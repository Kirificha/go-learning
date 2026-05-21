package main

import "fmt"

func main() {
	msg1 := make(chan string)
	msg2 := make(chan string)

	go func() {
		msg1 <- "Hello"
		msg1 <- "from ch1"
		close(msg1)
	}()

	go func() {
		msg2 <- "Hello"
		msg2 <- "from ch2"
		close(msg2)
	}()

	merged := fanIn(msg1, msg2)
	for msg := range merged {
		fmt.Println(msg)
	}

}
