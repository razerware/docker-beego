package LBC

import (
	"github.com/golang/glog"
	"github.com/razerware/docker_beego/models"
)

type ServiceStat struct {
	ServiceId     string `json:"service_id"`
	SwarmId       string `json:"swarm_id"`
	Replication   int    `json:"replication"`
	DesireReplica int    `json:"desire_replica"`
	models.ElasticInfo
}

func CheckServices(cs ClusterStat) ([]ServiceStat, []ServiceStat) {
	record := cs.ListServiceC()
	if record == nil {
		glog.Error("No services")
		return nil, nil
	}
	var ss []ServiceStat
	//map->json->struct
	if err := MapToStructSlice(record, &ss); err != nil {
		glog.Error(err)
		return nil, nil
	} else {
		glog.Info("map transfer ok")
	}
	var expandList, reduceList []ServiceStat
	for _, s := range ss {
		if s.DesireReplica > s.Replication {
			glog.Info("This service is not in good condition")
			continue
		}
		kv:=map[string]string{"service_id":s.ServiceId}
		CpuSql := InfluxSelect("mean", "cpu", cs.Username, "autogen", "container",
			"30s", kv)
		MemSql := InfluxSelect("mean", "mem", cs.Username, "autogen", "container",
			"30s", kv)
		ir := PullData(cs.Username, CpuSql+";"+MemSql)
		glog.Info(".....................Judging Service: ",s.ServiceId,".....................")
		result := Judgement(ir, s)

		if result > 0 {
			expandList = append(expandList, s)
		} else if result < 0 {
			reduceList = append(reduceList, s)
		}
		glog.Info(s.ServiceId," Checking result is ",result)
		//if s.DesireReplica<nodeNumber{
		//
		//}
	}
	return expandList, reduceList
}

func (ss ServiceStat) CombineData(tempCpu int, tempMem int) compareData {
	return compareData{tempCpu, tempMem, ss.CpuUpper, ss.CpuLower,
		ss.MemUpper, ss.MemLower, ss.ServiceId}
}

func (ss ServiceStat) CheckExpand() bool {
	if ss.DesireReplica+ss.Step <= ss.UpperLimit {
		return true
	}
	return false
}

func (ss ServiceStat) CheckReduce() bool {
	if ss.DesireReplica-ss.Step >= ss.LowerLimit {
		return true
	}
	return false
}
