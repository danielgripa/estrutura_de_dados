package main

import (
	"fmt"
)

//Elabore uma função e_primo() que retorna se um número é primo ou não.
// Caso o número não seja primo, liste todos os números pelos quais ele é divisível.


func ePrimo(valor int) (string, []int) {
	var primo string
	var divisores [] int

	for i :=2; i < valor; i++{
		if valor % i == 0 {
			divisores = append(divisores,i)
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
// 1, 1, 2, 3, 5, 8
package main

import "fmt"

func fibo(n int) int {
    if n <= 0 {
        return 0
    } else if n == 1 {
        return 1
    }

    anterior := 0
    vigente := 1
    for i := 2; i <= n; i++ {
        proximo := prev + current
        anterior = vigente
        vigente = proximo

    return vigente
}



func main() {
	fmt.Println(ePrimo(225))
	fmt.Println(fibo(16))
}





