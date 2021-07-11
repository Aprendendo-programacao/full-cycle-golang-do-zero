package main

import "fmt"

type Carro struct {
	Name string
}

func (c *Carro) andar() {
	c.Name = "BMW"
	fmt.Println(c.Name, "andou!")
}

func main() {

	carro := Carro{
		Name: "Ford Ka",
	}

	carro.andar()
	fmt.Println(carro.Name)
}