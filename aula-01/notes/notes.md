# Golang

### Características

* Linguagem de programação _Open Source_ que tem o objetivo de tornar os programadores mais produtivos

* Expressiva, concisa, limpa e eficiente

* Foi criada para aproveitar ao máximo dos recursos multicore e de rede

* Linguagem compilada --> compilação rápida e ao mesmo tempo trabalho com garbage collection

* Estaticamente tipada, compilada mas que ao mesmo tempo parece até uma linguagem dinamicamente tipada e interpretada

* Compilada em apenas um arquivo binário

### Curiosidades

* Criada pelo Google

* Projeto inicial em 2007

* A versão 1.0 foi lançada em 2012

* Na versão 1.5, a linguagem reescreveu seu compilador em Go

* Criadores: Robert Griesemar, Rob Pike, Ken Thompson

* **Empresas que utilizam Golang**

  * Google
  * Docker
  * Kubernetes
  * Dropbox
  * Digital Ocean
  * Prometheus
  * Uber
  * SoundCloud
  * Medium
  * Twitch
  * BBC
  * SendGrip

### Motivos para a criação do Golang

* **Limitações das principais linguagens utilizadas na Google** (Python, Java, C++)

  * Python: problemas com **lentidão**

  * C/C++: muita **complexidade** e **lentidão na compilação**

  * Java: complexidade gerada ao longo do tempo, principalmente na **manutenção** --> **verbosidade da linguagem**

* **Necessidades**

  * Multithreading e concorrência: não foram criadas, nativamente, com esse propósito

### Vantagens na utilização do Golang

* Simplicidade

* Framework de testes e profiling nativos
 
* Detecção de _Race conditions_

* Deploy simples

* Baixa curva de aprendizado

### Comandos 

* **Visualizar as variáveis de ambiente do Go**: `$ go env`

* **Executar um arquivo `.go`**: `$ go run <arquivo>`

* **Compilar um arquivo `.go`**: `$ go build <arquivo>`

* **Gerar o binário**: `$ go install`

* **Instalar um dependência externa**

  * **Sintaxe**: `$ go get <repositório>`

  * **Exemplo**: `$ go get github.com/google/uuid`