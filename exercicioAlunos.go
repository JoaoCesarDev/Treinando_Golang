package main

import "fmt"

type Aluno struct {
    Nome  string
    Nota int
}

func main() {
    alunos := make(map[string]Aluno)

    alunos["a1"] = Aluno{"João", 20}
    alunos["a2"] = Aluno{"Maria", 22}
	alunos["a3"] = Aluno{"Pedro", 19}
	alunos["a4"] = Aluno{"Ana", 21}

    media := 0

	for _, aluno := range alunos {
		media += aluno.Nota
	}
	media /= len(alunos)

	notaMaisAlta := alunos["a1"]

	for _, aluno := range alunos {
		if aluno.Nota > notaMaisAlta.Nota {
			notaMaisAlta = aluno
		}
	}

	fmt.Printf("A média das notas é: %d\n", media)

	fmt.Printf("A nota mais alta é do aluno: %s com nota %d\n", notaMaisAlta.Nome, notaMaisAlta.Nota)
}
