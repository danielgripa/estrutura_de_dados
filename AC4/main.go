package main

import (
	"fmt"
	"torrehanoi/torrehanoi"
)

func main() {
	var discos uint8
	fmt.Print("Escolha um número de até 63 discos que você quer jogar?")
	fmt.Scan(&discos)

	jogadasMinimas := torrehanoi.Hanoi(discos)

	if jogadasMinimas == 0 {
		fmt.Println("Reinicie o cálculo")
	}
	fmt.Println("São ",jogadasMinimas, " jogos possíveis")

}
