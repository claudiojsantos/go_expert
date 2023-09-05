package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum2(3, 4))
	fmt.Println(sum3(5, 60))
	fmt.Println(sum4(5, 60))

	valor, err := sum4(5, 404)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)
}

func sum(a int, b int) int {
	return a + b
}

func sum2(a, b int) int {
	return a + b
}

func sum3(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	} else {
		return a + b, false
	}
}

func sum4(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("valor muito alto")
	} else {
		return a + b, nil
	}
}
