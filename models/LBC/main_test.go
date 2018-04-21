package LBC

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/razerware/docker_beego/models"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Set("alsologtostderr", "true")
	//flag.Set("log_dir", "/tmp")
	flag.Set("v", "0")
	flag.Parse()
	models.MysqlConnectTest()
	configPath := "D:\\myProject\\docker-beego\\src\\github.com\\razerware\\docker_beego\\conf\\app.conf"
	glog.Info(configPath)
	if err := beego.LoadAppConfig("ini", configPath); err != nil {
		panic(err)
	}
	ret := m.Run()
	os.Exit(ret)

}
