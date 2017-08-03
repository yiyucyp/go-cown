package main

import (
	"big_cown/mysqlite"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type JsonType1 struct {
	ErrorCode int    `json:"errcode"`
	ErrorMsg  string `json:"errmsg"`
}

func main() {
	mysqlite.TestSqlite()
	b := []byte(`{"errcode":-3,"errmsg":"117.29.177.162无匹配的数据!"}`)
	var fjson interface{}
	var err error
	err = json.Unmarshal(b, &fjson)
	if err != nil {
		return
	}
	var vv JsonType1
	//err = json.Unmarshal(b, &vv)
	fmt.Print(err)
	resp, err1 := http.Get("http://ww.mydogo.com/api/kaokeacc/get_taokeacc.php")
	if err1 != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
	json.Unmarshal(body, &vv)
	//m := fjson.(map[string]interface{})
}

func TestMysql() {
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer db.Close()
	rows, err := db.Query("select*from users")
	if err != nil {
		fmt.Print(err)
		return
	}
	for i := 0; rows.Next(); i++ {

	}
}
