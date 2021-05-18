package main

import "fmt"

type Esportivo interface {
	ligarTurbo()
}

type Luxuoso interface {
	fazerBaliza()
}

type EsportivoLuxuoso interface {
	Esportivo
	Luxuoso
}

type Bmw7 struct{}

func (b Bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b Bmw7) fazerBaliza() {
	fmt.Println("Fazendo baliza...")
}

func main() {
	carro := Bmw7{}

	carro.ligarTurbo()
	carro.fazerBaliza()
}
