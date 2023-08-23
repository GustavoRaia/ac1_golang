package main

import (
	"ac3/ac3/modelo"
	"ac3/ac3/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Declarando variáveis
	var contatos [5]modelo.Contato
	var nome, email, opcao string
	var indice int

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("(1) para adicionar" +
				"\n(2) para remover" +
				"\n(3) para editar email" +
				"\n(4) para exibir todos os contatos salvos" +
				"\nou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			c := modelo.Contato{
				Nome:	nome,
				Email: email,
			}

			utils.AdicionaContato(c, &contatos)
		case "2":
			utils.ExcluiContato(&contatos)
		case "3":
			fmt.Print("Informe o índice do Email (0 até 4): ")
			fmt.Scanln(&indice)

			utils.EditaEmail(&contatos, indice)
		case "4":
			utils.MostraContatos(&contatos)
		default:
			return
		}
	}
}