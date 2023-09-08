package main

import "fmt"

func main() {
	var pessoa interface{} = "Claudio Santos"

	fmt.Println(pessoa)

	res, ok := pessoa.(string)
	fmt.Println(res, ok)
}
