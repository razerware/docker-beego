package models

import (
	"fmt"
	"github.com/razerware/docker_beego/client"
	"io/ioutil"
	"net/http"
)

func PullImages(name string, tag string) {
	url := client.GetClient()
	req, _ := http.NewRequest("POST", url, nil)
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("fromImage", "hello-world")
	q.Add("tag", "latest")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())

	client := &http.Client{}
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
