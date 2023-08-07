# Prática de Linguagem em Go

Este repositório contém exemplos e orientações como prática da linguagem de programação Go (Golang). A motivação por trás deste projeto é proporcionar um espaço para aprender, praticar e aprofundar meu conhecimento em Go.

## Executando Aplicativos

Para executar um aplicativo Go, você pode usar o comando. Por exemplo:

```sh
go run main.go
```

Este comando executará o programa diretamente, sem gerar um arquivo binário.

### Compilando para Windows

Para compilar a aplicação e gerar um arquivo binário, você pode usar o seguinte comando:

```sh
go build main.go
```

Isso criará um executável chamado main (ou main.exe no Windows) que você pode executar diretamente.

Se você deseja compilar a aplicação para Windows em um ambiente diferente, como Linux, pode usar a variável de ambiente GOOS:

```sh
GOOS=windows go build main.go
```

Isso criará um executável main.exe compatível com Windows.

## Exemplos e Tópicos Abordados

### Goroutines e Channels

Goroutines são funções ou métodos executados em concorrência.
O arquivo main.go contém exemplos de como usar goroutines e channels para executar tarefas concorrentes e trocar informações entre elas.

### Funções e Métodos

O arquivo main.go\_ demonstra como definir funções e métodos em Go, incluindo exemplos de declaração de funções, retornos múltiplos e uso de ponteiros.

### Consumindo APIs

O exemplo main.go também inclui um exemplo de como fazer uma requisição HTTP para uma API externa (ViaCEP) e como processar a resposta JSON.

## Como Contribuir

Sinta-se à vontade para explorar os exemplos e fazer alterações para aprimorar ou expandir os conceitos apresentados. Caso queira contribuir, siga as diretrizes a seguir:

1. Faça um fork deste repositório.
2. Crie uma branch para sua contribuição: git checkout -b minha-contribuicao.
3. Faça suas alterações e adições.
4. Faça commit das suas alterações: git commit -m "Minha contribuição".
5. Faça push da sua branch: git push origin minha-contribuicao.
6. Crie um Pull Request neste repositório.

## Recursos Adicionais

- [Documentação Oficial do Go](https://golang.org/doc/)
- [Tour pelo Go](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example](https://gobyexample.com/)
- [Awesome Go](https://github.com/avelino/awesome-go)
