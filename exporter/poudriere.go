package exporter

import (
	"bufio"
	"io"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

type PoudriereExporter struct {
}

type PoudriereStat struct {
	Set    string
	Ports  string
	Jail   string
	Build  string
	Status string
	Queue  string
	Built  string
	Fail   string
	Skip   string
	Ignore string
	Remain string
	Time   string
	Logs   string
}

func (s *PoudriereExporter) Scrape() error {
	cmd := exec.Command("/usr/local/bin/poudriere", "status", "-fH")
	out, err := cmd.StdoutPipe()

	err = cmd.Run()
	if err != nil {
		return err
	}

	stats, err := readPoudriereStats(out)
	if err != nil {
		return err
	}

	for _, s := range stats {
		log.Tracef("s: %+s", s)
	}

	return nil
}

func readPoudriereStats(r io.Reader) ([]PoudriereStat, error) {
	var stats []PoudriereStat

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")

		stat := PoudriereStat{
			Set:    parts[0],
			Ports:  parts[1],
			Jail:   parts[2],
			Build:  parts[3],
			Status: parts[4],
			Queue:  parts[5],
			Built:  parts[6],
			Fail:   parts[7],
			Skip:   parts[8],
			Ignore: parts[9],
			Remain: parts[10],
			Time:   parts[11],
			Logs:   parts[12],
		}

		stats = append(stats, stat)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
