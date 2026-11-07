package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func processDrawing(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("Drawing processing started")
		select {
		case <-ctx.Done():
			fmt.Println("Drawing canceled: ", ctx.Err())
			return
		case <-time.After(6 * time.Second):
			fmt.Println("Drawing completed: ", i)

		}
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go processDrawing(ctx, &wg)

	wg.Wait()
	fmt.Println("Main finished")

}
