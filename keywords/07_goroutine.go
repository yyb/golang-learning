package main

import (
	"fmt"
	"sync"
)

type IHuman interface {
	say() string
}

type Human struct {
	name string
}

func (h *Human) say() string {
	return "My name is " + h.name
}

func output(human IHuman) {
	fmt.Println(human.say())
}

var wg sync.WaitGroup

func main() {

	John := &Human{"John"}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			output(John)
		}
	}()
	James := &Human{"James"}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			output(James)
		}
	}()

	wg.Wait()
}
