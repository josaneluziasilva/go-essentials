package main

import ("fmt"
       "D:\Josane\Linuxtips\Go_Essentials\go-essentials\calc\operations"
)

func main() {
	fmt.Println("Calculadora #GoEssentiais")

	result := operations.Sum(5, 10) //criação e atribuição da variável
	fmt.Println(result)

	result = operations.Divide(10, 5) //só atribuindo
	fmt.Println(result)

	result = operations.Multiply(5, 5) //só atribuindo
	fmt.Println(result)

	result = operations.Substract(50, 10) //só atribuindo
	fmt.Println(result)

}
