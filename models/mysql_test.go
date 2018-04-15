package models

import (
	"flag"
	"os"
	"github.com/golang/glog"
	"testing"
	"github.com/astaxie/beego"
)

func TestMain(m *testing.M) {
	flag.Set("alsologtostderr", "true")
	//flag.Set("log_dir", "/tmp")
	flag.Set("v", "1")
	flag.Parse()
	MysqlConnectTest()
	configPath := "D:\\myProject\\docker-beego\\src\\docker-beego\\conf\\app.conf"
	glog.Info(configPath)
	if err := beego.LoadAppConfig("ini", configPath); err != nil {
		panic(err)
	}
	ret := m.Run()
	os.Exit(ret)

}

func TestMysqlConnect(t *testing.T) {
	MysqlConnect()
}
func TestMysqlQuery(t *testing.T) {
	record := MysqlQuery("SELECT username from `user`")
	glog.Info(record)
}
