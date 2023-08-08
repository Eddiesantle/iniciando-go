package utils

import (
	"errors"
	"fmt"
	"time"
)

// Definição da estrutura ViaCEP
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

// Função para somar dois números e retornar o resultado ou um erro
func Soma(a, b int) (int, error) {
	if a > b {
		return a + b, nil
	}
	return 0, errors.New("b não pode ser maior que a")
}

// Método que define o valor do CEP na estrutura ViaCEP
// Ponteiro( *STRUCT ): adicionando ele substitui, se deixar sem ele apenas copia
func (v *ViaCEP) SetCep(cep string) {
	v.Cep = cep
	fmt.Println(v.Cep)
}

// Método que retorna o endereço completo baseado nos campos da estrutura ViaCEP
func (v ViaCEP) EnderecoCompleto() string {
	// METODO porque tem uma struct anexado
	return fmt.Sprintf("%s, %s, %s, %s, %s", v.Logradouro, v.Complemento, v.Bairro, v.Localidade, v.Uf)
}

// Função que simula uma tarefa
func Task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second)
	}
}

// Função que representa um worker que recebe e executa tarefas
// LOADBALANCE -> Recebe e manda para worker
// WORKER -> Executa tarefa
func Worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
