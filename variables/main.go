package main

import "fmt"

const i float64 = 5.5

const pi = 3.14159

const timeToSleepInSeconds = 600

func main() {
	x := "Hello World!" //criação/declaração e atribuição de x -> :=
	fmt.Println(x)      //referencia a x

	var y string //criação/declaração de y -> var
	y = "Hello World!"
	fmt.Println(y) //referência a y

	var z, a, b int = 1, 2, 3 //criação/declaração e atribuição de z,a,b -> var
	fmt.Println(z)            //referência a z
	fmt.Println(a)            //referência a a
	fmt.Println(b)            //referência a b

	c, d := 3, 5
	fmt.Println(c) //referência a c
	fmt.Println(d) //referência a d

	var e int

	e = 1
	fmt.Println(e) //referência a e

	//estaticamente tipado + fortemente tipado

	e = 5

	var g bool
	g = true
	fmt.Println(g) //referência a g

	var h string
	h = "5"
	fmt.Println(h)

}
