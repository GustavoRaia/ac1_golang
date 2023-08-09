package main

import "fmt"

func main() {
	dia_semana()
}

func dia_semana() {
	var num int

	fmt.Println("Informe um número: ")
	fmt.Scanln(&num)

	switch num {
	case 0:
		fmt.Println("Começa no UM.")
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda-feira")
	case 3:
		fmt.Println("Terça-feira")
	case 4:
		fmt.Println("Quarta-feira")
	case 5:
		fmt.Println("Quinta-feira")
	case 6:
		fmt.Println("Sexta-feira")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Erro, dia inválido")
	}
}