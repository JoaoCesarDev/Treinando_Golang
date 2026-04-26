package main

import (
	"fmt"
)

type Produto struct {
	nome  string
	preco float64
}

func sistemaProdutos() {
	produto1 := Produto{"Produto 1", 10.0}
	produto2 := Produto{"Produto 2", 20.0}
	produto3 := Produto{"Produto 3", 30.0}

	produtos := []Produto{produto1, produto2, produto3}

	maiorPreco := produtos[0]

	for _, produto := range produtos {
		if produto.preco > maiorPreco.preco {
			maiorPreco = produto
		}
	}

		fmt.Printf(" O produto de Maior Preço: %s, %.2f\n", maiorPreco.nome, maiorPreco.preco)

}

func prods() {
	sistemaProdutos()
}