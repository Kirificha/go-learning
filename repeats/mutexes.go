package main

import (
	"fmt"
	"sync"
)

func main() {
	gor()
}

func worker(i int) {
	fmt.Printf("\nЯ рабочий %d", i)
}

func gor() {

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i <= 5; i++ {
		go func(n int) {
			worker(n)
			wg.Done()
		}(i)
	}

	wg.Wait()

}
