package main

import (
	"fmt"
)

type No struct {
	valor int
	anterior *No
	proximo *No
}

type Lista struct {
	tamanho int
	cabeca *No
}

func main() {
	lista := &Lista{cabeca: &No{}}

	// Adicionando números à lista.
	lista.adicionaLista(9, 1)
	lista.adicionaLista(4, 2)
	lista.adicionaLista(5, 3)
	lista.adicionaLista(3, 4)
	lista.exibeLista()

	fmt.Println("")

	// Removendo o 2º nó da Lista.
	lista.removeLista(2)
	lista.exibeLista()

	fmt.Println("")

	// Busca um nó na Lista pelo VALOR.
	resultado := lista.buscaLista(9)
	if(resultado != nil) {
		fmt.Println(resultado.valor)
	} else {
		fmt.Println("Não encontrado.")
	}

	lista.removeLista(1)
	fmt.Println("")
	
	lista.exibeLista()
	fmt.Println("")

	// Busca um nó na Lista pelo VALOR.
	resultado = lista.buscaLista(9)
	if(resultado != nil) {
		fmt.Println(resultado.valor)
	} else {
		fmt.Println("Não encontrado.")
	}
}

// Exibição dos nós em uma lista duplamente encadeada.
func (l *Lista) exibeLista() {
	no := l.cabeca.proximo
	
	for(no != nil) {
		fmt.Println(no.valor)
		no = no.proximo
	}	
}

// Busca de um Nó.
func (l *Lista) buscaLista(n int) *No {
	no := l.cabeca.proximo

	for(no != nil) {
		if(no.valor == n) {
			return no
		}
		no = no.proximo
	}

	return nil
}

// Inserção de um Nó.
func (l *Lista) adicionaLista(num, posicao int) {
	temp := &No{valor: num}
	p := l.cabeca

	for i := 0; i < posicao-1; i++ {
		p = p.proximo
	}

	temp.proximo = p.proximo

	// Verifica se o nó não é o último da lista
	if(p.proximo != nil) {
		p.proximo.anterior = temp
	}
	
	temp.anterior = p
	p.proximo = temp

	l.tamanho++
}

// Remoção de um Nó.
func (l *Lista) removeLista(posicao int) {
	p := l.cabeca

	for i := 0; i < posicao-1; i++ {
		p = p.proximo
	}

	p.proximo = p.proximo.proximo

	// Verifica se o nó não é o último da lista
	if(p.proximo != nil) {
		p.proximo.anterior = p
	}

	l.tamanho--
}
