package main

import (
	"fmt"
	"time"
)

// Função que simula uma tarefa
func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second)
	}
}

// Função que representa um worker que recebe e executa tarefas
// LOADBALANCE -> Recebe e manda par aworker
// WORKER -> Execulta tarefa
func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {

	// Exemplo de concorrência usando goroutines -> MULTITRADIN - rodando de forma assincrona
	// go task("Tarefa 1")
	// go task("Tarefa 2")
	// task("Tarefa 3")

	// Exemplo de comunicação entre goroutines usando canais - trading conversar entre si
	/* canal := make(chan string)

	go func () {
		canal <- "HELLO WORD!!!"
	}()
	msg := <-canal
	fmt.Println(msg) */

	// Exemplo de load balancing e workers usando goroutines e canais
	data := make(chan int)

	// T2 - trading
	/*
		go worker(1, data) // 2k memoria pra cada worker
		go worker(2, data)
	*/

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
