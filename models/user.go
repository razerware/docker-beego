package models

import (
	"fmt"
	"github.com/golang/glog"
)

type User struct {
	Uid      int    `json:"uid"`
	Uname    string `json:"uname"` //nickname
	Username string `json:"username"`
	Password string `json:"password"`
}

func ListUser() []User {
	//var uid int
	//c.Ctx.Input.Bind(&uid, "uid")
	sql := fmt.Sprintf("SELECT username,uid FROM `user`")
	record := MysqlQuery(sql)
	userList := []User{}
	if record==nil{
		glog.Info("record empty")
		return userList
	}
	for _, m := range record {
		value1, ok1 := m["uid"].(int)
		value2, ok2 := m["username"].(string)
		if ok1 && ok2 {
			userList = append(userList, User{Uid: value1, Username: value2})
		} else {
			glog.Error("ListUser error")
		}
	}
	return userList
	//fj:=FrontendJson{0,"",len(record),record}
	//c.Data["json"] = fj
	//c.ServeJSON()
}


