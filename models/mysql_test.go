package models

import (
	"github.com/golang/glog"
	"testing"
)

func TestMysqlConnect(t *testing.T) {
	MysqlConnect()
}
func TestMysqlQuery(t *testing.T) {
	record := MysqlQuery("SELECT username from `user`")
	glog.Info(record)
}
