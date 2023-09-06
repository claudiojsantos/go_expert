package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Uf         string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
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

	claudio.Cidade = "Maceio"

	fmt.Println("Nome:", claudio.Nome, "Idade:", claudio.Idade, "Ativo:", claudio.Ativo, "Cidade:", claudio.Cidade)

	claudio.Endereco.Cidade = "Maceió"

	fmt.Println("Nome:", claudio.Nome, "Idade:", claudio.Idade, "Ativo:", claudio.Ativo, "Cidade:", claudio.Cidade, "Cidade:", claudio.Endereco.Cidade)
}
