package LBC

import (
	"fmt"
	"github.com/golang/glog"
	"testing"
)

func TestInfluxResp_SortInfluxResp(t *testing.T) {
	sql := fmt.Sprintf("SELECT mean(cpu) AS mean_cpu FROM lzy.autogen.container " +
		"WHERE time > now() - 5m GROUP BY service_id")
	ir := PullData("lzy", sql)
	result := ir.SortInfluxResp("desc")
	glog.Info(result)
}
