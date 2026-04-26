package main

import (
	"fmt"
)

func contarPalavras() {
	palavras := []string{"Go", "java", "python", "JavaScript", "C++", "Go", "python", "java"}

	contador := make(map[string]int)

	for _, palavra := range palavras {
		contador[palavra]++	
	}

	for palavra, quantidade := range contador {
		fmt.Printf("A palavra '%s' aparece %d vezes.\n", palavra, quantidade)
	}
}

func contarPalavrasRepetidas() {
	contarPalavras()
}