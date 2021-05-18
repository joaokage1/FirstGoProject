package main

import "fmt"

type Produto struct {
	nome     string
	preco    float64
	desconto float64
}

//metodo funcao com receiver

func (p Produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 Produto

	produto1 = Produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := Produto{"Notebook", 4000.0, 0.1}
	fmt.Println(produto2, produto2.precoComDesconto())
}
