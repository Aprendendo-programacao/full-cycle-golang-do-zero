# Essa é a técnica para tratar erros em Golang

### Erro

* No Go, o erro é tratado logo após a chamada de uma função que possui como retorno, ou um dos retorno, `error`

* Por padrão, o nome da variável que recebe o erro é `err`

* **Exemplo**

  ```go
  response, err := http.Get("https://www.google.com.br/")

	if err != nil {
		// Tratamento do erro
    log.Fatal(err.Error())
	}

	// Código 
  ```

  ```go
  func main() {

    result, err := soma(7, 2)

    if err != nil {
      log.Fatal(err.Error())
    }

    fmt.Println(result)

  }

  func soma(x int, y int) (int, error) {
    result := x + y

    if result > 10 {
      return 0, errors.New("total maior que 10")
    }

    return result, nil
  }
  ```

* **Blank Identifier**

  > **Blank Identifier** (`_`) tem como função **descartar a variável**. No caso do uso desse recurso em um erro, o tratamento dela não é mais obrigatório

  ```go
  func main() {

    result, _ := soma(7, 2)

    fmt.Println(result)

  }

  func soma(x int, y int) (int, error) {
    result := x + y

    if result > 10 {
      return 0, errors.New("total maior que 10")
    }

    return result, nil
  }
  ```