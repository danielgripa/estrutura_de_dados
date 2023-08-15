package main

import (
	"fmt"
)

// Questão 1
func ePrimo(valor int) (string, []int) {
	var primo string
	var divisores []int

	for i := 2; i < valor; i++ {
		if valor%i == 0 {
			divisores = append(divisores, i)
		}
	}
	if len(divisores) == 0 {
		primo = "Ele é primo"
	} else {
		primo = "Não é primo, ele é divisível por"
	}
	return primo, divisores
}

// Questão 2
func fibo(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	anterior := 0
	vigente := 1
	for i := 2; i <= n; i++ {
		proximo := anterior + vigente
		anterior = vigente
		vigente = proximo
	}

	return vigente
}

// Questão 3
func diaDaSemana(valor int) string {
	var dia string
	switch valor {
	case 1:
		dia = "Domingo"
	case 2:
		dia = "Segunda-Feira"
	case 3:
		dia = "Terça-Feira"
	case 4:
		dia = "Quarta-Feira"
	case 5:
		dia = "Quinta-Feira"
	case 6:
		dia = "Sexta-Feira"
	case 7:
		dia = "Sábado"
	default:
		dia = "Valor inválido"
	}
	return dia
}

// Questão 4
func eBissexto(valor int) string {
	var resposta string
	if valor%4 == 0 && valor%100 != 0 || valor%400 == 0 {
		resposta = "É Bissexto"
	} else {
		resposta = "Não é bissexto"
	}
	return resposta
}

//Testando as funções das questões abaixo
func main() {
	fmt.Println(ePrimo(225))
	fmt.Println(fibo(12))
	fmt.Println(diaDaSemana(3))
	fmt.Println(eBissexto(1640))
}
