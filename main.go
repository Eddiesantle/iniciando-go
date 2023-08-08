package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Eddiesantle/iniciando-go.git/utils"
)

func main() {

	// Exemplo de utilização de funções e métodos
	// fmt.Println(utils.Soma(1, 2))
	// resultado, err := utils.Soma(1, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(resultado)

	webViaCep()
	//loadbalanceWorker()

}

func webViaCep() {
	// (:=) quando é declarado a primeira vez usar ( : )
	// x := 1
	//  cep := utils.ViaCEP{
	// 	Cep: "01001-000",
	// 	Logradouro: "Praça da sé",
	// 	Complemento: "Lado impar",
	// 	Bairro: "sé",
	// 	Localidade: "São Paulo",
	// }
	// fmt.Println(cep)

	// Exemplo de requisição HTTP para obter informações do CEP usando a API ViaCEP
	cep := "01001000"
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		//parar programa
		panic(err)
	}
	defer req.Body.Close()

	// Lê o corpo da requisição
	res, err := io.ReadAll(req.Body)
	if err != nil {
		//parar programa
		panic(err)
	}

	// Converte o JSON para a estrutura ViaCEP
	var data utils.ViaCEP
	err = json.Unmarshal(res, &data)
	if err != nil {
		//parar programa
		panic(err)
	}

	// Se o retorno da tag for com letra maiuscula, diferente do struct ViaCEP que começa com letra maiuscula
	// {"cep:", "06011010"}
	// Imprime a localidade do CEP e o endereço completo
	fmt.Println(data.Localidade)
	fmt.Println(data.EnderecoCompleto())

	println("Hello, word")
	// Inicia um servidor HTTP na porta 8080
	http.ListenAndServe(":8080", nil)
}

func loadbalanceWorker() {
	// Exemplo de concorrência usando goroutines -> MULTITRADIN - rodando de forma assincrona
	// go utils.Task("Tarefa 1")
	// go utils.Task("Tarefa 2")
	// utils.Task("Tarefa 3")

	// Exemplo de comunicação entre goroutines usando canais - trading conversar entre si
	// canal := make(chan string)

	// go func () {
	// 	canal <- "HELLO WORD!!!"
	// }()
	// msg := <-canal
	// fmt.Println(msg)

	// Exemplo de load balancing e workers usando goroutines e canais
	data := make(chan int)

	// T2 - trading
	// go worker(1, data) // 2k memoria pra cada worker
	// go worker(2, data)

	// Iniciar múltiplos workers
	qtdWorkers := 100
	for i := 0; i < qtdWorkers; i++ {
		go utils.Worker(i, data)
	}

	// Enviar tarefas para os workers // T1 - trading
	for i := 0; i < 10000; i++ {
		data <- i
	}

	// Fechar o canal para encerrar os workers
	close(data)
}
