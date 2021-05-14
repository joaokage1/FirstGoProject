package main

import "fmt"

func main() {

	i := 1 //Go Não tem aritmética de ponteiros

	var p *int = nil

	p = &i // pegando endereço da variável
	*p++
	i++

	fmt.Println(p, *p, i, &i)
}
