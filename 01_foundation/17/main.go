package main

import "fmt"

type Cliente struct {
	nome string
}

func (c *Cliente) andou() {
	fmt.Printf("O Cliente %v andou\n", c.nome)

	fmt.Printf("O nome do cliente é %v\n", c.nome)
}

func main() {
	cj := Cliente{nome: "Claudio Santos"}

	cj.andou()
}
