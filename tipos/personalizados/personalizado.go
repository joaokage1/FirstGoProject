package main

import "fmt"

type Nota float64

func (n Nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n Nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 6.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	}

	return "R"
}

func main() {
	fmt.Println(notaParaConceito(9.0))
	fmt.Println(notaParaConceito(2.1))
}
