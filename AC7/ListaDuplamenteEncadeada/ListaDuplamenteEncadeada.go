package main

import "fmt"

type NoDuplo struct {
	Valor    int
	Anterior *NoDuplo
	Proximo  *NoDuplo
}

type ListaDuplamenteEncadeada struct {
	Cabeca *NoDuplo
	Cauda  *NoDuplo
}

func NovaListaDuplamenteEncadeada() *ListaDuplamenteEncadeada {
	return &ListaDuplamenteEncadeada{}
}

func (l *ListaDuplamenteEncadeada) InsereValorListaOrdenada(v int) {
	novoNo := &NoDuplo{Valor: v}

	if l.Cabeca == nil {
		l.Cabeca = novoNo
		l.Cauda = novoNo
	} else if v <= l.Cabeca.Valor { //Se for menor que o primeiro valor, é preciso refazer a cabeça da lista;
		novoNo.Proximo = l.Cabeca
		l.Cabeca.Anterior = novoNo
		l.Cabeca = novoNo
	} else if v >= l.Cauda.Valor { // Se for maior que o último valor ja existente na lista;
		novoNo.Anterior = l.Cauda
		l.Cauda.Proximo = novoNo
		l.Cauda = novoNo
	} else { // Se for em algum lugar no meio da lista;
		atual := l.Cabeca
		for atual != nil && v > atual.Valor {
			atual = atual.Proximo
		}

		novoNo.Anterior = atual.Anterior
		novoNo.Proximo = atual
		atual.Anterior.Proximo = novoNo
		atual.Anterior = novoNo
	}
}

func (l *ListaDuplamenteEncadeada) ExibeListaOrdenada() {
	if l.Cabeca == nil {
		fmt.Println("A lista está vazia")
		return
	}

	no := l.Cabeca
	for no != nil {
		fmt.Println(no.Valor)
		no = no.Proximo
	}
	fmt.Println()
}

func (l *ListaDuplamenteEncadeada) RemoveValorListaOrdenada(v int) {
	if l.Cabeca == nil {  //Se a lista estiver vazia;
		fmt.Println("A lista está vazia, não tem o que remover.")
		return
	}


	if l.Cabeca.Valor == v {  //Se o valor estiver na cabeça, removerá aqui;
		l.Cabeca = l.Cabeca.Proximo
		if l.Cabeca != nil {
			l.Cabeca.Anterior = nil
		}
		return
	}

	if l.Cauda.Valor == v { // Se o valor estiver na cauda, removerá aqui;
		l.Cauda = l.Cauda.Anterior
		if l.Cauda != nil {
			l.Cauda.Proximo = nil
		}
		return
	}

	atual := l.Cabeca
	for atual != nil && atual.Valor != v {
		atual = atual.Proximo
	}

	if atual != nil {
		atual.Anterior.Proximo = atual.Proximo
		atual.Proximo.Anterior = atual.Anterior
	}
}



func main() {
	lista := NovaListaDuplamenteEncadeada()

	listaTeste.ExibeListaCircular()
	listaTeste.InsereValorListaCircular(4)
	listaTeste.ExibeListaCircular()
	listaTeste.InsereValorListaCircular(2)
	listaTeste.ExibeListaCircular()
	listaTeste.InsereValorListaCircular(9)
	listaTeste.ExibeListaCircular()
	listaTeste.RemoveValorListaCircular(4)  // tirar um do meio;
	listaTeste.ExibeListaCircular()
	listaTeste.RemoveValorListaCircular(99) // tirar valor que não existe na lista;
	listaTeste.ExibeListaCircular()
	listaTeste.InsereValorListaCircular(5)  // adicionando um 5;
	listaTeste.ExibeListaCircular()
	listaTeste.RemoveValorListaCircular(2) // tirar um da cabeça;
	listaTeste.ExibeListaCircular()
	listaTeste.RemoveValorListaCircular(9) // tirar um do final;
	fmt.Println("Sucesso.")
