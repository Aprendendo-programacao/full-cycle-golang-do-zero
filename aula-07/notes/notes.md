# Structs, Composição e Json com Golang

### struct

* **Criar uma `struct`**

  ```go
  type Cliente struct {
    Nome  string
    Email string
    CPF   int
  }
  ```

* **Declarar uma `struct`**

  ```go
  cliente := Cliente{
		Nome:  "Wesley",
		Email: "wesley@gmail.com",
		CPF:   12345678900,
	}

	fmt.Println(cliente) // {Wesley wesley@gmail.com 12345678900}
  ```

  **OU**

  ```go
  cliete2 := Cliente{
		"Mari",
		"mari@gmail.com",
		98765432111,
	}

  fmt.Printf("Nome: %s, Email: %s, CPF: %d\n", cliete2.Nome, cliete2.Email, cliete2.CPF) // Nome: Mari, Email: mari@gmail.com, CPF: 98765432111
  ```

### Composição

```go
type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

type ClienteInternacional struct {
	Cliente
	Pais string
}
```

```go
cliete3 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Davi",
			Email: "davi@gmail.com",
			CPF:   12312312300,
		},
		Pais: "EUA",
	}
	fmt.Printf("Nome: %s, Email: %s, CPF: %d, Pais %s\n", cliete3.Nome, cliete3.Email, cliete3.CPF, cliete3.Pais) // Nome: {Davi davi@gmail.com %!s(int=12332112300)}, Email: davi@gmail.com, CPF: 12332112300, Pais> EUA
```

```go
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
	Pais string
}

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
	cliente1.Exibe() // Exibindo cliente pelo método:  Wesley
	cliente2.Exibe() // Exibindo cliente pelo método:  Mari
	cliente3.Exibe() // Exibindo cliente pelo método:  Davi
```

### JSON

* **Converter uma `struct` em JSON**

  ```go
  cliente := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Davi",
			Email: "davi@gmail.com",
			CPF:   12312312300,
		},
		Pais: "EUA",
	}

  clienteJson, _ := json.Marshal(cliente)

  fmt.Println(clienteJson) // retorna um slice de bytes ([]byte)
  fmt.Println(string(clienteJson)) // {"Nome":"Davi","Email":"davi@gmail.com","CPF":12312312300,"Pais":"EUA"}
  ```

* **TAGs**

  ```go
  type Cliente struct {
    Nome  string
    Email string
    CPF   int
  }

  type ClienteInternacional struct {
    Cliente
    Pais string `json:"pais"`
  }

  func main() {

    cliente3 := ClienteInternacional{
      Cliente: Cliente{
        Nome:  "Davi",
        Email: "davi@gmail.com",
        CPF:   12312312300,
      },
      Pais: "EUA",
    }

    clienteJson, _ := json.Marshal(cliente3)

    fmt.Println(string(clienteJson)) // {"Nome":"Davi","Email":"davi@gmail.com","CPF":12312312300,"pais":"EUA"}

  }
  ```

* **Converter JSON para `struct`**

  ```go
  jsonCliente := `{"Nome":"Davi","Email":"davi@gmail.com","CPF":12312312300,"pais":"EUA"}`
  cliente4 := ClienteInternacional{}

  err = json.Unmarshal([]byte(jsonCliente), &cliente4)

  if err != nil {
    log.Fatal(err.Error())
  }

  fmt.Println(cliente4) // {{Davi davi@gmail.com 12312312300} EUA}
  ```