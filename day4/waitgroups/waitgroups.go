package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	wg.Wait()
}

func worker(id int) {
	fmt.Printf("Работник %d стартовал\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Работник %d закончил\n", id)
}
