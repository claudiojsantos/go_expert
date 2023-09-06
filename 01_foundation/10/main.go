package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	claudio := Cliente{
		Nome:  "Claudio",
		Idade: 37,
		Ativo: true,
	}

	fmt.Println("Nome:", claudio.Nome, "Idade:", claudio.Idade, "Ativo:", claudio.Ativo)

	claudio.Ativo = false

	fmt.Println("Nome:", claudio.Nome, "Idade:", claudio.Idade, "Ativo:", claudio.Ativo)
}
