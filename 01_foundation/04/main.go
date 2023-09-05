package main

import "fmt"

func main() {
	var meuArray [3]int

	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Printf("O tipo de meuArray é %T\n", meuArray)
	fmt.Printf("A ultima posição é %v\n", meuArray[len(meuArray)-1])
	fmt.Printf("A primeira posição é %v\n", meuArray[0])

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d e o valor é %d\n", i, v)
	}
}
