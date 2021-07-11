package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo método: ", c.Nome)
}

type ClienteInternacional struct {
	Cliente
	Pais string `json:"pais"`
}

func main() {

	cliente1 := Cliente{
		Nome:  "Wesley",
		Email: "wesley@gmail.com",
		CPF:   12345678900,
	}
	fmt.Println(cliente1)

	cliente2 := Cliente{
		"Mari",
		"mari@gmail.com",
		98765432111,
	}
	fmt.Printf("Nome: %s, Email: %s, CPF: %d\n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Davi",
			Email: "davi@gmail.com",
			CPF:   12312312300,
		},
		Pais: "EUA",
	}
	cliente1.Exibe()
	cliente2.Exibe()
	cliente3.Exibe()

	clienteJson, err := json.Marshal(cliente3)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(clienteJson))

	jsonCliente := `{"Nome":"Davi","Email":"davi@gmail.com","CPF":12312312300,"pais":"EUA"}`
	cliente4 := ClienteInternacional{}

	err = json.Unmarshal([]byte(jsonCliente), &cliente4)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(cliente4)
}
