package crowtool

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ExampleMysql(dataSourceName string) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer db.Close()
}
