package LBC

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/razerware/docker_beego/models"
	"sort"
)

type InfluxResp struct {
	Results []struct {
		StatementID int `json:"statement_id"`
		Series      `json:"series"`
	} `json:"results"`
}

type Series []struct {
	Name string `json:"name"`
	Tags struct {
		NodeID    string `json:"node_id"`
		ServiceID string `json:"service_id"`
	} `json:"tags"`
	Columns []string        `json:"columns"`
	Values  [][]interface{} `json:"values"`
}

func (s Series) Len() int {
	return len(s)
}
func (s Series) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Series) Less(i, j int) bool {
	return s[i].Values[0][1].(float64) < s[j].Values[0][1].(float64)
}

func PullData(user string, sql string) InfluxResp {
	paras := map[string]string{}
	var ir InfluxResp
	paras["pretty"] = "true"
	paras["db"] = user
	paras["q"] = sql
	paras["u"] = "admin"
	paras["p"] = "admin"
	code, body, err := models.MyGet("http://"+beego.AppConfig.String("InfluxUrl")+"/query", paras)
	//code, body, err := MyGet(beego.AppConfig.String("InfluxUrl"), paras)
	if err != nil || code > 200 {
		glog.Error(err)
		return ir
	}
	glog.V(1).Info("InfluxData data")
	glog.V(1).Info(string(body))
	json.Unmarshal(body, &ir)
	return ir
}

func InfluxSelect(function, field, db, policy, measurement, period string,kv map[string]string) string {
	sql := fmt.Sprintf("SELECT %s(%s) AS %s_%s FROM %s.%s.%s WHERE time > now() - %s",
		function, field, function, field, db, policy, measurement, period)
	for k,v :=range kv{
		sql+=fmt.Sprintf(" AND %s='%s'",k,v)
	}
	return sql
}

func (ir InfluxResp) SortInfluxResp(mode string) []string {
	var result []string
	if len(ir.Results[0].Series) < 1 {
		return result
	}
	s := ir.Results[0].Series
	if mode == "asc" {
		sort.Sort(s)
	} else {
		sort.Sort(sort.Reverse(s))
	}
	for _, i := range s {
		if i.Tags.NodeID != "" {
			result = append(result, i.Tags.NodeID)
		} else if i.Tags.ServiceID != "" {
			result = append(result, i.Tags.ServiceID)
		}

	}
	return result
}
