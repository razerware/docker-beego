package models

import (
	"github.com/golang/glog"
	"testing"
)

func TestGetServiceInfo(t *testing.T) {
	si := GetServiceInfo("10.109.252.172", "pp4tit2ovgwjjsemhdee3vwlx")
	glog.Info(si)
}
