package main

import (
	"fmt"
	"time"
)

func goroutines() {

	go func(num int) {
		fmt.Println(num)
	}(1)
	go func(num int) {
		fmt.Println(num)
	}(2)
	go func(num int) {
		fmt.Println(num)
	}(3)
	time.Sleep(time.Second)
}

func goroutines1() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

func merge(a, b <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				out <- v
			case v := <-b:
				out <- v
			}
		}
	}()
	return out
}
