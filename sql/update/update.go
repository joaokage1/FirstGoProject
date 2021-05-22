package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = (?) where id = (?)")
	stmt.Exec("Jaqueline", 1)

	stmt2, _ := db.Prepare("delete from usuarios where id = (?)")
	stmt2.Exec(3)
}
