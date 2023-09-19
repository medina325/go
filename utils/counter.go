package utils

import (
	"fmt"
	"time"
)

func Counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
