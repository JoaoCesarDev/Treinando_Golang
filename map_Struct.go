package main

import "fmt"

type Produto1 struct {
    Nome  string
    Preco float64
}

func main2() {
    produtos := make(map[string]Produto1)

    produtos["p1"] = Produto1{"Notebook", 3500}
    produtos["p2"] = Produto1{"Mouse", 100}

    fmt.Println(produtos["p1"])
}