// Similar ao C, nós precisamos definir um arquivo "main" que
// servirá como entrypoint do programa. No caso, o nome do arquivo
// pode ser qualquer coisa, mas deve importar o pacote main,
// e definir uma função main
package main

import (
	"fmt"

	"github.com/medina325/go/utils"
)

func somaComStatus(x int, y int) (int, bool) {
	if x > 10 {
		return x + y, true
	}

	return x + y, false
}

func helloWorld() {
	a := "Hello, World!"
	println(a)
}

func main() {
	var option int
	fmt.Print(`Escolha uma das opções abaixo:
1 - run hello world
2 - run sum with status
3 - run server
4 - run counter threads
5 - run load balancing worker threads
	`)
	_, err := fmt.Scan(&option)

	if err != nil {
		fmt.Println("Fez caquinha")
		return
	}

	switch option {
	case 1:
		helloWorld()
	case 2:
		res, status := somaComStatus(10, 25)
		println(res, status)
	case 3:
		utils.Server()
	case 4:
		go utils.Counter()
		go utils.Counter()
		utils.Counter()
	case 5:
		utils.SimulateWorkers()
	}
}
