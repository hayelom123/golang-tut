package main

import (
	"context"
	"fmt"
)

func worker(ctx context.Context, progress chan<- int, finished chan<- struct{}) {

	defer close(finished)

	for i := 1; i <= 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", ctx.Err())
			return

		case progress <- i:
			// Report completed task
		}
	}

}

func main() {

	ctx, cancel := context.WithCancel(
		context.Background(),
	)
	defer cancel()

	progress := make(chan int)
	finished := make(chan struct{})

	go worker(ctx, progress, finished)

	// Receive two completed tasks
	for i := 0; i < 2; i++ {
		task := <-progress
		fmt.Println("Processed task:", task)
	}

	// Cancel remaining tasks
	cancel()

	// Wait for worker to stop
	<-finished

	fmt.Println("Main finished")

}

/*

Concept			What it does
go worker(...)	Starts a concurrent worker
progress <- i	Sends a task number
<-progress		Receives a task number
cancel()		Requests cancellation
<-ctx.Done()	Detects cancellation
<-finished		Waits for worker completion
*/
