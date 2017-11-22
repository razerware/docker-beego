package models

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var DB *sql.DB
var DbError error
func MysqlConnect() {
	DB, DbError = sql.Open("mysql", "root:root@tcp(10.109.252.172)/docker")
}
