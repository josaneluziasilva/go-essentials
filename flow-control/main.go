package main

import "fmt"

const MAX int = 100

func main() {
	var n int
	n = 0

	expr := n < MAX

	//	fmt.Println("1 - Oi, estou em uma execução 100% linear")
	//	fmt.Println("2 - Oi, estou em uma execução 100% linear")
	//	fmt.Println("3 - Oi, estou em uma execução 100% linear")

	if expr {
		fmt.Printf("%d é menor do que MAX == %d\n", n, MAX)
	} else {
		fmt.Printf("%d é maior do que MAX == %d\n", n, MAX)
	}

	for ; n < MAX; n = n + 1 {
		fmt.Println(n % 2)
		if n%2 == 0 {
			fmt.Printf("%d é par\n", n)
		} else {
			fmt.Printf("%d é ímpar\n", n)
		}

	}

	for i := 0; i < 10; i++ {
	}

	for {
		n = n + 1
		fmt.Println(n)

		if n > MAX {
			break
		}
	} // loop infinito (não existe while em Go, apenas for infinito)

	switch n < MAX {
	case true:
		{
			fmt.Println("n é menor do que MAX")
		}
	case false:
		{
			fmt.Println("n não é menor do que MAX")
		}
	}
}
