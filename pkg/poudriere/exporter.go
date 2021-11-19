package poudriere

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	poudriereStatusDesc = prometheus.NewDesc(
		"poudriere_status",
		"Poudrere status",
		[]string{"ports", "jail", "status"},
		nil,
	)
)

type Exporter struct {
	logger log.Logger
}

type Stat struct {
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

func NewExporter(logger log.Logger) (*Exporter, error) {
	return &Exporter{
		logger: log.With(logger, "exporter", "poudriere"),
	}, nil
}

func (s *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- poudriereStatusDesc
}

func (s *Exporter) Collect(ch chan<- prometheus.Metric) {
	cmd := exec.Command("/usr/local/bin/poudriere", "status", "-fH")

	out, err := cmd.Output()
	if err != nil {
		_ = level.Error(s.logger).Log("err", err.Error())
		return
	}

	r := bytes.NewReader(out)

	stats, err := readPoudriereStats(r)
	if err != nil {
		_ = level.Error(s.logger).Log("err", err.Error())
		return
	}

	for _, s := range stats {
		ch <- prometheus.MustNewConstMetric(poudriereStatusDesc, prometheus.GaugeValue, float64(s.Queue), s.Ports, s.Jail, "queue")
		ch <- prometheus.MustNewConstMetric(poudriereStatusDesc, prometheus.GaugeValue, float64(s.Built), s.Ports, s.Jail, "built")
		ch <- prometheus.MustNewConstMetric(poudriereStatusDesc, prometheus.GaugeValue, float64(s.Fail), s.Ports, s.Jail, "fail")
		ch <- prometheus.MustNewConstMetric(poudriereStatusDesc, prometheus.GaugeValue, float64(s.Skip), s.Ports, s.Jail, "skip")
		ch <- prometheus.MustNewConstMetric(poudriereStatusDesc, prometheus.GaugeValue, float64(s.Ignore), s.Ports, s.Jail, "ignore")
		ch <- prometheus.MustNewConstMetric(poudriereStatusDesc, prometheus.GaugeValue, float64(s.Remain), s.Ports, s.Jail, "remain")
	}
}

func readPoudriereStats(r io.Reader) ([]Stat, error) {
	var stats []Stat

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

		stat := Stat{
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
