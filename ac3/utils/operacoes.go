package utils

import(
	"fmt"
	"ac3/ac3/modelo"
	"strconv"
)

func AdicionaContato(c modelo.Contato, listaContatos *[5]modelo.Contato) {
	for ind, contato := range listaContatos {
		if (contato == modelo.Contato{}) {
			listaContatos[ind] = c
			break
		}
	}
}

func ExcluiContato(listaContatos *[5]modelo.Contato) {
	if (listaContatos[0] == modelo.Contato{}) {
		fmt.Println("Lista de contatos vazia!")
	}

	for ind, contato := range listaContatos {
		if (contato == modelo.Contato{}) {
			listaContatos[ind-1] = modelo.Contato{}
		}
	}
}

func AlterarEmail(c *modelo.Contato, novoEmail string) {
	c.Email = novoEmail
}

func EditaEmail (listaContatos *[5]modelo.Contato,indice int) {
	// For não percorre o array de Contatos
	for i:=0; i<5; i++ {
		/* VERIFICAÇÃO PARA INDEX NÃO CORRESPONDENTE

		if (contato == modelo.Contato{}) {
			fmt.Println("Lista de contatos vazia!")
		}
		*/
		if (i == indice) {
			var novoEmail string

			fmt.Println("Informe um novo nome de Email: ")
			fmt.Scanln(&novoEmail)

			AlterarEmail(/*Ponteiro de contato*/ &listaContatos[indice], novoEmail)

			// listaContatos[indice].Email = novoEmail // Funcionou sem a função
		}
	}
}

func MostraContatos(listaContatos *[5]modelo.Contato) {
	for i:= 0; i<5; i++{
		if (listaContatos[0] == modelo.Contato{}) {
			fmt.Println("\nLISTA DE CONTATOS VAZIA!")
			break
		} else if (listaContatos[i] == modelo.Contato{}) {
			break
		} else {
			fmt.Println("\nContato nº " + strconv.Itoa(i+1))
			fmt.Println("Nome: " + listaContatos[i].Nome)
			fmt.Println("Email: " + listaContatos[i].Email)
			fmt.Println("\n------------------------------")
		}
	}
}