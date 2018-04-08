package controllers

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/golang/glog"
)

func MyPost(url string, send_bytes []byte) (int,[]byte,error){
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(send_bytes))
	if err != nil {
		// handle error
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Error(err)
		// handle error
	} else {
		glog.V(1).Info("MyPost successed!")
	}
	code:=resp.StatusCode
	return code,body,err
}

func MyGet(url string, stats interface{}) {
	client := &http.Client{}
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle error
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	cs := stats
	json.Unmarshal(body, &cs)
	fmt.Println(cs)
}
