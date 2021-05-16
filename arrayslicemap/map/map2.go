package main

import "fmt"

func main() {

	funcionarios := map[string]float64{
		"José":     12313.45,
		"Gabriela": 3313.45,
		"João":     5313.45,
		"Pedro":    2313.45,
	}

	funcionarios["Toddy"] = 3000.00
	delete(funcionarios, "Inexistente")

	for nome, salario := range funcionarios {
		fmt.Println(nome, salario)
	}
}
