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
	fmt.Println("Calculadora #GoEssentiais")

	result := sum(5, 10) //criação e atribuição
	fmt.Println(result)

	result = divide(10, 5) //só atribuindo
	fmt.Println(result)

	result = multiply(5, 5) //só atribuindo
	fmt.Println(result)

	result = substract(50, 10) //só atribuindo
	fmt.Println(result)

}
