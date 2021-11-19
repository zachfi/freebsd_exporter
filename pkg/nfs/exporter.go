package nfs

import (
	"bytes"
	"encoding/json"
	"os/exec"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	nfsServerOpertionsDesc = prometheus.NewDesc(
		"nfs_server_operations",
		"NFS server operations",
		[]string{"operation"},
		nil,
	)
)

type Exporter struct {
	logger log.Logger
}

func NewExporter(logger log.Logger) (*Exporter, error) {
	return &Exporter{
		logger: log.With(logger, "exporter", "nfs"),
	}, nil
}

func (s *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- nfsServerOpertionsDesc
}

func (s *Exporter) Collect(ch chan<- prometheus.Metric) {
	cmd := exec.Command("/usr/bin/nfsstat", "-E", "--libxo=json")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		_ = level.Error(s.logger).Log("err", err.Error())
		return
	}

	var stats Stat
	err = json.Unmarshal(out.Bytes(), &stats)
	if err != nil {
		_ = level.Error(s.logger).Log("err", err.Error())
		return
	}

	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Getattr), "getattr")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Setattr), "setattr")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Lookup), "lookup")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Readlink), "readlink")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Read), "read")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Write), "write")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Create), "create")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Remove), "remove")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Rename), "rename")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Link), "link")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Symlink), "symlink")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Rmdir), "mkdir")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Rmdir), "rmdir")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Readdir), "readdir")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Rdirplus), "rdirplus")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Access), "access")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Mknod), "mknod")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Fsstat), "fsstat")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Fsinfo), "fsinfo")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Pathconf), "pathconf")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Commit), "commit")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Setclientid), "setclientid")

	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Setclientidcfrm), "setclientidcfrm")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Clientstats.Operations.Open), "open")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Opendwgr), "opendwgr")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Openattr), "openattr")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Opencfrm), "opencfrm")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Delepurge), "delepurge")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Delreg), "delreg")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Getfh), "getfh")

	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Lock), "lock")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Lockt), "lockt")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Locku), "locku")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Close), "close")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Verify), "verify")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Nverify), "nverify")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Putfh), "putfh")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Putpubfh), "putpubfh")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Putrootfh), "putrootfh")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Renew), "renew")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Restore), "restore")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Savefh), "savefh")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Secinfo), "secinfo")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Rellockown), "rellockown")
	ch <- prometheus.MustNewConstMetric(
		nfsServerOpertionsDesc, prometheus.GaugeValue, float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.V4Create), "v4create")
}

type Stat struct {
	Version string `json:"__version"`
	Nfsstat struct {
		Nfsv4 struct {
			Clientstats struct {
				Operations struct {
					Getattr       int `json:"getattr"`
					Setattr       int `json:"setattr"`
					Lookup        int `json:"lookup"`
					Readlink      int `json:"readlink"`
					Read          int `json:"read"`
					Write         int `json:"write"`
					Create        int `json:"create"`
					Remove        int `json:"remove"`
					Rename        int `json:"rename"`
					Link          int `json:"link"`
					Symlink       int `json:"symlink"`
					Mkdir         int `json:"mkdir"`
					Rmdir         int `json:"rmdir"`
					Readdir       int `json:"readdir"`
					Rdirplus      int `json:"rdirplus"`
					Access        int `json:"access"`
					Mknod         int `json:"mknod"`
					Fsstat        int `json:"fsstat"`
					Fsinfo        int `json:"fsinfo"`
					Pathconf      int `json:"pathconf"`
					Commit        int `json:"commit"`
					Setclientid   int `json:"setclientid"`
					Setclientidcf int `json:"setclientidcf"`
					Lock          int `json:"lock"`
					Lockt         int `json:"lockt"`
					Locku         int `json:"locku"`
					Open          int `json:"open"`
					Opencfr       int `json:"opencfr"`
					Nfsv41        struct {
						Opendowngr   int `json:"opendowngr"`
						Close        int `json:"close"`
						Rellckown    int `json:"rellckown"`
						Freestateid  int `json:"freestateid"`
						Putrootfh    int `json:"putrootfh"`
						Delegret     int `json:"delegret"`
						Getacl       int `json:"getacl"`
						Setacl       int `json:"setacl"`
						Exchangeid   int `json:"exchangeid"`
						Createsess   int `json:"createsess"`
						Destroysess  int `json:"destroysess"`
						Destroyclid  int `json:"destroyclid"`
						Layoutget    int `json:"layoutget"`
						Getdevinfo   int `json:"getdevinfo"`
						Layoutcomit  int `json:"layoutcomit"`
						Layoutreturn int `json:"layoutreturn"`
						Reclaimcompl int `json:"reclaimcompl"`
						Readdatas    int `json:"readdatas"`
						Writedatas   int `json:"writedatas"`
						Commitdatas  int `json:"commitdatas"`
						Openlayout   int `json:"openlayout"`
						Createlayout int `json:"createlayout"`
					} `json:"nfsv41"`
					Nfsv42 struct {
						Ioadvise    int `json:"ioadvise"`
						Allocate    int `json:"allocate"`
						Copy        int `json:"copy"`
						Seek        int `json:"seek"`
						Seekdatas   int `json:"seekdatas"`
						Getextattr  int `json:"getextattr"`
						Setextattr  int `json:"setextattr"`
						Rmextattr   int `json:"rmextattr"`
						Listextattr int `json:"listextattr"`
					} `json:"nfsv42"`
				} `json:"operations"`
				Client struct {
					Openowner int `json:"openowner"`
					Opens     int `json:"opens"`
					Lockowner int `json:"lockowner"`
					Locks     int `json:"locks"`
					Delegs    int `json:"delegs"`
					Localown  int `json:"localown"`
					Localopen int `json:"localopen"`
					Locallown int `json:"locallown"`
					Locallock int `json:"locallock"`
				} `json:"client"`
				RPC struct {
					Timedout int `json:"timedout"`
					Invalid  int `json:"invalid"`
					Xreplies int `json:"xreplies"`
					Retries  int `json:"retries"`
					Requests int `json:"requests"`
				} `json:"rpc"`
				Cache struct {
					Attrhits    int `json:"attrhits"`
					Attrmisses  int `json:"attrmisses"`
					Lkuphits    int `json:"lkuphits"`
					Lkupmisses  int `json:"lkupmisses"`
					Biorhits    int `json:"biorhits"`
					Biormisses  int `json:"biormisses"`
					Biowhits    int `json:"biowhits"`
					Biowmisses  int `json:"biowmisses"`
					Biorlhits   int `json:"biorlhits"`
					Biorlmisses int `json:"biorlmisses"`
					Biodhits    int `json:"biodhits"`
					Biodmisses  int `json:"biodmisses"`
					Direhits    int `json:"direhits"`
					Diremisses  int `json:"diremisses"`
					Cache       struct {
					} `json:"cache"`
				} `json:"cache"`
			} `json:"clientstats"`
			Serverstats struct {
				Operations struct {
					Getattr         int `json:"getattr"`
					Setattr         int `json:"setattr"`
					Lookup          int `json:"lookup"`
					Readlink        int `json:"readlink"`
					Read            int `json:"read"`
					Write           int `json:"write"`
					Create          int `json:"create"`
					Remove          int `json:"remove"`
					Rename          int `json:"rename"`
					Link            int `json:"link"`
					Symlink         int `json:"symlink"`
					Mkdir           int `json:"mkdir"`
					Rmdir           int `json:"rmdir"`
					Readdir         int `json:"readdir"`
					Rdirplus        int `json:"rdirplus"`
					Access          int `json:"access"`
					Mknod           int `json:"mknod"`
					Fsstat          int `json:"fsstat"`
					Fsinfo          int `json:"fsinfo"`
					Pathconf        int `json:"pathconf"`
					Commit          int `json:"commit"`
					Lookupp         int `json:"lookupp"`
					Setclientid     int `json:"setclientid"`
					Setclientidcfrm int `json:"setclientidcfrm"`
					Open            int `json:"open"`
					Openattr        int `json:"openattr"`
					Opendwgr        int `json:"opendwgr"`
					Opencfrm        int `json:"opencfrm"`
					Delepurge       int `json:"delepurge"`
					Delreg          int `json:"delreg"`
					Getfh           int `json:"getfh"`
					Lock            int `json:"lock"`
					Lockt           int `json:"lockt"`
					Locku           int `json:"locku"`
					Close           int `json:"close"`
					Verify          int `json:"verify"`
					Nverify         int `json:"nverify"`
					Putfh           int `json:"putfh"`
					Putpubfh        int `json:"putpubfh"`
					Putrootfh       int `json:"putrootfh"`
					Renew           int `json:"renew"`
					Restore         int `json:"restore"`
					Savefh          int `json:"savefh"`
					Secinfo         int `json:"secinfo"`
					Rellockown      int `json:"rellockown"`
					V4Create        int `json:"v4create"`
					Nfsv41          struct {
						Backchannelctrl int `json:"backchannelctrl"`
						Bindconntosess  int `json:"bindconntosess"`
						Exchangeid      int `json:"exchangeid"`
						Createsess      int `json:"createsess"`
						Destroysess     int `json:"destroysess"`
						Freestateid     int `json:"freestateid"`
						Getdirdeleg     int `json:"getdirdeleg"`
						Getdevinfo      int `json:"getdevinfo"`
						Getdevlist      int `json:"getdevlist"`
						Layoutcommit    int `json:"layoutcommit"`
						Layoutget       int `json:"layoutget"`
						Layoutreturn    int `json:"layoutreturn"`
						Secinfnoname    int `json:"secinfnoname"`
						Sequence        int `json:"sequence"`
						Setssv          int `json:"setssv"`
						Teststateid     int `json:"teststateid"`
						Wantdeleg       int `json:"wantdeleg"`
						Destroyclid     int `json:"destroyclid"`
						Reclaimcompl    int `json:"reclaimcompl"`
					} `json:"nfsv41"`
					Nfsv42 struct {
						Allocate    int `json:"allocate"`
						Copy        int `json:"copy"`
						Copynotify  int `json:"copynotify"`
						Deallocate  int `json:"deallocate"`
						Ioadvise    int `json:"ioadvise"`
						Layouterror int `json:"layouterror"`
						Layoutstats int `json:"layoutstats"`
						Offloadcncl int `json:"offloadcncl"`
						Offloadstat int `json:"offloadstat"`
						Readplus    int `json:"readplus"`
						Seek        int `json:"seek"`
						Writesame   int `json:"writesame"`
						Clone       int `json:"clone"`
						Getextattr  int `json:"getextattr"`
						Setextattr  int `json:"setextattr"`
						Listextattr int `json:"listextattr"`
						Rmextattr   int `json:"rmextattr"`
					} `json:"nfsv42"`
				} `json:"operations"`
				Server struct {
					Clients   int `json:"clients"`
					Openowner int `json:"openowner"`
					Opens     int `json:"opens"`
					Lockowner int `json:"lockowner"`
					Locks     int `json:"locks"`
					Delegs    int `json:"delegs"`
				} `json:"server"`
				Cache struct {
					Inprog    int `json:"inprog"`
					Nonidem   int `json:"nonidem"`
					Misses    int `json:"misses"`
					Cachesize int `json:"cachesize"`
					Tcppeak   int `json:"tcppeak"`
				} `json:"cache"`
			} `json:"serverstats"`
		} `json:"nfsv4"`
	} `json:"nfsstat"`
}
