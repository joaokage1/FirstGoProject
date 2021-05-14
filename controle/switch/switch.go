package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	var conceito string
	switch nota {
	case 10:
		fallthrough
	case 9:
		conceito = "Conceito A"
	case 8, 7:
		conceito = "Conceito B"
	case 6, 5:
		conceito = "Conceito C"
	case 4, 3:
		conceito = "Conceito D"
	case 2, 1, 0:
		conceito = "Conceito R"
	default:
		conceito = "Nota inválida"
	}

	return conceito
}

func main() {
	fmt.Println(notaParaConceito(4.9))
	fmt.Println(notaParaConceito(5.9))
	fmt.Println(notaParaConceito(7.9))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(1))
}
