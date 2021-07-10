# Iniciando com Golang na prática - Aula 2

### Comandos

* **Gerar um executável para um OS específico**

  * **Sintaxe**: `$ GOOS=<OS> go build <arquivo>`

  * **Exemplo**: `$ GOOS=windows go build hello.go`

    > **OBS**: para usuários Windows: `$env:GOOS="linux"; go build hello.go`

### Sintaxe

* **Imports**

  > OBS: _imports_ desnecessário = erro

  * **1 _import_**

    ```go
    import "fmt"
    ```

  * **2 ou + _import_**

    ```go
    import (
      "fmt"
      "..."
    )
    ```