package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var contatos [5]Contato

	for {
		fmt.Println("Selecione uma operação através do seu teclado numérico:")
		fmt.Println("1 - Adicionar contato")
		fmt.Println("2 - Excluir último contato")

		var opcao int
		fmt.Print("Opção: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			fmt.Print("Digite o nome completo do contato: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			nome := scanner.Text()
			fmt.Print("Digite o e-mail do contato: ")
			scanner.Scan()
			email := scanner.Text()

			novoContato := contatos.Contato{Nome: nome, Email: email}
			contatos = utils.adicionaContato(novoContato, contatos)

		case 2:
			contatos = utils.excluiContato(contatos)

		default:
			fmt.Println("Atenção: Selecione uma opção numérica conforme o índice acima.")
		}
	}
}
