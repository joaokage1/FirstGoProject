package main

import "fmt"

type Item struct {
	produtoId  int
	quantidade int
	preco      float64
}

type Pedido struct {
	userId int
	itens  []Item
}

func (p Pedido) valorTotal() float64 {

	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return total
}

func main() {
	pedido := Pedido{
		userId: 1,
		itens: []Item{
			{
				produtoId:  1,
				quantidade: 4,
				preco:      4.99,
			},
			{
				produtoId:  2,
				quantidade: 2,
				preco:      8.99,
			},
			{
				produtoId:  6,
				quantidade: 10,
				preco:      1.00,
			},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())
}
