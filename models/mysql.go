package models

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"encoding/json"
)

var DB *sql.DB
var DbError error

func MysqlConnect() {
	DB, DbError = sql.Open("mysql", "root:root@tcp(10.109.252.172)/docker")
	if DbError != nil {
		glog.Error(DbError)
	}
	db := DB
	err := db.Ping()
	if err != nil {
		glog.Error(err)
	} else {
		glog.Info("mysql connect ok")
	}
}

// This method return []map[string]interface{} rather than interface{}
// so we can get value of returns
func MysqlQuery(sql string) ([]map[string]interface{}) {
	if DbError != nil {
		glog.Error(DbError)
	}
	db := DB
	err := db.Ping()
	var record []map[string]interface{}
	if err != nil {
		glog.Error(err)
		return record
	} else {
		glog.Info("mysql connect ok")
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		glog.Error(err)
		return record
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		glog.Error(err)
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
			glog.V(1).Info(v)
			entry[col] = v
		}
		record = append(record, entry)
	}
	return record
}
func MysqlQuery_abandon(sql string, st interface{}) (interface{}) {
	record := MysqlQuery(sql)
	if len(record) < 1 {
		glog.Error("no result")
		return record
	}
	glog.V(1).Info(record)
	temp, err := json.Marshal(record)
	if err != nil {
		glog.Error(err)
	}
	err = json.Unmarshal(temp, &st)
	if err != nil {
		glog.Error(err)
	}
	return st
}
func MysqlInsert(sql string) (int64, int64, error) {
	glog.V(1).Info("In MysqlInsert func")
	if DbError != nil {
		glog.Error(DbError)
	}
	db := DB
	err := db.Ping()
	if err != nil {
		glog.Error(err)
		return 0, 0, err
	} else {
		glog.V(1).Info("mysql connect ok")
	}
	stmt, _ := db.Prepare(sql)
	defer stmt.Close()

	ret, err := stmt.Exec()

	if err != nil {
		glog.Error(err)
		return 0, 0, err
	}
	var LastInsertId, RowsAffected int64
	if LastInsertId, err = ret.LastInsertId(); nil == err {
		glog.V(1).Info("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err = ret.RowsAffected(); nil == err {
		glog.V(1).Info("RowsAffected:", RowsAffected)
	}
	return LastInsertId, RowsAffected, nil
}
