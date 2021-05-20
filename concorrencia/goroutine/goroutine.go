package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d) \n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "Porque você não fala comigo", 3)
	//fale("João", "Só posso falar depois de você", 1)

	//go fale("Maria", "ei...", 500)
	//go fale("João", "opa...", 500)

	//time.Sleep(time.Second * 2)
	//fmt.Println("Fim")

	go fale("Maria", "Entendi", 10)
	fale("João", "Parabéns", 5)
}
