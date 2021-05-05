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

	nfsServerOperationsGetattr.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Getattr))
	nfsServerOperationsSetattr.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Setattr))
	nfsServerOperationsLookup.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Lookup))
	nfsServerOperationsReadlink.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Readlink))
	nfsServerOperationsRead.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Read))
	nfsServerOperationsWrite.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Write))
	nfsServerOperationsCreate.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Create))
	nfsServerOperationsRemove.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Remove))
	nfsServerOperationsRename.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Rename))
	nfsServerOperationsLink.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Link))
	nfsServerOperationsSymlink.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Symlink))
	nfsServerOperationsMkdir.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Mkdir))
	nfsServerOperationsRmdir.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Rmdir))
	nfsServerOperationsReaddir.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Readdir))
	nfsServerOperationsRdirplus.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Rdirplus))
	nfsServerOperationsAccess.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Access))
	nfsServerOperationsMknod.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Mknod))
	nfsServerOperationsFsstat.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Fsstat))
	nfsServerOperationsFsinfo.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Fsinfo))
	nfsServerOperationsPathconf.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Pathconf))
	nfsServerOperationsCommit.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Commit))
	nfsServerOperationsLookupp.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Lookup))
	nfsServerOperationsSetclientid.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Setclientid))
	nfsServerOperationsSetclientidcfrm.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Setclientidcfrm))
	nfsServerOperationsOpen.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Open))
	nfsServerOperationsOpenattr.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Openattr))
	nfsServerOperationsOpendwgr.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Opendwgr))
	nfsServerOperationsOpencfrm.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Opencfrm))
	nfsServerOperationsDelepurge.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Delepurge))
	nfsServerOperationsDelreg.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Delreg))
	nfsServerOperationsGetfh.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Getfh))
	nfsServerOperationsLock.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Lock))
	nfsServerOperationsLockt.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Lockt))
	nfsServerOperationsLocku.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Locku))
	nfsServerOperationsClose.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Close))
	nfsServerOperationsVerify.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Verify))
	nfsServerOperationsNverify.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Nverify))
	nfsServerOperationsPutfh.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Putfh))
	nfsServerOperationsPutpubfh.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Putpubfh))
	nfsServerOperationsPutrootfh.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Putrootfh))
	nfsServerOperationsRenew.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Renew))
	nfsServerOperationsRestore.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Restore))
	nfsServerOperationsSavefh.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Savefh))
	nfsServerOperationsSecinfo.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Secinfo))
	nfsServerOperationsRellockown.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.Rellockown))
	nfsServerOperationsV4Create.WithLabelValues().Set(float64(stats.Nfsstat.Nfsv4.Serverstats.Operations.V4Create))

	return nil
}
