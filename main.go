// Similar ao C, nós precisamos definir um arquivo "main" que
// servirá como entrypoint do programa. No caso, o nome do arquivo
// pode ser qualquer coisa, mas deve importar o pacote main,
// e definir uma função main
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func soma(x int, y int) int {
	return x + y
}

func soma_com_status(x int, y int) (int, bool) {
	if x > 10 {
		return x + y, true
	}

	return x + y, false
}

type Course struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Desc: %s, Price: %d", c.Name, c.Description, c.Price)
}

func home(w http.ResponseWriter, r *http.Request) {
	course := Course{
		Name:        "Golang",
		Description: "Oin",
		Price:       100,
	}
	println(course.Name)
	course.GetFullInfo()

	json.NewEncoder(w).Encode(course)
}

func hello(w http.ResponseWriter, r *http.Request) {
	// println(r)
	w.Write([]byte("Hello, World!!!"))
}

func server() {
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func worker(workerId int, channel chan int) {
	for x := range channel {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	// a := "Hello, World!"
	// println(a)
	// println(soma(1, 2))

	res, status := soma_com_status(10, 2)
	println(res, status)

	// go counter() // thread 1
	// go counter() // thread 2

	channel := make(chan int)

	go worker(1, channel)
	go worker(2, channel)

	for i := 0; i < 10; i++ {
		channel <- i
	}

	server()

	horas := 1
	minutos := 1
	segundos := 1
	fmt.Printf("%d, %d, %d", horas, minutos, segundos)

}
