package main

import (
	"fmt"
	"time"
)

// START_OMIT

func Producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i // write data to this channel
		fmt.Println("create :", i)
	}
}

func Consumer(queue <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-queue // same, read
		fmt.Println("receive:", v)
	}
}

func main() {
	queue := make(chan int, 88)
	go Producer(queue)
	go Consumer(queue)
	time.Sleep(1 * time.Second)
}

// END_OMIT
