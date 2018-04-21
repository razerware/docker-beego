package LBC

import (
	"github.com/golang/glog"
	"github.com/razerware/docker_beego/models"
	"testing"
)

func TestCheckCluster(t *testing.T) {
	clusters := models.ListClusterAll()

	CheckCluster(clusters[0])

}

func TestClusterStat_ReduceCluster(t *testing.T) {
	clusters := models.ListClusterAll()
	var cs ClusterStat
	CountClusterNode()
	MapToStruct(clusters[0], &cs)
	cs.ReduceCluster()
	glog.Info(cs.CheckReduce())
}

func TestClusterStat_GetNodeVersion(t *testing.T) {
	clusters := models.ListClusterAll()
	var cs ClusterStat
	CountClusterNode()
	MapToStruct(clusters[0], &cs)
	version := GetNodeInfo(cs.ManagerIP, "s8t7t1qjyr1i5by48gxql47m4")
	glog.Info(version)
}

func TestClusterStat_UpdateServices(t *testing.T) {
	clusters := models.ListClusterAll()
	var cs ClusterStat
	CountClusterNode()
	MapToStruct(clusters[0], &cs)
	cs.UpdateServices([]string{"pp4tit2ovgwjjsemhdee3vwlx"})
}

func TestClusterStat_UpdateCluster(t *testing.T) {
	clusters := models.ListClusterAll()
	var cs ClusterStat
	CountClusterNode()
	MapToStruct(clusters[0], &cs)
	cs.UpdateCluster()
}

func TestCountClusterNode(t *testing.T) {
	a:=[]int{1,2,3,4,5}
	a=a[:1]
	glog.Info(a)
}