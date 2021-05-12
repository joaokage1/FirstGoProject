package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	//cuidado
	fmt.Println("Teste " + string(97))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//String para int
	num, _ := strconv.Atoi("123") //ignorando retorno de erro
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") // true ou 1

	if b {
		fmt.Println("Verdadeiro")
	}
}
