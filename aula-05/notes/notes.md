# Golang: Você precisa saber isso sobre funções

### Tipos de funções

* **Função sem retorno**

  ```go
  func soma(a int, b int) {
    fmt.Printf("Resultado: %v", a + b)
  }
  ```

* **Função com retorno normal**

  ```go
  func soma(a int, b int) int {
    return a + b, "somou!"
  }
  ```

* **Função com retorno múltiplo**

  ```go
  func soma(a int, b int) (int, string) {
    return a + b, "somou!"
  }
  ```

* **Função com retorno nomeado**

  ```go
  func soma(a int, b int) (result int) {
    result = a + b
    return
  }
  ```

* **Função variadic**

  ```go
  func main() {
    result := somaTudo(1,2, 3, 4, 5, 6, 7, 8, 9, 10)

    fmt.Println(result)
  }

  func somaTudo(x ...int) int {
    resultado := 0

    for _, value := range x {
      resultado += value
    }

    return resultado
  }
  ```

* **Funções anônimas**

  ```go
  resultado := func(x ...int) func() int {
		result := 0
		
		for _, value := range x {
			result += value
		}
		
		return func() int {
			return result * result
		}
	}
  ```

  > OBS: a função anônima retorna um função (`func() int`)

  * **Imprimir a referência na memória da função anônima**

    ```go
    fmt.Println(resultado(10, 10, 10)) // 0x497860
    ```

  * **Imprimir a o resultado da função anônima**

    ```go
    fmt.Println(resultado(10, 10, 10)()) // 900
    ```

    > **Sintaxe**: `variavel(parametros, ...)()`

### Métodos

> Métodos são funções que estão relacionadas a uma `struct` (**binding**)

* **Exemplo**

  ```go
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
  ```