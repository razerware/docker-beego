package models

import (
	"testing"
	"github.com/golang/glog"
	"fmt"
	"reflect"
)

func TestPullService(t *testing.T) {
	CheckService(User{Uid:1,Username:"lzy"})
}

func TestPullData(t *testing.T) {
	PullData("lzy", "SELECT sum(cpu) AS sum_cpu2 FROM lzy.autogen.container "+
		"WHERE time > now() - 5m AND serviceid='nu1nlcy2u07drt5d0lzuwmh9d'; "+
		"SELECT mean(cpu) AS mean_cpu FROM lzy.autogen.vm WHERE time > now() - 5m AND "+
		"swarmid='tiaho6zrcp13e56u5wr0d7arl' GROUP BY time(1m) FILL(null)")
}

func TestLzytest(t *testing.T) {
	Lzytest(nil)
}

func TestJudgement(t *testing.T) {
	ir := PullData("lzy", "SELECT sum(cpu) AS sum_cpu2 FROM lzy.autogen.container "+
		"WHERE time > now() - 1d AND serviceid='nu1nlcy2u07drt5d0lzuwmh9d'; "+
		"SELECT mean(cpu) AS mean_cpu FROM lzy.autogen.vm WHERE time > now() - 1d AND "+
		"swarmid='tiaho6zrcp13e56u5wr0d7arl'")
	Judgement(ir, ServiceStat{})
}

func TestInfluxSelect(t *testing.T) {
	sql:=InfluxSelect("sum", "cpu", "lzy", "autogen", "container",
		"1d", "serviceid", "nu1nlcy2u07drt5d0lzuwmh9d")
	glog.Info(sql)
	ir:=PullData("lzy",sql)
	glog.Info(ir)

}

func TestCountClusterNode(t *testing.T) {
	CountClusterNode()
}
func hello() {
	fmt.Println("Hello world!")
}
func TestLzytest2(t *testing.T) {
	var h2 func()
	if true{
		h2=hello
	}
	fv := reflect.ValueOf(h2)
	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	fv.Call(nil)
}

func TestCheckCluster(t *testing.T) {
	for i:=0;i<5;i++{
		CheckCluster(User{Uid:1,Username:"lzy"})
	}
	CheckCluster(User{Uid:1,Username:"lzy"})
}