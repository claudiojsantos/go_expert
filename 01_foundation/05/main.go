package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Printf("O tipo de slice é %T\n", slice)

	fmt.Printf("len = %d e cap = %d %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])
	// len = 0 e cap = 10 []

	fmt.Printf("len = %d e cap = %d %v\n", len(slice[:4]), cap(slice[:4]), slice[:4])
	// len = 4 e cap = 10 [10 20 30 40]

	fmt.Printf("len = %d e cap = %d %v\n", len(slice[2:]), cap(slice[2:]), slice[2:])
	// len = 8 e cap = 8 [30 40 50 60 70 80 90 100]

	slice = append(slice, 110)
	fmt.Printf("len = %d e cap = %d %v\n", len(slice[:2]), cap(slice[:2]), slice[:2])

}
