package main

import "fmt"

//1.1 задача  объявить переменные всех основных типов, 3 спосабами var с типом var без типа
//:=
//вывести значения через %T

func main() {
	var a int = 1
	var b float64 = 1.22
	var c string = "assdasd"
	var d bool = true
	var f byte = 'a'
	var g rune = 'b'

	var a2 = 2
	var b2 = 2.11
	var c2 = "world"
	var d2 = "false"
	var f2 = 'c'
	var g2 = 'd'

	a3 := 1
	b3 := 2.22
	c3 := "fuck"
	d3 := false
	f3 := 'h'
	g3 := 'j'

	fmt.Printf("a = %v\n, (%T)\n", a, a)
	fmt.Printf("b = %v\n, (%T)\n", b, b)
	fmt.Printf("c = %v\n, (%T)\n", c, c)
	fmt.Printf("d = %v\n, (%T)\n", d, d)
	fmt.Printf("f = %v\n, (%T)\n", f, f)
	fmt.Printf("g = %v\n, (%T)\n", g, g)

	fmt.Printf("a2 = %v\n, (%T)\n", a2, a2)
	fmt.Printf("b2 = %v\n, (%T)\n", b2, b2)
	fmt.Printf("c2 = %v\n, (%T)\n", c2, c2)
	fmt.Printf("d2 = %v\n, (%T)\n", d2, d2)
	fmt.Printf("f2 = %v\n, (%T)\n", f2, f2)
	fmt.Printf("g2 = %v\n, (%T)\n", g2, g2)

	fmt.Printf("a3 = %v\n, (%T)\n", a3, a3)
	fmt.Printf("b3 = %v\n, (%T)\n", b3, b3)
	fmt.Printf("c3 = %v\n, (%T)\n", c3, c3)
	fmt.Printf("d3 = %v\n, (%T)\n", d3, d3)
	fmt.Printf("f3 = %v\n, (%T)\n", f3, f3)
	fmt.Printf("g3 = %v\n, (%T)\n", g3, g3)

	try()
}

func try() {
	a4 := 123
	b4 := 345

	c4 := a4
	a4 = b4
	b4 = c4

	fmt.Println(a4, b4)

	c5, d5 := 5, 10

	c5, d5 = d5, c5

	fmt.Println(c5, d5)
}
