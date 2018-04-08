package controllers

type User struct {
	Uid int `json:"uid"`
	Uname string `json:"uname"` //nickname
	Username string `json:"username"`
	Password string `json:"password"`
}
