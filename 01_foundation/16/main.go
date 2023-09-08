package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Uf         string
}

type Pessoa interface {
	MudarStatus()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) MudarStatus() {
	c.Ativo = !c.Ativo
	fmt.Printf("Status do cliente %s alterado para %v\n", c.Nome, c.Ativo)
}

func AlteracaoStatus(pessoa Pessoa) {
	pessoa.MudarStatus()
}

func main() {
	claudio := &Cliente{
		Nome:  "Claudio",
		Idade: 37,
		Ativo: true,
	}

	fmt.Println("Nome:", claudio.Nome, "Status Origem:", claudio.Ativo)

	claudio.MudarStatus()

	AlteracaoStatus(claudio)
}
