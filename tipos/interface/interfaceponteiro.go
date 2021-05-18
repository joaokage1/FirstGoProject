package main

import "fmt"

type Esportivo interface {
	ligarTurbo()
}

type Ferrari struct {
	modelo      string
	turboLigado bool
	velocidade  int
}

func (f *Ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := Ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	var carro2 Esportivo = &Ferrari{"F41", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
