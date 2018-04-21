package LBC

import (
	"github.com/razerware/docker_beego/models"
	"time"
	"github.com/golang/glog"
)

func DataCollecter() {
	sleep := make(chan int)
	go func() {
		for {
			sleep <- 1
			time.Sleep(30 * time.Second)
		}
	}()
	for {
		select {
		case <-sleep:
			clusters := models.ListClusterAll()
			CountClusterNode()
			Check(clusters)
			glog.Info("....................time wait.....................")
			glog.Info("....................time wait.....................")
			glog.Info("....................time wait.....................")
			glog.Info("....................time wait.....................")
		}
	}
}

func Check(clusters []map[string]interface{}) {
	for _, cluster := range clusters {
		cs := CheckCluster(cluster)
		if !cs.UpdateCluster(){
			glog.Error("update cluster error")
			return
		}
		if cs.ELBFlag > 0 {
			glog.Info("cluster ",cs.SwarmID," result ",cs.ELBFlag)
			cs.ExpandCluster()
			continue
		} else if cs.ELBFlag < 0 {
			glog.Info("cluster ",cs.SwarmID," result ",cs.ELBFlag)
			cs.ReduceCluster()
			continue
		}else if cs.Ready==0 {
			glog.Info("cluster resizing")
			continue
		}
		glog.Info("cluster ",cs.SwarmID," result ",cs.ELBFlag)
		//if ELBFlag==0
		CheckServices(cs)
	}
	//CountClusterNode()
}
