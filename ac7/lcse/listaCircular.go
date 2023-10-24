package main

import (
	"fmt"
)

type No struct {
	valor int
	proximo *No
}

type Lista struct {
	inicio *No
	fim *No
}

func main() {
	lista := &Lista{}

	lista.adicionaLista(1)
	lista.adicionaLista(5)
	lista.adicionaLista(4)
	lista.adicionaLista(9)
	lista.exibeLista()

	fmt.Println("")

	lista.removePrimeiroNo()
	lista.exibeLista()

	fmt.Println("")

	lista.removePrimeiroNo()
	lista.exibeLista()

	fmt.Println("")

	lista.adicionaLista(14)
	lista.exibeLista()
}

// Exibição dos Nós em uma Lista Circular

// Função para exibir os números de uma lista circular simplesmente encadeada
func (l *Lista) exibeLista() {
	no := l.inicio

	for(no != nil) {
		fmt.Println(no.valor)
		no = no.proximo
	}
}

// Inserção de um Nó no início da Lista
func (l *Lista) adicionaLista(num int) {
	no := No{valor: num}

	if(l.inicio == nil) {
		l.inicio = &no

	}
	if(l.fim != nil) {
		l.fim.proximo = &no
	}

	l.fim = &no
}


// Exclusão do Primeiro Nó da Lista
func (l *Lista) removePrimeiroNo() {

	// Verifica se o primeiro item da lista é vazio (lista vazia)
    if(l.inicio == nil) {
        fmt.Println("Lista vazia.")
        return
    }

    l.inicio = l.inicio.proximo
}