package main

import (
	"fmt"

	"github.com/cod3rcursos/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.google.com/", "https://www.ahnegao.com.br/"),
		html.Titulo("https://www.pelando.com.br/quente", "https://www.youtube.com/"),
	)

	fmt.Println("Primeiro mais rápido :", <-c)
	fmt.Println("Segundo mais rápido :", <-c)
	fmt.Println("Terceiro mais rápido :", <-c)
	fmt.Println("Quarto mais rápido :", <-c)
}
