// Função realizada originalmente em sala de aula
// ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓

/*
package main

import (
	"fmt"
)

func main() {
	numeros := []int{1, 2, 4, 6, 3, 5, 7, 12, 18}
	fmt.Println(achaSoma(numeros, 15))
	fmt.Println(achaSoma(numeros, 32))
	fmt.Println(achaSoma(numeros, 16))
	fmt.Println(achaSoma(numeros, 3))
	fmt.Println(achaSoma(numeros, 8))
	fmt.Println(achaSoma(numeros, 26))
	fmt.Println(achaSoma(numeros, 5))
	fmt.Println(achaSoma(numeros, 30))
}

func achaSoma(lista []int, alvo int) (int, int) {
	i:=0
	for(i < len(lista)){
		j := i+1
		for(j < len(lista)) {
			soma := lista[i]+lista[j]
			if(soma == alvo) {
				return lista[i], lista[j]
			}
			j++
		}
		i++
	}
	return -1, -1
}
*/

// ----------------------------------------------

package main

import (
	"fmt"
)

func achaSoma(lista []int, alvo int) (int, int) {
    num1 := 0 // Primeiro número recebe primeiro item da lista
	num2 := len(lista)-1 // Segundo número recebe o último item da lista

    for num1 < num2 { // Verifica se o primeiro item da lista é menor que o último
        soma := lista[num1] + lista[num2] // Define a variável soma para guardar o valor dos 2 itens 

        if soma == alvo {
            return lista[num1], lista[num2] // Se o alvo é encontrado (igual ao valor da soma), retorna os valores da soma
        } else if soma < alvo {
            num1++ // Se a soma é menor que o alvo, o valor menor (à esquerda) vai para a próxima casa à direita
        } else {
            num2-- // Se a soma é maior que o alvo, o maior valor (à direita) vai para a casa anterior (à esquerda)
        }
    }
    return -1, -1 // Se a soma não for igual ao alvo durante a execução do array não haverá retorno, logo, irá retornar (-1,-1)
}

func main() {
	// numeros := []int{1, 2, 4, 6, 3, 5, 7, 12, 18} // Array não crescente para testes
    numeros := []int{1, 2, 4, 6, 7, 12, 18} // Array crescente
	// alvo := 25 // Valor alvo a ser encontrado (por soma dos itens) dentro do Array - OPCIONAL

    fmt.Println(achaSoma(numeros, 1))
    fmt.Println(achaSoma(numeros, 5))
    fmt.Println(achaSoma(numeros, 13))
    fmt.Println(achaSoma(numeros, 15))
    fmt.Println(achaSoma(numeros, 22))
    fmt.Println(achaSoma(numeros, 30))
	/*
    fmt.Println(achaSoma(numeros, 4))
    fmt.Println(achaSoma(numeros, 5))
    fmt.Println(achaSoma(numeros, 6))
    fmt.Println(achaSoma(numeros, 7))
    fmt.Println(achaSoma(numeros, 8))
    fmt.Println(achaSoma(numeros, 9))
    fmt.Println(achaSoma(numeros, 10))
    fmt.Println(achaSoma(numeros, 11))
    fmt.Println(achaSoma(numeros, 12))
    fmt.Println(achaSoma(numeros, 13))
    fmt.Println(achaSoma(numeros, 14))
    fmt.Println(achaSoma(numeros, 15))
    fmt.Println(achaSoma(numeros, 16))
    fmt.Println(achaSoma(numeros, 17))
    fmt.Println(achaSoma(numeros, 18))
    fmt.Println(achaSoma(numeros, 19))
    fmt.Println(achaSoma(numeros, 20))
    fmt.Println(achaSoma(numeros, 21))
    fmt.Println(achaSoma(numeros, 22))
    fmt.Println(achaSoma(numeros, 23))
    fmt.Println(achaSoma(numeros, 24))
    fmt.Println(achaSoma(numeros, 25))
    fmt.Println(achaSoma(numeros, 26))
    fmt.Println(achaSoma(numeros, 27))
    fmt.Println(achaSoma(numeros, 28))
    fmt.Println(achaSoma(numeros, 29))
    fmt.Println(achaSoma(numeros, 30))
	*/
}

// Possibilidade de somas para o Array crescente
// 3, 5, 6, 7, 8, 9,
// 10, 11, 13, 14, 16, 18, 19,
// 20, 22, 24, 25,
// 30