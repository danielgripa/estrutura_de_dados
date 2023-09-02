package torrehanoi

import (
	"fmt"
	"math"
)

func Hanoi(discos uint8) uint64 {
	if discos == 1 {
		fmt.Println("Uma jogada é suficiente. Apenas movendo o disco da primeira torre para a última.")
		return 1
	} else if discos > 63 || discos < 0 {
		fmt.Println("Só é possível calcular em jogos de 1 até 63 discos.")
		return 0
	}
	
	return uint64(math.Pow(2,float64(discos)))-1

}
