package client

import (
	"github.com/astaxie/beego"
)
type DockerClient struct {
	Ip string
}

func  GetClient() string{
	client:=beego.AppConfig.String("docker_api_ip")
	return client
}
