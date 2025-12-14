package main

import "fmt"

func multiply(a int, b int) int {
	fmt.Println("Multiplicando %d * %d\n", a, b)
	return a * b
}

func sum(a int, b int) int {
	fmt.Println("Somando %d + %d\n", a, b)
	return a + b
}

func divide(a int, b int) int {
	fmt.Println("Dividindo %d / %d\n", a, b)
	return a / b
}

func substract(a int, b int) int {
	fmt.Println("Subtraindo %d - %d\n", a, b)
	return a - b
}
