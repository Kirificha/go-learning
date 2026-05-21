package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	go worker(1, jobs, &wg)
	go worker(2, jobs, &wg)
	go worker(3, jobs, &wg)

	wg.Wait()
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Воркер %d обработал задачу %d\n", id, job)
	}
}
