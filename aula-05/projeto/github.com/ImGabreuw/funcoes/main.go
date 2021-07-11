package main

import "fmt"

type Carro struct {
	Nome string
}

func (c Carro) andar() {
	fmt.Printf("O carro %v andou!", c.Nome)
}

func main() {

	carro := Carro{Nome: "BMW"}

	carro.andar()
}

func soma(a int, b int) (result int) {
	result = a + b
	return
}

func somaTudo(x ...int) int {
	resultado := 0

	for _, value := range x {
		resultado += value
	}

	return resultado
}