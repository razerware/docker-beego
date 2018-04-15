package models

import (
	"testing"
	"github.com/golang/glog"
)

func TestListUser(t *testing.T) {
	record := ListUser()
	glog.Info(record)
}
