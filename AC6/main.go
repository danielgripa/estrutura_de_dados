package main

import "fmt"

const M = 5

func main() {
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

func insereOrd(v *[M]int, n *int, novoValor int) {
	fmt.Println("Valor de n é: ", *n)
	fmt.Println("tentando inserir ", novoValor)

	if *n >= M {
		fmt.Println("overflow")
		return
	} else {
		posicao := 0
		for posicao < *n && v[posicao] < novoValor {
			posicao++
		}

		for i := *n; i > posicao; i-- {
			v[i] = v[i-1]
		}

		v[posicao] = novoValor

		(*n)++
	}
}
