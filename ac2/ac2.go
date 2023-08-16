package main

import "fmt"

type Contato struct {
	Nome	string
	Email	string
}

func main() {
	var listaContatos [5]Contato
	var opcao int

	fmt.Println("Escolha uma opção: ")
	fmt.Println("Opçao 1: Adicionar Contato")
	fmt.Println("Opçao 2: Excluir Contato")
	fmt.Println("Opçao 3: Mostrar Contatos")
	fmt.Scanln(&opcao)

	// Alimentando os contatos no Array
	// listaContatos[0] = Contato{"Gustavo", "gustavo@example.com"}
	// listaContatos[1] = Contato{"Azevedo", "azevedo@example.com"}
	// listaContatos[2] = Contato{"Raia", "raia@example.com"}
	// listaContatos[3] = Contato{"de", "de@example.com"}
	// listaContatos[4] = Contato{"Siqueira", "siqueira@example.com"}

	switch opcao {
	case 1:
		fmt.Println("\nAdicionar Contatos")
		listaContatos = adicionaContato(listaContatos)
	case 2:
		fmt.Println("\nExcluir Contatos")
		listaContatos = excluiContato(listaContatos)
	case 3:
		fmt.Println("\nMostrar Contatos")
		fmt.Println(listaContatos)
	default:
		fmt.Println("Fechando")
	}
}

func adicionaContato (c [5]Contato) [5]Contato {
	var i int
	for _, contatos := range c {

		if (contatos == Contato{}) {
			var nome string
			var email string

			fmt.Println("Informe um nome: ")
			fmt.Scanln(&nome)
			c[i] = Contato{Nome: nome}

			// Não lê o segundo Scan

			fmt.Println("Informe um email: ")
			fmt.Scanln(&email)
			c[i] = Contato{Email: email}

			fmt.Println(c[i].Nome)

			i++
			break
		} else {
			fmt.Println("TODOS USUÁRIOS CADASTRADOS")
			break
		}
	}
	fmt.Println(c)
	return c
}

func excluiContato(c [5]Contato) [5]Contato {
	var contato Contato
	i:=4 // Contador
	for i >= 0 {
		if (contato != Contato{}) {
			c[i] = Contato{Nome: ""}
			c[i] = Contato{Email: ""}
			break
		} else {
			fmt.Println("Contato Vazio")
		}
		i--
	}
	fmt.Println(c)

	return c
}