# Pseudocodigo - Lista Circular Simplesmente Encadeada:

### Declarações
Declarar uma estrutura de nó com valor do contexto em que se trabalha e ponteiro para o próximo nó;
Declarar uma lista com um ponteiro para a cabeça da lista.


### Função para Inicializar a Lista Circular
    Criar um nó "Cabeça" da lista com um valor nulo e o ponteiro que aponta para o próximo nó = cabeça;
       


### Função Exibir()
   Se Cabeça da lista for vazio 
      Retornar que a lista está vazia
   
   Senão
      Enquanto o próximo do nó da vez não for igual ao valor da Cabeça
         Imprimr o Valor do nó
         Avançar para o próximo nó


 

### Função Inserção(Valor)
   Ler o valor a ser inserido na lista e armazenar em um novo nó
   Se a lista estiver vazia
      Declarar o novo nó como a cabeça da lista
      Atualizar o próximo do novo nó para apontar para ele mesmo
   
   Senão
      Percorrer a lista até o nó o qual o Proximo aponte para a cabeça da lista
      Atualizar o próximo do último nó para apontar para o novo nó
      Atualizar o próximo do novo nó para apontar para Cabeça da lista
   




### Função Remoção(Valor)
   Se Cabeça da lista for vazio
      Retornar que não tem o que remover da lista em questão
     
     
   Senão
      Enquanto o próximo do nó não for igual a Cabeça da lista e o Valor do nó não for igual ao Valor a ser removido
         O nó anterior vai receber o valor do nó da vez
         Avançar para o próximo nó
      
         
      Se o Valor foi encontrado
         Se o nó anterior for vazio 
            Atualizar a cabeça da lista para apontar para o próximo nó
            
         Senão
            Atualizar o próximo do nó anterior para apontar para o próximo nó
	 
      
   
