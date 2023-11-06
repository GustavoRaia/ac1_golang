package main

import (
	"fmt"
)

type No struct {
	Chave    int
	Esq, Dir *No
}

type Arvore struct {
	Raiz *No
}

func main() {
	arvore := Arvore{}

	n1 := &No{Chave: 2}
	n2 := &No{Chave: 5}
	n3 := &No{Chave: 1}
	n4 := &No{Chave: 9}
	n5 := &No{Chave: 0}
	n6 := &No{Chave: 7}
	n7 := &No{Chave: 8}

	arvore.Raiz = n1
	n1.Esq = n2
	n2.Esq = n3
	n2.Dir = n4
	n1.Dir = n5
	n4.Dir = n6
	n6.Esq = n7

	fmt.Println("Pré Ordem")
	fmt.Println("")

	imprimePreOrdem(arvore.Raiz)

	fmt.Println("")
	fmt.Println("Simétrico")
	fmt.Println("")

	imprimeSimetrico(arvore.Raiz)

	fmt.Println("")
	fmt.Println("Pós Ordem")
	fmt.Println("")

	imprimePosOrdem(arvore.Raiz)
}

func buscaArvore(valor int, no *No) int {
	if no == nil {
		return 0
	}

	if valor == no.Chave {
		return 1
	}

	if valor < no.Chave {
		if no.Esq == nil {
			return 2
		}
		return buscaArvore(valor, no.Esq)
	}

	if valor > no.Chave {
		if no.Dir == nil {
			return 3
		}
		return buscaArvore(valor, no.Dir)
	}

	return -1
}

func insere(valor int, arvore *Arvore) {
	no := arvore.Raiz
	f := buscaArvore(valor, no)

	if f == 1 {
		fmt.Println("Inserção inválida")
		return
	}

	novoNo := &No{Chave: valor, Esq: nil, Dir: nil}

	if f == 0 {
		arvore.Raiz = novoNo
	} else if f == 2 {
		no.Esq = novoNo
	} else {
		no.Dir = novoNo
	}
}

func imprimePreOrdem(no *No) {
	if no != nil {
		fmt.Println(no.Chave)
		imprimePreOrdem(no.Esq)
		imprimePreOrdem(no.Dir)
	}
}

func imprimeSimetrico(no *No) {
	if no != nil {
		imprimeSimetrico(no.Esq)
		fmt.Println(no.Chave)
		imprimeSimetrico(no.Dir)
	}
}

func imprimePosOrdem(no *No) {
	if no != nil {
		imprimePosOrdem(no.Esq)
		imprimePosOrdem(no.Dir)
		fmt.Println(no.Chave)
	}
}
