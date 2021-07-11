# Aprenda a trabalhar com ponteiros usando Golang

### Ponteiros

![](https://media.geeksforgeeks.org/wp-content/uploads/20190705160332/Pointers-in-Golang.jpg)

* **Variável = armazena um valor**

* **Ponteiro = armazena o endereço na memória** de um valor

* **Obter o endereço na memória de um valor (`&`)**

  ```go
  a := 10
	fmt.Println(&a) // 0xc00018c000
  ```

  > `0xc00018c000` é o endereço na memória onde o valor 10 está alocado (armazenado)

* **Ponteiro (`*`)**

  ```go
  a := 10
	fmt.Println(&a) // 0xc00018c000

	var b *int = &a
	fmt.Println(b) // 0xc0000b6010
  ```

  > A variável b armazena um **endereço na memória** de um **valor do tipo `int`**

  <br>

  ```go
	a := 10
	var b *int = &a
	fmt.Println(*b) // 10
  ```

  > Ao colocar o **`*` na frente de um ponteiro**, é **recuperado o valor do apontamento** da variável b  (armazena o endereço na memória do valor 10)

  <br>

  ```go
  a := 10
  var b *int = &a
  fmt.Println(*b) // 10

  *b = 50
  fmt.Println(*b) // 50
  ```

  > `*b = 50` = atribuir 50 como novo valor do endereço na memória (`0xc0000b6010`) que estava armazenado no ponteiro "b"

  <br>

  ```go
  a := 10

  var b *int = &a
  *b = 50

  c := &a
  *c = 60

  fmt.Println(a)
  fmt.Println(*b)
  fmt.Println(*c)
  ```

  > Como `a`, `b` e `c` estão apontando para o mesmo endereço na memória (`0xc0000b6010`) logo, se o valor armazenado no endereço `0xc0000b6010` for alterado todas as variáveis que apontam para ela também serão alteradas

  <br>

  * **Exemplo**

      ```go
      func main() {
        a := 10

        fmt.Println(a) // 10
        alterarValor(&a)
        fmt.Println(a) // 15
      }

      func alterarValor(a *int) {
        *a = 15
      }
      ``` 

  * **Ponteiro e `struct`**

    * **Problema:** alteração de um atributo apenas no escopo local

      ```go
      type Carro struct {
        Name string
      }

      func (c Carro) andar() {
        c.Name = "BMW" // O atributo "Name" só foi alterado no escopo local (função "andar()")
        fmt.Println(c.Name, "andou!")
      }

      func main() {

        carro := Carro{
          Name: "Ford Ka",
        }

        carro.andar() // BMW andou!
        fmt.Println(carro.Name) // BMW andou!
      }
      ```

    * **Solução:** alteração do valor na memória com ponteiro

      ```go
      type Carro struct {
        Name string
      }

      func (c *Carro) andar() {
        c.Name = "BMW" // a alteração do atributo se torna global ao utilizar o ponteiro, pois foi alterado o valor na memória
        fmt.Println(c.Name, "andou!")
      }

      func main() {

        carro := Carro{
          Name: "Ford Ka",
        }

        carro.andar()
        fmt.Println(carro.Name)
      }
      ```