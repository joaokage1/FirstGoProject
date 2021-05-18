package main

import "fmt"

type Imprimivel interface {
	toString() string
}

type Pessoa struct {
	nome      string
	sobrenome string
}

func (p Pessoa) toString() string {
	return fmt.Sprintf("Pessoa: %s %s", p.nome, p.sobrenome)
}

type Produto struct {
	nome  string
	preco float64
}

func (p Produto) toString() string {
	return fmt.Sprintf("Produto: %s %.2f", p.nome, p.preco)
}

func imprimir(x Imprimivel) {
	fmt.Println(x.toString())
}

//Interfaces são implementadas implicitamente

func main() {
	var coisa Imprimivel = Pessoa{
		nome:      "João",
		sobrenome: "Silva",
	}

	imprimir(coisa)

	coisa = Produto{
		nome:  "Computador",
		preco: 7999.99,
	}

	imprimir(coisa)
}
