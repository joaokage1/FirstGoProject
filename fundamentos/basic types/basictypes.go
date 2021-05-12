package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (positivos)... uinnt8(byte) uint16(short) uint32(int) uint64(long)

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O maior valor para um int", i1)
	fmt.Println("O tipo do i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("O i2 é", i2)

	//numeros reais (float32, float64)
	var x float32 = 49.99 // ou var x = (float32)49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Olá meu nome é João"
	fmt.Println(s1)
	fmt.Println("Tamanho da string 1 é", len(s1))

	//string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	João`
	fmt.Println("Tamanho da string 2 é", len(s2))

	//char??? -> Não existe char em GO
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
}
