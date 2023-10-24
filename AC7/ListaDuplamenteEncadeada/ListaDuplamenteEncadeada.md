# Pseudocodigo - Lista Duplamente Encadeada:

### Declarações
Declarar uma estrutura de nó com valor do contexto em que se trabalha e dois ponteiros: um para apontar o próximo nó e outro para o anterior;
Declarar uma lista com um ponteiro para a cabeça da lista.


### Função InicializarListaDuplamenteEncadeada():
    Lista recebe Nova ListaDuplamenteEncadeada
    Lista.Inicial recebe Nulo
    Lista.Final recebe Nulo

    Retorne Lista

### Função InserirListaDuplamenteEncadeada(Lista, Valor):
    NovoNó recebe Novo Nó
    NovoNó.Valor recebe Valor

    Se Lista.Tamanho = 0:
        Lista.Inicial recebe NovoNó
        Lista.Final recebe NovoNó
    Senão:
        NovoNó.Anterior recebe Lista.Final
        Lista.Final.Próximo recebe NovoNó
        Lista.Final recebe NovoNó


### Função RemoverListaDuplamenteEncadeada(Lista):
    Se Lista.Tamanho = 0:
        Retorne Erro("A lista está vazia")
    Senão:
        PrimeiroNó recebe Lista.Inicial
        Lista.Inicial recebe PrimeiroNó.Próximo

        Se Lista.Inicial for Nulo:
            Lista.Final recebe Nulo
        Senão:
            Lista.Inicial.Anterior recebe Nulo

        Retorne PrimeiroNó.Valor

### Função ExibirListaDuplamenteEncadeada(Lista):
    NóAtual recebe Lista.Inicial
    Enquanto NóAtual for diferente de Nulo:
        Imprima NóAtual.Valor
        NóAtual recebe NóAtual.Próximo