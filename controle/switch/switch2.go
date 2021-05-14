package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch { // switch true

	case t.Hour() < 12:
		fmt.Printf("%s \nBom dia", t)
	case t.Hour() < 18:
		fmt.Printf("%s \nBoa tarde", t)
	default:
		fmt.Printf("%s \nBoa noite", t)
	}
}
