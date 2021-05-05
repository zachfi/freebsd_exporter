package exporter

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

func Scrape() error {

	cmd := exec.Command("/usr/bin/nfsstat", "-E", "--libxo=json")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}

	var stats NFSStat
	err = json.Unmarshal(out.Bytes(), &stats)
	if err != nil {
		return err
	}

	nfsServerOperationsGetattr.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Getattr))

	return nil
}
