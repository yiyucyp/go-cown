package crowtool

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ExampleSqlite(dataSourceName string) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer db.Close()
	//select
	rows, err := db.Query("select* from users")
	if err != nil {
		fmt.Print(err)
		return
	}
	for i := 0; rows.Next(); i++ {

	}

	db.Exec("delete from dx_running")
}
func OpenDB(dataSource string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSource)
	return db, err
}

func ExecDB(db *sql.DB, sqlString string) {
	db.Exec(sqlString)
}
