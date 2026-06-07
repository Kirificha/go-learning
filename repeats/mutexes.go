package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	thousand()
	fmt.Println(counter)
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

var counter int

var sc sync.Mutex

func thousand() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			sc.Lock()
			counter++
			defer sc.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
