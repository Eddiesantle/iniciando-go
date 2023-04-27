package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second)
	}
}

// LOADBALANCE -> recebe e manda par aworker
// 	WORKER -> execulta tarefa

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {

	// MULTITRADIN - rodando de forma assincrona
	// go task("Tarefa 1")
	// go task("Tarefa 2")
	// task("Tarefa 3")

	// trading conversar entre si
	/* canal := make(chan string)

	go func () {
		canal <- "HELLO WORD!!!"
	}()
	msg := <-canal
	fmt.Println(msg) */

	// loadbalance and workers
	data := make(chan int)

	// T2 - trading
	/* 
	go worker(1, data) // 2k memoria pra cada worker
	go worker(2, data) 
	*/
	qtdWorkers := 100
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data)
	}

	// T1 - trading
	for i := 0; i < 10000; i++ {
		data <- i
	}

}
