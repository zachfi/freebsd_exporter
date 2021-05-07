package exporter

import (
	"bufio"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type PoudriereExporter struct {
}

type PoudriereStat struct {
	Set    string
	Ports  string
	Jail   string
	Build  string
	Status string
	Queue  int
	Built  int
	Fail   int
	Skip   int
	Ignore int
	Remain int
	Time   time.Duration
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
		poudriereStatusQueue.WithLabelValues(s.Ports, s.Jail).Set(float64(s.Queue))
		poudriereStatusBuilt.WithLabelValues(s.Ports, s.Jail).Set(float64(s.Built))
		poudriereStatusFail.WithLabelValues(s.Ports, s.Jail).Set(float64(s.Fail))
		poudriereStatusSkip.WithLabelValues(s.Ports, s.Jail).Set(float64(s.Skip))
		poudriereStatusIgnore.WithLabelValues(s.Ports, s.Jail).Set(float64(s.Ignore))
		poudriereStatusRemain.WithLabelValues(s.Ports, s.Jail).Set(float64(s.Remain))
	}

	return nil
}

func readPoudriereStats(r io.Reader) ([]PoudriereStat, error) {
	var stats []PoudriereStat

	zero, err := time.Parse("15:04:05", "00:00:00")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")

		queue, err := strconv.Atoi(parts[5])
		if err != nil {
			return nil, err
		}

		built, err := strconv.Atoi(parts[6])
		if err != nil {
			return nil, err
		}

		fail, err := strconv.Atoi(parts[7])
		if err != nil {
			return nil, err
		}

		skip, err := strconv.Atoi(parts[8])
		if err != nil {
			return nil, err
		}

		ignore, err := strconv.Atoi(parts[9])
		if err != nil {
			return nil, err
		}

		remain, err := strconv.Atoi(parts[10])
		if err != nil {
			return nil, err
		}

		statTime, err := time.Parse("15:04:05", parts[11])
		if err != nil {
			return nil, err
		}

		stat := PoudriereStat{
			Set:    parts[0],
			Ports:  parts[1],
			Jail:   parts[2],
			Build:  parts[3],
			Status: parts[4],
			Queue:  queue,
			Built:  built,
			Fail:   fail,
			Skip:   skip,
			Ignore: ignore,
			Remain: remain,
			Time:   statTime.Sub(zero),
			Logs:   parts[12],
		}

		stats = append(stats, stat)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
