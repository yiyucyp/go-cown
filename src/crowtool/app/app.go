package main

import (
	"github.com/crowtool"
)

func main() {
	crowtool.ExampleJosn()
	crowtool.ExampleSqlite("./db.db")

	crowtool.ExampleMysql("root:root@test")
}
