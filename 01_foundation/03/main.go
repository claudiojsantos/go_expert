package main

import "fmt"

type ID int

var (
	a ID = 1
)

func main() {
	fmt.Printf("O tipo de a é %T\n", a)
	fmt.Printf("O tipo de a é %v\n", a)
}
