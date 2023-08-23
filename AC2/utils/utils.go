package utils

import (
	"contatos"
	"fmt"
)

func adicionaContato(contato contatos.Contato, arrayContatos []contatos.Contato) []contatos.Contato {
	for i := range arrayContatos {
		if arrayContatos[i].Nome == "" && arrayContatos[i].Email == "" {
			arrayContatos[i] = contato
			fmt.Println("Inclusão realizada.")
			break
		} else {
			fmt.Println("Não há espaço disponível para cadastrar mais contato.")
		}
	}

	return arrayContatos
}

func excluiContato(arrayContatos []contatos.Contato) []contatos.Contato {
	for i := len(arrayContatos) - 1; i >= 0; i-- {
		if arrayContatos[i].Nome != "" && arrayContatos[i].Email != "" {
			arrayContatos[i] = contatos.Contato{}
			fmt.Println("Exclusão realizada com sucesso.")
			break

		} else {
			fmt.Println("Não há mais contatos para excluir.")
		}
	}

	return arrayContatos
}
