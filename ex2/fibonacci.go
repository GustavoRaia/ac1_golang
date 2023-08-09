package main

import "fmt"

func main() {
	fibo()
}

func fibo() {
	var num int
	i := 1
	a, b := 0, 1

	fmt.Println("Informe um número: ")
	fmt.Scanln(&num)

	for i < num {
		a, b = b, a+b
		i++
	}
	fmt.Print(num)
	fmt.Print("-ésimo: ")
	fmt.Println(b)
}