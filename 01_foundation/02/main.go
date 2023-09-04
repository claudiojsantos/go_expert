package main

const a = "Essa é uma variavel constante e não pode ser alterada"

var (
	b string
	c int
	d bool
	e float64

	f string  = "Essa é uma variavel do tipo string inicializada"
	g int     = 10
	h bool    = true
	i float64 = 10.5
)

func main() {
	j := "Essa é uma variavel do tipo string com shorthand - declarada apenas dentro da função"
	k := 10
	l := true
	m := 10.5

	println(a)
	println("------------")
	println(b)
	println(c)
	println(d)
	println(e)
	println("------------")
	println(f)
	println(g)
	println(h)
	println(i)
	println("------------")
	println(j)
	println(k)
	println(l)
	println(m)
}
