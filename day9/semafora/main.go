package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sem := make(chan struct{}, 3)
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{}
			fmt.Printf("горутина %d работает\n", id)
			time.Sleep(time.Second)
			<-sem
		}(i)
	}

	wg.Wait()
}
