package main

func main() {
	kanal := make(chan int)
	<-kanal
}
