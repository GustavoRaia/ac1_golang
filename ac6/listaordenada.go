package main

import (
	"fmt"
	"sort"
	"strconv"
)

const M = 5

func main() {
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

// Verifica se o valor procurado já esta na lista
func busca1(v [M]int, n, x int) int {
	i := 0
	for i < n {
		if v[i] == x {
			return i
		}
		i++
	}
	return -1
}

func insereOrd(v *[M]int, n *int, novoValor int) {
	if(*n == M) { // Verifica se o valor de N é o mesmo do tamanho da lista
		fmt.Println("\nOverflow") // Mensagem de Erro
	} else { // Próximo bloco de instruções
		fmt.Println("\nTentando inserir " + strconv.Itoa(novoValor))
		ind := busca1(*v, *n, novoValor) // Chama função que retorna -1 caso novoValor não esteja inserido na lista

		if(ind == -1) { // Se o valor procurado não estiver na lista
			fmt.Println(strconv.Itoa(novoValor) + " não encontrado, inserindo na lista")

			// BLOCO PARA ORDENAR A LISTA DE FORMA CRESCENTE

			/* Tentativa de fazer funcionar
			for i := 0; i < *n; i++ {
				fmt.Println(v[i])

				if novoValor > v[i] && novoValor > v[i+1] {
					// verifica se é maior que os 2 valores
					n1 := v[i]
					n2 := v[i+1]

					fmt.Println(n1)
					fmt.Println(n2)

				} else if novoValor > v[i] && novoValor < v[i+1] {
					// verifica se está entre os 2 valores
					continue
				} else {
					v[i] = novoValor
				}
			}
			*/

			v[*n] = novoValor // Atribui o novo valor a uma casa na lista (indice *n)
			*n++ // Incrementa para que seja inserido na próxima casa

			slice := v[:*n] // Copia o valor da lista para um slice novo
			sort.Ints(slice) // Utiliza a biblioteca sort para ordenar os itens

		} else { // Se o valor procurado já estiver na lista
			fmt.Println("\nElemento já existe na tabela")
		}
	}
}