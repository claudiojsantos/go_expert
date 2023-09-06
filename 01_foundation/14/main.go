package main

import "fmt"

func main() {
	a := 10
	b := a

	fmt.Println(a)
	fmt.Println(b)

	var ponteiro *int = &a
	*ponteiro = 20

	fmt.Println(a)
	fmt.Println(b)

	c := &a

	fmt.Println(c)

	*c = 30

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*c)
}
