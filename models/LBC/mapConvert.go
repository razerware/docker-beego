package LBC

import (
	"encoding/json"
	"github.com/golang/glog"
)

func MapToStruct(data map[string]interface{}, v interface{}) error {

	if temp, ok := json.Marshal(data); ok == nil {
		if err := json.Unmarshal(temp, &v); err != nil {
			glog.Error(ok)
			return ok
		} else {
			glog.V(1).Info("map->json->struct ok! ", v)
		}
	} else {
		return ok
	}
	return nil
}

func MapToStructSlice(data []map[string]interface{}, v interface{}) error {

	if temp, ok := json.Marshal(data); ok == nil {
		if err := json.Unmarshal(temp, &v); err != nil {
			glog.Error(ok)
			return ok
		} else {
			glog.V(1).Info("map->json->struct ok! ", v)
		}
	} else {
		return ok
	}
	return nil
}
