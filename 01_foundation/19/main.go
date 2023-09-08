package main

type MyNumber int

type Number interface {
	~int | float64
}

func soma[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

func main() {
	m := map[string]int{
		"item1": 10,
		"item2": 20,
		"item3": 30,
	}

	m2 := map[string]float64{
		"item1": 10.43,
		"item2": 20.3,
		"item3": 30.22,
	}

	m3 := map[string]MyNumber{
		"item1": 10,
		"item2": 20,
		"item3": 30,
	}

	println(soma(m))
	println(soma(m2))
	println(soma(m3))
}
