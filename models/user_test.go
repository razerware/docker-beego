package models

import (
	"github.com/golang/glog"
	"testing"
)

func TestListUser(t *testing.T) {
	record := ListUser()
	glog.Info(record)
}
