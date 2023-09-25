# Golang - Go
Esse README irá servir como o histórico do meu aprendizado em Go. Tudo que eu for aprendendo,
na ordem em que eu fui aprendendo, estará escrito aqui (pelo menos até eu considerar necessário).

## Primeiros comandos

```bash
go mod init <nome-projeto>
go run <filename>
go build
GOOS=windows go build # GOOS pode linux, mac, etc.
```

## Conceitos

- variáveis
```go
// Versão verbosa
var a string
a = "Hello"

// Versão menos verbosa
// Nesse caso ele infere que a variável é uma string
// e como é fortemente tipada, não podemos mais mudar
// o tipo
a := "Hello"
```

- funções
Funções também é bem parecido em C, onde precisamos
especificar o tipo dos argumentos, e tipo do retorno da
função também. 

Mas o interessante é que podemos retornar mais de um tipo
(parece o retorno de tuplasem Python).

```go
func soma(x int, y int) (int, bool) {
    if x > 10 {
        return x + y, true
    }

    return x + y, false
}

res, status := soma(10, 2)
```

- imports
Os imports sempre são feitos assim:

```go
package main
import "net/http"

func main() {
    http.ListenAndServe(":8080", nil)
}
```

- structs
O Go não tem suporte a orientação a objetos. Mas temos structs, que
são amplamente utilizadas para implementar diversos
design patterns.

```go
type Course struct {
    Name string
    Description string
    Price int
}

// main
course := Course{
    Name: "Golang",
    Description: "Golang Course",
    Price: 100,
}

println(course.Name)
```

Nós temos "métodos" para essas structs, que nada mais são
que funções que recebem um argumento do tipo da struct. É uma
sintaxe levemente diferente de funções normais.

```go
func (c Course) GetFullInfo() string {
	return fmt.Sprintf(c.Name, c.Description, c.Price)
}

// main
course.GetFullInfo
```

> Repare que não retornei println, mas sim fmt.Sprintf

**CONVENÇÃO IMPORTANTE**
Todo nome em go, que começar em maiúsculo, significa que é público,
enquanto que todo nome que começa em minúsculo, significa que é privado.

Portanto `Course` e `GetFullInfo` são públicos, e qualquer pacote
consegue acessá-los.

---

- lidando com json
Em Go, frequentemente utilizamos structs para representar DTOs - 
Data Transfer Object - no backend, i.e., eventualmente vamos precisar mandar
essa struct para um frontend, ou para outro serviço. O formato que geralmente
é utilizado é o JSON, e portanto é bem fácil traduzir structs para json:

```go
type Course struct {
    //                 tag para modificarmos o campo do json
	Name        string `json:name`
	Description string `json:description`
	Price       int    `json:price`
}
```

- paralelismo com threads - go routines
```go
func counter() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
        time.Sleep(time.Second)
    }
}

// thread 0
func main() {
    go counter() // thread 1
    go counter() // thread 2
}
```

## Misc
- [Sobre timeouts](#https://go.dev/blog/concurrency-timeouts)