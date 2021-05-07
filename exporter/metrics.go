package exporter

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	nfsServerOperationsGetattr = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_getattr",
	}, []string{})
	nfsServerOperationsSetattr = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_setattr",
	}, []string{})
	nfsServerOperationsLookup = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_lookup",
	}, []string{})
	nfsServerOperationsReadlink = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_readlink",
	}, []string{})
	nfsServerOperationsRead = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_read",
	}, []string{})
	nfsServerOperationsWrite = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_write",
	}, []string{})
	nfsServerOperationsCreate = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_create",
	}, []string{})
	nfsServerOperationsRemove = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_remove",
	}, []string{})
	nfsServerOperationsRename = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_rename",
	}, []string{})
	nfsServerOperationsLink = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_link",
	}, []string{})
	nfsServerOperationsSymlink = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_symlink",
	}, []string{})
	nfsServerOperationsMkdir = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_mkdir",
	}, []string{})
	nfsServerOperationsRmdir = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_rmdir",
	}, []string{})
	nfsServerOperationsReaddir = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_readdir",
	}, []string{})
	nfsServerOperationsRdirplus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_rdirplus",
	}, []string{})
	nfsServerOperationsAccess = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_access",
	}, []string{})
	nfsServerOperationsMknod = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_mknod",
	}, []string{})
	nfsServerOperationsFsstat = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_fsstat",
	}, []string{})
	nfsServerOperationsFsinfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_fsinfo",
	}, []string{})
	nfsServerOperationsPathconf = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_pathconf",
	}, []string{})
	nfsServerOperationsCommit = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_commit",
	}, []string{})
	nfsServerOperationsLookupp = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_lookupp",
	}, []string{})
	nfsServerOperationsSetclientid = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_setclientid",
	}, []string{})
	nfsServerOperationsSetclientidcfrm = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_setclientidcfrm",
	}, []string{})
	nfsServerOperationsOpen = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_open",
	}, []string{})
	nfsServerOperationsOpenattr = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_openattr",
	}, []string{})
	nfsServerOperationsOpendwgr = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_opendwgr",
	}, []string{})
	nfsServerOperationsOpencfrm = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_opencfrm",
	}, []string{})
	nfsServerOperationsDelepurge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_delepurge",
	}, []string{})
	nfsServerOperationsDelreg = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_delreg",
	}, []string{})
	nfsServerOperationsGetfh = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_getfh",
	}, []string{})
	nfsServerOperationsLock = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_lock",
	}, []string{})
	nfsServerOperationsLockt = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_lockt",
	}, []string{})
	nfsServerOperationsLocku = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_locku",
	}, []string{})
	nfsServerOperationsClose = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_close",
	}, []string{})
	nfsServerOperationsVerify = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_verify",
	}, []string{})
	nfsServerOperationsNverify = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_nverify",
	}, []string{})
	nfsServerOperationsPutfh = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_putfh",
	}, []string{})
	nfsServerOperationsPutpubfh = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_putpubfh",
	}, []string{})
	nfsServerOperationsPutrootfh = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_putrootfh",
	}, []string{})
	nfsServerOperationsRenew = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_renew",
	}, []string{})
	nfsServerOperationsRestore = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_restore",
	}, []string{})
	nfsServerOperationsSavefh = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_savefh",
	}, []string{})
	nfsServerOperationsSecinfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_saveinfo",
	}, []string{})
	nfsServerOperationsRellockown = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_rellockown",
	}, []string{})
	nfsServerOperationsV4Create = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_v4_create",
	}, []string{})

	poudriereStatusQueue = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "poudriere_status_queue",
	}, []string{"ports", "jail"})
	poudriereStatusBuilt = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "poudriere_status_built",
	}, []string{"ports", "jail"})
	poudriereStatusFail = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "poudriere_status_fail",
	}, []string{"ports", "jail"})
	poudriereStatusSkip = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "poudriere_status_skip",
	}, []string{"ports", "jail"})
	poudriereStatusIgnore = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "poudriere_status_ignore",
	}, []string{"ports", "jail"})
	poudriereStatusRemain = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "poudriere_status_remain",
	}, []string{"ports", "jail"})
)

func init() {
	prometheus.MustRegister(
		nfsServerOperationsGetattr,
		nfsServerOperationsSetattr,
		nfsServerOperationsLookup,
		nfsServerOperationsReadlink,
		nfsServerOperationsRead,
		nfsServerOperationsWrite,
		nfsServerOperationsCreate,
		nfsServerOperationsRemove,
		nfsServerOperationsRename,
		nfsServerOperationsLink,
		nfsServerOperationsSymlink,
		nfsServerOperationsMkdir,
		nfsServerOperationsRmdir,
		nfsServerOperationsReaddir,
		nfsServerOperationsRdirplus,
		nfsServerOperationsAccess,
		nfsServerOperationsMknod,
		nfsServerOperationsFsstat,
		nfsServerOperationsFsinfo,
		nfsServerOperationsPathconf,
		nfsServerOperationsCommit,
		nfsServerOperationsLookupp,
		nfsServerOperationsSetclientid,
		nfsServerOperationsSetclientidcfrm,
		nfsServerOperationsOpen,
		nfsServerOperationsOpenattr,
		nfsServerOperationsOpendwgr,
		nfsServerOperationsOpencfrm,
		nfsServerOperationsDelepurge,
		nfsServerOperationsDelreg,
		nfsServerOperationsGetfh,
		nfsServerOperationsLock,
		nfsServerOperationsLockt,
		nfsServerOperationsLocku,
		nfsServerOperationsClose,
		nfsServerOperationsVerify,
		nfsServerOperationsNverify,
		nfsServerOperationsPutfh,
		nfsServerOperationsPutpubfh,
		nfsServerOperationsPutrootfh,
		nfsServerOperationsRenew,
		nfsServerOperationsRestore,
		nfsServerOperationsSavefh,
		nfsServerOperationsSecinfo,
		nfsServerOperationsRellockown,
		nfsServerOperationsV4Create,
	)
}

func StartMetricsServer(bindAddr string) {
	d := http.NewServeMux()
	d.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(bindAddr, d)
	if err != nil {
		log.Fatal("Failed to start metrics server, error is:", err)
	}
}
