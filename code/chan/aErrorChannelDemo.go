package main

import (
	"fmt"
	"time"
)

//
//chan is closed , trigger panic
//往已经关闭的channel写入数据会panic的。因为main在开辟完两个goroutine之后，立刻关闭了ch
func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
