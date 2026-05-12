package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина", id, "остановлена")
			return
		default:
			fmt.Println("Горутина", id, "работает")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, 1)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(time.Second)
}
