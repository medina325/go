package goroutines

import (
	"fmt"
	"math/rand"
	"time"
)

func Timeout(timeoutThres time.Duration) {
	// channel de bools, com buffer com 1 espaço
	timeout := make(chan bool, 1)

	ch := make(chan string)

	go func() {
		randomWait := randomNumber(10) * time.Second
		fmt.Println(randomWait)
		time.Sleep(randomWait)
		ch <- "Escrevi essa mensagem em " + randomWait.String() + " segundos, fez sentido? Espero que sim!"
	}()

	go func() {
		time.Sleep(timeoutThres)
		timeout <- true
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
		fmt.Println("Leitura de ch aconteceu antes de " + timeoutThres.String() + " segundos")
	case <-timeout:
		fmt.Printf("Timeout depois de %s segundos\n", timeoutThres)
	}

}

func randomNumber(_range int) time.Duration {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return time.Duration(rng.Intn(_range))
}
