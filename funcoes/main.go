package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}

func substract(a int, b int) int {
	return a - b
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(multiply(5, 20))
	fmt.Println(divide(10, 2))
	fmt.Println(substract(10, 10))
}
