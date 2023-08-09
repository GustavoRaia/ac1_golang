package main

import "fmt"

func main () {
	e_primo()
}

func e_primo() {
	var num int
	var contador int
	var i int
	var divisores int

	fmt.Println("Informe um número: ")
	fmt.Scanln(&num)

	for i < num {
		i++
		if num % i == 0 {
			contador += 1
		}
	}

	if contador == 2 {
		fmt.Println("É primo")
	} else {
		fmt.Println("Não é primo")
		fmt.Println("Divisores: ")
		for divisores < num {
			divisores++
			if num % divisores == 0{
				fmt.Println(divisores)
			}
		}
	}
}