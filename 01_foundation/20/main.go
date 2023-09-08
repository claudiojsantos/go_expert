package main

import (
	"fmt"

	"go_expert/matematica"
)

func main() {
	s := matematica.Soma(1, 2, 3, 4, 5)

	fmt.Println("O resultado da soma é", s)
}
