package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Pessoa :)
type Pessoa struct {
	ID   int    `json: "id"`
	Nome string `json: "nome"`
}

// UsuarioHandler analisa o request e delega para funcao adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}

}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:mysql@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var p Pessoa
	db.QueryRow("select * from usuarios where id = (?)", id).Scan(&p.ID, &p.Nome)

	json, _ := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:mysql@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, _ := db.Query("select * from usuarios")
	defer rows.Close()

	var pessoas []Pessoa

	for rows.Next() {
		var p Pessoa
		rows.Scan(&p.ID, &p.Nome)
		pessoas = append(pessoas, p)
	}

	json, _ := json.Marshal(pessoas)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
