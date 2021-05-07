package exporter

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestReadPoudriereStats(t *testing.T) {
	file, err := os.Open("../testdata/poudriere_status.txt")
	require.NoError(t, err)
	defer file.Close()

	stats, err := readPoudriereStats(file)
	require.NoError(t, err)
	require.NotNil(t, stats)

	zero, err := time.Parse("15:04:05", "00:00:00")
	require.NoError(t, err)
	t.Logf("zero.Minutes: %+v", zero)

	statTime, err := time.Parse("15:04:05", "00:33:28")
	require.NoError(t, err)
	t.Logf("statTime.Minutes: %+v", statTime)

	expected := []PoudriereStat{
		{
			Set:    "-",
			Ports:  "default",
			Jail:   "larch12",
			Build:  "2021-05-06_00h32m33s",
			Status: "done",
			Queue:  19,
			Built:  19,
			Fail:   0,
			Skip:   0,
			Ignore: 0,
			Remain: 0,
			Time:   statTime.Sub(zero),
			Logs:   "/usr/local/poudriere/data/logs/bulk/larch12-default/2021-05-06_00h32m33s",
		},
	}

	require.Equal(t, expected[0], stats[0])
}
