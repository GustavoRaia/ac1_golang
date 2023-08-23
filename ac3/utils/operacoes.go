package utils

import(
	"fmt"
	"ac3/modelo"
	"strconv"
)

// Função para adicionar os contatos ao Array
func AdicionaContato(c modelo.Contato, listaContatos *[5]modelo.Contato) {

	// Percorre o Array de contatos
	for ind, contato := range listaContatos {

		// Verifica se o último elemento do Array está preenchido
		if (listaContatos[4] != modelo.Contato{}) {
			// Se sim, envia uma mensagem de erro e para a execução do loop
			fmt.Println("\n == LISTA DE CONTATOS CHEIA ==")
			break
		}

		// Verifica se o contato de indice atual é vazio
		if (contato == modelo.Contato{}) {
			// Se sim, cadastra um novo contato
			listaContatos[ind] = c
			break
		}
	}
}

// Função de excluir último contato cadastrado da lista
func ExcluiContato(listaContatos *[5]modelo.Contato) {
	for i := 0; i < len(listaContatos); i++  {
		// Verifica se o primeiro elemento é um Contato vazio
		if (listaContatos[0] == modelo.Contato{}) {
			// Se for, envia uma mensagem de erro e termina a execução da função
			fmt.Println("\n == LISTA DE CONTATOS VAZIA ==")
			break
		}
		
		// Verifica se o último contato da lista não é vazio
		if (listaContatos[4] != modelo.Contato{}) {
			// Se não for(case true), exclui o cadastro e encerra a execução do loop
			listaContatos[4] = modelo.Contato{}
			fmt.Println("\nContato de nº5 apagado.")
			break
		}
		
		// Verifica se o contato de indice atual é vazio
		if (listaContatos[i] == modelo.Contato{}) {
			// Exclui o último contato cadastrado da lista
			listaContatos[i-1] = modelo.Contato{}
			fmt.Println("\nContato de nº" + strconv.Itoa(i) + " apagado.")
		}
	}
}

// Função que auxilia a func EditarEmail()
func AlterarEmail(c *modelo.Contato, novoEmail string) {
	c.Email = novoEmail
}

// Função de Editar Email
func EditaEmail (listaContatos *[5]modelo.Contato,indice int) {

	// For não percorre o array de Contatos
	for i := 0; i < len(listaContatos); i++ {

		// Verifica se o primeiro elemento é um Contato vazio
		if (listaContatos[0] == modelo.Contato{}) {
			// Se for, envia uma mensagem de erro e termina a execução da função
			fmt.Println("\n == LISTA DE CONTATOS VAZIA ==")
			break
		}

		if (indice>4){
			fmt.Println("\nDigite números de 1 e 5")
			break
		}

		// Verifica se o índice fornecido é igual ao valor atual de i
		if (i == indice) {
			
			// Declara a variável novoEmail
			var novoEmail string

			// Pede ao usuário um novo valor
			fmt.Println("Informe um novo nome de Email: ")
			fmt.Scanln(&novoEmail)

			// Verifica se o email novo é igual ao antigo
			if(novoEmail == listaContatos[indice].Email){
				// Se sim, envia uma mensagem alarmando o usuário
				fmt.Println("\nO Email já é esse!")
			} else {
				// Se não, será alterado
				AlterarEmail(/*Ponteiro de contato*/ &listaContatos[indice], novoEmail)
				fmt.Println("\nEmail alterado com sucesso!")
			}
		}
	}
}

// Função de Exibir Contatos
func MostraContatos(listaContatos *[5]modelo.Contato) {

	// Percorre a lista de Contatos
	for i := 0; i < len(listaContatos); i++{

		if (listaContatos[0] == modelo.Contato{}) {
			// Quebra se o primeiro contato for vazio e mostra uma mensagem
			fmt.Println("\n == LISTA DE CONTATOS VAZIA == ")
			break

		} else if (listaContatos[i] == modelo.Contato{}) {
			// Quebra / Não Exibe quando é um Contato Vazio
			break
			
		} else {
			// Printa os atributos do contato no indice(i) atual da lista
			fmt.Println("\n------------------------------")
			fmt.Println("\nContato nº " + strconv.Itoa(i+1))
			fmt.Print("\nNome: " + listaContatos[i].Nome)
			fmt.Println("Email: " + listaContatos[i].Email)
		}
	}
}