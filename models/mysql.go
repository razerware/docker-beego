package models

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var DB *sql.DB
var DbError error

func MysqlConnect() {
	DB, DbError = sql.Open("mysql", "root:root@tcp(10.109.252.172)/docker")
	if DbError != nil {
		fmt.Println(DbError)
	}
	db := DB
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("mysql connect ok")
	}
}

func MysqlQuery(sql string) ([]map[string]interface{}) {
	if DbError != nil {
		fmt.Println(DbError)
	}
	db := DB
	err := db.Ping()
	var record  []map[string]interface{}
	if err != nil {
		fmt.Println(err)
		return record
	} else {
		fmt.Println("mysql connect ok")
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println("gg")
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	//rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return record
	}
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	//https://studygolang.com/articles/10512
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			fmt.Println(v)
			entry[col] = v
		}
		record=append(record,entry)
	}
	return record
}
