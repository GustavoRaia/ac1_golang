# Lista Circular Simplesmente Encadeada

Uma lista circular simplesmente encadeada é uma estrutura de dados em que cada elemento, chamado de nó, contém dois componentes principais: um valor de dados e um ponteiro que aponta para o próximo nó na lista. A principal diferença entre uma lista circular simplesmente encadeada e uma lista simplesmente encadeada tradicional é que, na lista circular, o último nó aponta de volta para o primeiro nó, formando um ciclo contínuo.

## Exibição dos Nós em uma Lista Circular

1 - Nó recebe o valor de inicio da Lista

2 - Para cada Nó diferente de NULO:

        Exiba na tela o valor do Nó atual
        Atribua ao nó atual o valor de próximo

## Inserção de um Nó no início da Lista

1 - Variavel nó recebe valor a ser adicionado
2 - Se primeiro item da lista for vazio:

        Atribui o início da Lista ao valor do ponteiro de Nó
3 - Se o final da lista não for vazio:

        Atribui ao valor de 'próximo' o valor do ponteiro de Nó
4 - Atribui ao fim da Lista o valor do ponteiro de Nó


## Exclusão do Primeiro Nó da Lista

1 - Se valor de início da lista for NULO:

        Retorna mensagem de erro
        Encerra a função
        
2 - Atribui o valor de 'próximo' ao valor do início da lista
