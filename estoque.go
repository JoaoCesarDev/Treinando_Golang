package main

import "fmt"

func main1() {
    estoque := map[string]int{
        "Notebook": 10,
        "Mouse": 50,
    }

    // adicionando
    estoque["Teclado"] = 20

    // vendendo produto
    estoque["Mouse"] -= 5

    fmt.Println(estoque)
}