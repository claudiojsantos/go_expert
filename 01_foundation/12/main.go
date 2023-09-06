package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Uf         string
}

func (c Cliente) MudarStatus() {
	c.Ativo = !c.Ativo
	fmt.Printf("Status do cliente %s alterado para %v\n", c.Nome, c.Ativo)
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

	claudio.Ativo = !claudio.Ativo

	fmt.Println("Status do cliente", claudio.Nome, "alterado para", claudio.Ativo)

	fmt.Println("Nome:", claudio.Nome, "Idade:", claudio.Idade, "Ativo:", claudio.Ativo, "Cidade:", claudio.Cidade, "Cidade:", claudio.Endereco.Cidade)

	claudio.MudarStatus()
}
