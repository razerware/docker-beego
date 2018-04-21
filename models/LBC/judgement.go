package LBC

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"github.com/golang/glog"
	"strings"
)

type compareData struct {
	tempCpu  int
	tempMem  int
	cpuUpper int
	cpuLower int
	memUpper int
	memLower int
	mapKey   string //serviceId or swarmId
}

var (
	//MontiorMap map[string]Monitor
	monitorMap = utils.NewBeeMap()
	clusterMap = utils.NewBeeMap()
)

type Monitor struct {
	CpuFlag int
	MemFlag int
}

type Stat interface {
	CombineData(int, int) compareData
}

func Judgement(ir InfluxResp, stat Stat) int {
	if len(ir.Results) < 1 {
		glog.Error("InfluxResp not contain both CPU or MEM data")
		return 0
	}
	var tempCpu, tempMem int
	for _, i := range ir.Results {

		if len(i.Series) == 0 {
			glog.Error("Need both Cpu and Mem data from InfluxDB")
			return 0
		}
		if strings.Contains(i.Series[0].Columns[1], "cpu") {
			v, _ := i.Series[0].Values[0][1].(float64)
			tempCpu = int(v)
			glog.V(1).Info("this is Cpu data:", tempCpu)
		} else if strings.Contains(i.Series[0].Columns[1], "mem") {
			v, _ := i.Series[0].Values[0][1].(float64)
			tempMem = int(v)
			glog.V(1).Info("this is Mem data:", tempMem)
		}
	}

	cd := stat.CombineData(tempCpu, tempMem)
	vMonitor := monitorMap.Get(cd.mapKey)
	return JudgeByData(vMonitor, cd)

	//switch stat := statInterface.(type) {
	//case ServiceStat:
	//	glog.V(1).Info("check Service")
	//	vMonitor = monitorMap.Get(stat.ServiceId)
	//	cd := compareData{tempCpu, tempMem, stat.CpuUpper, stat.CpuLower,
	//		stat.MemUpper, stat.MemLower, stat.ServiceId}
	//	return JudgeByData(vMonitor, cd)
	//case ClusterStat:
	//	glog.V(1).Info("check Cluster")
	//	vMonitor = monitorMap.Get(stat.SwarmId)
	//	cd := compareData{tempCpu, tempMem, stat.CpuUpper, stat.CpuLower,
	//		stat.MemUpper, stat.MemLower, stat.SwarmId}
	//	return JudgeByData(vMonitor, cd)
	//default:
	//	return 0
	//}
}

func JudgeByData(vMonitor interface{}, data compareData) int {
	monitor, _ := vMonitor.(Monitor)
	glog.Info("threshol is cpu: ",data.cpuLower,"-",data.cpuUpper," mem: ",data.memLower,"-",data.memUpper)
	monitor.CpuFlag = AlgoCompare(data.tempCpu, data.cpuUpper, data.cpuLower, monitor.CpuFlag)
	monitor.MemFlag = AlgoCompare(data.tempMem, data.memUpper, data.memLower, monitor.MemFlag)
	DynamicPeriod, _ := beego.AppConfig.Int("DynamicPeriod")
	//periodFlag := 1<<uint(DynamicPeriod) - 1
	periodFlag:=DynamicPeriod
	glog.Info("After compare monitor is: ", monitor)
	result := 0
	if monitor.CpuFlag >= periodFlag || monitor.MemFlag >= periodFlag {
		//expand/reduce need some CD so they cannot operated together
		monitor.CpuFlag = 0
		monitor.MemFlag = 0
		glog.Info("Need Expand")
		result = 1
	} else if monitor.CpuFlag <= -periodFlag && monitor.MemFlag <= -periodFlag {
		//expand is "or" but reduce need "and"
		monitor.CpuFlag = 0
		monitor.MemFlag = 0
		glog.Info("Need reduce")
		result = -1
	} else {
		glog.Info("No Need to expand/reduce")
		result = 0
	}
	//flag cannot exceed periodFlag
	monitor.CpuFlag = thresholdFlag(monitor.CpuFlag, periodFlag)
	monitor.MemFlag = thresholdFlag(monitor.MemFlag, periodFlag)
	monitorMap.Set(data.mapKey, monitor)
	glog.V(1).Info(monitorMap.Items())
	glog.V(1).Info("Judgement complete")
	return result
}

func AlgoCompare(test int, upper int, lower int, flag int) int {

	if test > upper {
		if flag > 0 {
			flag++
		} else {
			flag = 1
		}
	} else if test < lower {
		if flag < 0 {
			flag--
		} else {
			flag = -1
		}
	} else {
		flag = 0
	}
	return flag
}

func thresholdFlag(input int, threshold int) int {
	if input > threshold {
		return threshold
	}
	if input <= -threshold {
		return -threshold
	}
	return input
}
