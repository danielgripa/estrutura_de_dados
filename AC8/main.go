//Grupo: Daniel Gripa e Thalles Diepes

package main

import "fmt"

type No struct {
	valor int
	esq   *No
	dir   *No
}

type ArvoreBinaria struct {
	raiz *No
}

func (arvore *ArvoreBinaria) insereValor(valor int) string {
	novoNo := &No{valor: valor}

	if arvore.raiz == nil {
		arvore.raiz = novoNo
		return fmt.Sprintf("Valor %d inserido como raiz da árvore.", valor)
	}

	noAtual := arvore.raiz
	for {
		if valor < noAtual.valor {
			if noAtual.esq == nil {
				noAtual.esq = novoNo
				return fmt.Sprintf("Valor %d inserido à esquerda de %d.", valor, noAtual.valor)
			}
			noAtual = noAtual.esq
		} else if valor > noAtual.valor {
			if noAtual.dir == nil {
				noAtual.dir = novoNo
				return fmt.Sprintf("Valor %d inserido à direita de %d.", valor, noAtual.valor)
			}
			noAtual = noAtual.dir
		} else {
			return fmt.Sprintf("Inserção inválida: valor %d já existe na árvore.", valor)
		}
	}
}

func (arvore *ArvoreBinaria) buscaValor(valor int, no *No) (*No, bool) {
	if no == nil || no.valor == valor {
		return no, no != nil
	}

	if valor < no.valor {
		return arvore.buscaValor(valor, no.esq)
	}
	return arvore.buscaValor(valor, no.dir)
}

func main() {
	arvore := ArvoreBinaria{}
	valoresTeste := []int{99, 4, 98, 30, 97, 18} // não tem, tem, não tem, tem, etc...

	fmt.Println(arvore.insereValor(13))
	fmt.Println(arvore.insereValor(25))
	fmt.Println(arvore.insereValor(8))
	fmt.Println(arvore.insereValor(17))
	fmt.Println(arvore.insereValor(2))
	fmt.Println(arvore.insereValor(12))
	fmt.Println(arvore.insereValor(22))
	fmt.Println(arvore.insereValor(18))
	fmt.Println(arvore.insereValor(30))
	fmt.Println(arvore.insereValor(4))
	fmt.Println(arvore.insereValor(11))
	fmt.Println(arvore.insereValor(19))
	fmt.Println(arvore.insereValor(27))

	for _, valor := range valoresTeste {
		_, encontrado := arvore.buscaValor(valor, arvore.raiz) //Se encontrado for true, encontrou. Não necessitando do valor em si para guardar aqui, por isso ignorei o valor em si com "_".
		if encontrado {
			fmt.Println("Valor", valor, "encontrado")
		} else {
			fmt.Println("Não foi posssível achar o valor:", valor)
		}
	}
}
