package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	nome      string
	sobrenome string
}

func (p Pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *Pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")

	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {

	p1 := Pessoa{"Pedro", "Silva"}

	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Joao Silva")

	fmt.Println(p1.getNomeCompleto())
}
