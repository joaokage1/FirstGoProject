package main

import "fmt"

func main() {

	/*
		Mapas devem ser inicializados
	*/
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[151216354] = "Pedro"
	aprovados[654513453] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println("Vai ser apagado:", aprovados[151216354])

	delete(aprovados, 151216354)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
}
