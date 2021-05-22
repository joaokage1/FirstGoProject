package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Pessoa struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:mysql@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query("select * from usuarios where id >= (?)", 1)
	defer rows.Close()

	for rows.Next() {
		var p Pessoa
		rows.Scan(&p.id, &p.nome)
		fmt.Println(p)
	}
}
