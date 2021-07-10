# Mão na massa com variáveis, tipos e pacotes

### Variáveis

* **Declaração de uma variável**

  ```go
  var lang string
  ```

* **Atribuição de uma valor a uma variável**

  ```go
  lang = "Go"
  ```

* **Declaração e atribuição de uma variável**

  ```go
  lang := "Go"
  lang = "JavaScript"
  ```

  > OBS: `:=` só é necessário apenas na declaração e na inferência do tipo, nas próximas atribuições basta utilizar `=`

  ```go
  var lang string = "Go"
  ```

### Tipos de dado

Tipo | Exemplo
:-----: | :-----:
string | `"Golang"`
int | `10`
float64 | `3.14`
bool | `false`

* **Exemplo**

  ```go
  a := 10
  b := "World"
  c := 3.14
  d := false
  e := `Hello

    World
    `

  fmt.Printf("%T \n", a)
  fmt.Printf("%T \n", b)
  fmt.Printf("%T \n", c)
  fmt.Printf("%T \n", d)
  fmt.Printf("%T \n", e)
  ```

### Funções

* **Função com retorno**

  ```go
  func Soma(a int, b int) int {
    return a + b
  }
  ```

* **Chamar uma função**

  ```go
  resultado := Soma(1, 1)
	fmt.Printf("Tipo: %T | Valor: %v", resultado, resultado) // Tipo: int | Valor: 2%
  ```

* **Importar funções**

  ```go
  package math

  func Soma(a int, b int) int {
	  return a + b
  }
  ```

  ```go
  package main

  import (
    "fmt"
    "variaveis-tipos-e-pacotes/math" // nome-da-aplicacao/pacotes/...
  )

  func main() {
    resultado := math.Soma(1, 1)
    fmt.Printf("Tipo: %T | Valor: %v", resultado, resultado) // 2
  }

  ```

### Visibilidade (variável/função/`struct`/atributo)

* **IMPORTANTE**: Go é **case-sensitive**

* **Função exportada**

  ```go
  func Soma(a int, b int) int {
    return a + b
  }
  ```

  > 1º letra do nome maiúscula = "público"

* **Função não exportada**

  ```go
  func soma(a int, b int) int {
    return a + b
  }
  ```

  > 1º letra do nome minúscula = "privado"

* **Documentação de funções exportadas**

  > É uma boa prática sempre documentar as funções exportadas

  ```go
  // <Nome da função> <descrição sobre o que ela faz>
  func Soma(a int, b int) int {
    return a + b
  }
  ```