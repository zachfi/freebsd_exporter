package poudriere

import (
	"strings"
	"testing"
	"time"
)

// readPoudriereStats must tolerate the variants of `poudriere status -fH`
// output we see in the wild: HH:MM:SS for active/done builds, and "0" or
// empty when no time has been recorded yet (idle/queued/never-built).
func TestReadPoudriereStats(t *testing.T) {
	const input = "" +
		"-\tdefault\tlarch12\t2021-05-06_00h32m33s\tdone\t19\t19\t0\t0\t0\t0\t00:33:07\t/var/log/larch12-default\n" +
		// Idle row with "0" in the time column — previously caused the
		// whole scrape to abort with `parsing time "0" as "15:04:05"`.
		"-\tdefault\tlarch13\t-\tidle\t0\t0\t0\t0\t0\t0\t0\t-\n" +
		// Empty time field: also seen in idle states.
		"-\tpersonal\tlarch12\t-\tidle\t0\t0\t0\t0\t0\t0\t\t-\n" +
		// Active build with HH:MM:SS.
		"-\tpersonal\tlarch13\t2021-05-06_14h26m19s\tparallel_build\t4\t2\t0\t0\t0\t2\t00:01:48\t/var/log/larch13-personal\n" +
		// Truncated/malformed line: skipped, must not abort.
		"garbage\n" +
		// Blank line: skipped.
		"\n"

	stats, err := readPoudriereStats(strings.NewReader(input))
	if err != nil {
		t.Fatalf("readPoudriereStats returned unexpected error: %v", err)
	}

	if got, want := len(stats), 4; got != want {
		t.Fatalf("got %d stats, want %d", got, want)
	}

	cases := []struct {
		jail   string
		ports  string
		status string
		want   time.Duration
	}{
		{"larch12", "default", "done", 33*time.Minute + 7*time.Second},
		{"larch13", "default", "idle", 0},
		{"larch12", "personal", "idle", 0},
		{"larch13", "personal", "parallel_build", 1*time.Minute + 48*time.Second},
	}

	for i, c := range cases {
		s := stats[i]
		if s.Jail != c.jail || s.Ports != c.ports || s.Status != c.status {
			t.Errorf("stat[%d]: got (jail=%q ports=%q status=%q), want (%q %q %q)",
				i, s.Jail, s.Ports, s.Status, c.jail, c.ports, c.status)
		}
		if s.Time != c.want {
			t.Errorf("stat[%d] (%s/%s): got Time=%v, want %v", i, c.ports, c.jail, s.Time, c.want)
		}
	}
}
