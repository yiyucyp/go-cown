package mysqlite

import "database/sql"
import "fmt"
import _ "github.com/mattn/go-sqlite3"

// just test struct
type Users struct {
	UserID int
	Uname  string
}

// just test fun
func TestSqlite() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		fmt.Print(err)
		return
	}
	sql := `create table users(userId integer, uname text)`
	db.Exec(sql)
	sql = `insert into users(userId, unaame) values(1, 'Mike')`
	db.Exec(sql)
	rows, err := db.Query("select * from users")
	if err != nil {
		return
	}
	defer rows.Close()
	var users []Users = make([]Users, 0)
	for rows.Next() {
		var u Users
		rows.Scan(&u.UserID, &u.Uname)
		users = append(users, u)
	}
}
