package main

import (
	"encoding/json"
	"fmt"
)

type Produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"` //slice
}

func main() {
	//struct para json
	p1 := Produto{1, "Notebook", 3999.99, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))
	//fmt.Println(p1)

	//json para struct
	var p2 Produto
	jsonString := `{"id":2,"nome":"Caneta","preco":2.99,"tags":["Papelaria"]}`

	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)
}
