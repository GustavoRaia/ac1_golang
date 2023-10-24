# Lista Linear Duplamente Encadeada

Numa lista linear duplamente encadeada cada elemento possui, além do espaço para armazenamento da informação, um espaço para armazenar a referencia da localização de memória onde se encontra o próximo elemento da lista e outro espaço para armazenar a referencia da localização de memória onde se encontra o elemento anterior.

## Exibição dos Nós em uma Lista Duplamente Encadeada

1 - Variável nó recebe o valor do próximo elemento da lista
2 - Para cada nó diferente de NULO:
        Exibe na tela o valor do Nó.
        Atribui à variável nó o valor do elemento seguinte

## Busca de um Nó

1 - Variável nó recebe o valor do próximo elemento da lista
2 - Para cada nó diferente de NULO:
        Se valor do nó for igual ao valor buscado:
            retorna Nó
        Nó recebe o elemento seguinte
3 - Retorna NULO

## Inserção de um Nó

1 - Variável temporária recebe o valor do ponteiro de um Nó (com valor a ser inserido)
2 - Variável de posição recebe o valor da cabeça da Lista
3 - Para cada elemento da Lista:
        Atribui o próximo valor à cabeça da Lista
4 - Se o proximo elemento não for NULO:
        Atribui ao valor anterior o valor da cabeça da Lista
5 - Atribui o valor da cabeça da Lista ao valor anterior da variável temporária
6 - Atribui ao valor seguinte à cabeça da Lista o valor da variável temporária
7 - Incrementa o tamnho da lista em 1

## Remoção de um Nó

1 - Variável de posição recebe o valor da cabeça da Lista
2 - Para cada elemento da Lista:
        Atribui o próximo valor à cabeça da Lista
3 - Valor do elemento recebe o valor do próximo do próximo elemento ('pula' uma casa)
4 - Se o proximo elemento não for NULO:
        Atribui ao valor anterior o valor da cabeça da Lista
5 - Diminui o tamanho da Lista em 1 casa