package main

import "fmt"

type Carro struct {
	nome       string
	velocidade int
}

type Ferrary struct {
	Carro       //campos anônimos
	turboLigado bool
}

func main() {
	f := Ferrary{}
	f.nome = "F40"
	f.velocidade = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s esta com turbo ligado ? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}
