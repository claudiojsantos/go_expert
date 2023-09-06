package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(5, 404, 20, 1, 3) * 3
	}()

	fmt.Println(total)

	fmt.Println(sum(5, 404, 20, 1, 3))
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
