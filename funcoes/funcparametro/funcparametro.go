package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(f1 func(int, int) int, p1, p2 int) int {
	return f1(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 5, 2)
	fmt.Println(resultado)
}
