package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	numChan := make(chan int)
	strChan := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			numChan <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			num := <-numChan
			strChan <- strconv.Itoa(num)
		}
	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 10; i++ {
		str := <-strChan
		fmt.Println(str)
	}
}
