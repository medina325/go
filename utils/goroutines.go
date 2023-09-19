package utils

import (
	"fmt"
	"time"
)

func worker(workerId int, channel chan int) {
	for x := range channel {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func SimulateWorkers() {
	channel := make(chan int)

	go worker(1, channel)
	go worker(2, channel)

	for i := 0; i < 100; i++ {
		channel <- i
	}

}
