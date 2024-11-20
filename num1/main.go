package main

import (
	"fmt"
	"sync"
	"time"
)

var shetchik int
var mu sync.Mutex

func increment() {
	mu.Lock()
	shetchik++
	mu.Unlock()
}

func main() {
	for i := 0; i < 5; i++ {
		go increment()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Конечное значение счетчика:", shetchik)
}
