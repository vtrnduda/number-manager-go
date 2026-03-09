## Number Manager in Go 🧮
Um utilitário de linha de comando (CLI) simples e eficiente para gerenciar listas de números inteiros, focado em demonstrar os fundamentos da linguagem Go, como slices, funções modulares e tratamento de erros.

## Funcionalidades
O programa oferece um menu interativo com as seguintes opções:

- Add Number: Insere um novo inteiro na lista.
- List Numbers: Exibe todos os números armazenados.
- Remove by Index: Exclui um elemento validando os limites do slice.
- Statistics: Calcula Mínimo, Máximo e Média (com funções granulares).
- Safe Division: Realiza divisões tratando erros de divisor zero.
- Clear List: Esvazia o armazenamento atual.

## Tecnologias Utilizadas
Linguagem: Go (Golang)

Estruturas: Slices, Variadic Functions, Multiple Returns, Error Handling (if err != nil).

## Como Executar

### Pré-requisitos
Ter o Go instalado (versão 1.18 ou superior recomendada).

### Passo a Passo
Salve o código em um arquivo chamado main.go.

Abra o terminal na pasta do arquivo.

 Inicialize o módulo:

```Bash
go mod init number-manager
```

Execute o projeto:

```Bash
go run main.go
```