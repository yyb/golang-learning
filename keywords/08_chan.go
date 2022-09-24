package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type Message struct {
	body string
}

func main() {

	// 管道传递结构
	var pipline chan Message
	pipline = make(chan Message)
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		message := <-pipline

		fmt.Println(message.body)
	}()

	pipline <- Message{"Hello"}

	wg.Wait()

	// 只读和只写管道
	var rw chan string
	rw = make(chan string)
	go read(rw)
	go write(rw)
	fmt.Println(math.Pow(10, 6))
	time.Sleep(time.Second)

}

func write(w chan<- string) {
	w <- "hello world"
}

func read(r <-chan string) {
	fmt.Println(<-r)
}
