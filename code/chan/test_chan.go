package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START_OMIT

var demoChan = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("seed: %v\n", value)
	demoChan <- value

}

func send2(v chan int) {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("seed2: %v\n", value)
	v <- value

}

func main() {

	defer close(demoChan)
	go send()
	go send2(demoChan)
	fmt.Printf("wait\n")
	value := <-demoChan
	fmt.Printf("received: %v\n", value)

	fmt.Printf("end\n")

}

// END_OMIT
