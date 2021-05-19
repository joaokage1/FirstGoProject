package main

import "math"

// Iniciando com letra maiúscula é Público (Dentro e fora do pacote)
// Letra minúscula é Privado (Dentro do pacote)

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {

	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return cx, cy
}

//Distancia é responsavel por calcular a distancia entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
