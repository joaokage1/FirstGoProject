package main

import (
	"fmt"
)

func main() {

	var i int = 0
	for i <= 10 { //nao tem while
		fmt.Print(i)
		i++
	}

	fmt.Println()
	for i := 0; i <= 20; i += 2 {
		fmt.Print(i)
	}

	fmt.Println("\nMinsturando...")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("É par ")
		} else {
			fmt.Print("É impar ")
		}
	}

	//for {
	// laço infinito
	//fmt.Println("Para sempre...")
	//time.Sleep(time.Second)
	//}
}
