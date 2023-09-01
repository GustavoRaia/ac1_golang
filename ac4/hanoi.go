package main

import (
	"fmt"
	"strconv"
)

func main() {
	var discos int

	// Pede a quantidade de discos desejado
	fmt.Println("\nDigite a quantidade de discos: ")
	fmt.Scanln(&discos) // Lê a quantidade
	fmt.Println("")

	movimentos := hanoi(discos, "A", "C", "B")

	fmt.Println("\nMovimentos necessários: " + strconv.Itoa(movimentos))
	fmt.Println("")
}

func hanoi(n int, origem, destino, auxiliar string) int {
	movimentos := 0

	if n == 1 {
		fmt.Println("Disco " + strconv.Itoa(n) + ": " + origem + " → " + destino)
		return 1
	}

	// Move da torre de origem para a torre auxiliar usando a torre de destino como auxiliar
	// n ← n-1
	// Origem ← Origem
	// Destino ← Auxiliar
	// Auxiliar ← Origem
	movimentos += hanoi(n-1, origem, auxiliar, destino)

	// Move o disco restante da torre de origem para a torre de destino
	fmt.Println("Disco " + strconv.Itoa(n) + ": " + origem + " → " + destino)
	movimentos++

	// Move da torre auxiliar para a torre destino usando a torre de origem como auxiliar
	// n ← n-1
	// Origem ← auxiliar
	// Destino ← Destino
	// Auxiliar ← Origem
	movimentos += hanoi(n-1, auxiliar, destino, origem)

	return movimentos
}