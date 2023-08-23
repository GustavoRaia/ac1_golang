package main

import "fmt"

func main() {
	e_primo()
	fibo()
	dia_semana()
	e_bissexto()
}

// Exercício 1 - Números Primos
func e_primo() {
	var num int
	var contador int
	var i int
	var divisores int

	fmt.Println("Exercício 1 - Números Primos")
	fmt.Println("Informe um número: ")
	fmt.Scanln(&num)

	for i < num {
		i++
		if num%i == 0 {
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
			if num%divisores == 0 {
				fmt.Println(divisores)
			}
		}
	}
}

// Exercício 2 - Sequência de Fibonacci
func fibo() {
	var num int
	i := 1
	a, b := 0, 1

	fmt.Println("Exercício 2 - Sequência de Fibonacci")
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

// Exercício 3 - Dia da Semana
func dia_semana() {
	var num int

	fmt.Println("Exercício 3 - Dia da Semana")
	fmt.Println("Informe um número: ")
	fmt.Scanln(&num)

	switch num {
	case 0:
		fmt.Println("Começa no UM.")
		break
	case 1:
		fmt.Println("Domingo")
		break
	case 2:
		fmt.Println("Segunda-feira")
		break
	case 3:
		fmt.Println("Terça-feira")
		break
	case 4:
		fmt.Println("Quarta-feira")
		break
	case 5:
		fmt.Println("Quinta-feira")
		break
	case 6:
		fmt.Println("Sexta-feira")
		break
	case 7:
		fmt.Println("Sábado")
		break
	default:
		fmt.Println("Erro, dia inválido")
		break
	}
}

// Exercício 4 - Ano Bissexto
func e_bissexto() {
	var ano int

	fmt.Println("Exercício 4 - Ano Bissexto")
	fmt.Println("Informe um ano: ")
	fmt.Scanln(&ano)

	if ano%4 == 0 {
		if ano%100 == 0 && ano%400 != 0 {
			fmt.Println("Não é bissexto")
		} else {
			fmt.Println("É bissexto")
		}
	} else {
		fmt.Println("Não é bissexto")
	}
}
