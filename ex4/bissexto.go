package main

import "fmt"

func main() {
	e_bissexto()
}

func e_bissexto() {
	var ano int

	fmt.Println("Informe um ano: ")
	fmt.Scanln(&ano)

	if (ano % 4 == 0){
		if ano % 100 == 0 && ano % 400 != 0{
			fmt.Println("Não é bissexto")
		} else {
			fmt.Println("É bissexto")
		}
	} else {
		fmt.Println("Não é bissexto")
	}
}