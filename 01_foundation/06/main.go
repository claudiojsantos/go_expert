package main

import "fmt"

func main() {
	salarios := map[string]float64{"Claudio": 10000.0, "Renata": 2000.0, "Pedro": 5000.0}

	fmt.Printf("Salarios: %v\n", salarios)
	// Salarios: map[Claudio:10000 Pedro:5000 Renata:2000]

	fmt.Println(salarios["Claudio"])
	// 10000

	delete(salarios, "Pedro")

	fmt.Printf("Salarios: %v\n", salarios)

	salarios["Aldo"] = 20000.0

	fmt.Printf("Salarios: %v\n", salarios)

	sal1 := make(map[string]float64)
	sal2 := map[string]float64{}

	sal1["Aldo"] = 10000.0
	sal2["Antonio"] = 10000.0

	fmt.Printf("Salarios: %v\n", sal1)
	fmt.Printf("Salarios: %v\n", sal2)

	// Salarios: map[Aldo:10000]
	// Salarios: map[Antonio:10000]

	for nome, salario := range salarios {
		fmt.Printf("O salario do %s é %.2f\n", nome, salario)
	}

	// O salario do Claudio é 10000.00
	// O salario do Renata é 2000.00
	// O salario do Aldo é 20000.00

	for _, salario := range salarios {
		fmt.Printf("O salario é de %.2f\n", salario)
	}

	// O salario é de 10000.00
	// O salario é de 2000.00
	// O salario é de 20000.00
}
