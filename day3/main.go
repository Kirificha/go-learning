package main

import "fmt"

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() { a <- 1; b <- 2 }()

	y := merge(a, b)

	fmt.Println(<-y)
	fmt.Println(<-y)
}
