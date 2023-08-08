package main

// Definindo a estrutura ViaCEP para armazenar os dados do CEP
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

func main() {

	// Exemplo de concorrência usando goroutines -> MULTITRADIN - rodando de forma assincrona
	// go task("Tarefa 1")
	// go task("Tarefa 2")
	// task("Tarefa 3")

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
		go worker(i, data)
	}

	// Enviar tarefas para os workers // T1 - trading
	for i := 0; i < 10000; i++ {
		data <- i
	}

	// Fechar o canal para encerrar os workers
	close(data)
}

/*
func main() {

	// (:=) quando é declarado a primeira vez usar ( : )
	// x := 1

	//  cep := ViaCEP{
	// 	Cep: "01001-000",
	// 	Logradouro: "Praça da sé",
	// 	Complemento: "Lado impar",
	// 	Bairro: "sé",
	// 	Localidade: "São Paulo",
	// }
	// fmt.Println(cep)

	// Exemplo de utilização de funções e métodos
	// fmt.Println(soma(1,2))
	// resultado, err := soma(1, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(resultado)

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
	var data ViaCEP
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
*/
