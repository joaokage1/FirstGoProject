package main

import "fmt"

func main() {

	funcsProLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 3231.00,
			"Guga Pereira":   5432.34,
		},
		"J": {
			"João": 3046.50,
			"José": 6026.50,
		},
		"P": {
			"Pedro": 4046.50,
		},
	}

	for letra, funcs := range funcsProLetra {
		fmt.Println(letra, funcs)
	}

	delete(funcsProLetra, "P")

	fmt.Println("REMOVENDO P")
	for letra, funcs := range funcsProLetra {
		fmt.Println(letra, funcs)
	}

	fmt.Println()
	for _, funcs := range funcsProLetra {
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
