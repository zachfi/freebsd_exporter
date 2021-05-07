package exporter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadPoudriereStats(t *testing.T) {
	file, err := os.Open("../testdata/poudriere_status.txt")
	require.NoError(t, err)
	defer file.Close()

	stats, err := readPoudriereStats(file)
	require.NoError(t, err)
	require.NotNil(t, stats)

	expected := []PoudriereStat{
		{
			Set:    "-",
			Ports:  "default",
			Jail:   "larch12",
			Build:  "2021-05-06_00h32m33s",
			Status: "done",
			Queue:  "19",
			Built:  "19",
			Fail:   "0",
			Skip:   "0",
			Ignore: "0",
			Remain: "0",
			Time:   "00:33:07",
			Logs:   "/usr/local/poudriere/data/logs/bulk/larch12-default/2021-05-06_00h32m33s",
		},
	}

	require.Equal(t, expected[0], stats[0])
}
