package main

import "sync"

func fanIn(ch1, ch2 <-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	merge := func(ch <-chan string) {
		defer wg.Done()
		for msg := range ch {
			out <- msg
		}
	}

	wg.Add(2)
	go merge(ch1)
	go merge(ch2)

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}
