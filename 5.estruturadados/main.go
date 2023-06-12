package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func main() {
	pessoa := Pessoa{
		nome:  "Maria",
		idade: 25,
	}
	fmt.Println("Nome:", pessoa.nome)
	fmt.Println("Idade:", pessoa.idade)
}
