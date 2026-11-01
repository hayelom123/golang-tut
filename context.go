
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, finished chan<- struct{}) {
    defer close(finished)

    fmt.Println("Worker started")

    <-ctx.Done()
    // ctx.Done() provides a channel that's closed when the context is canceled.
    // Because this worker waits on that channel, it won't continue until cancellation happens.

    fmt.Println("Worker stopped:", ctx.Err())
}

func main() {

	 ctx1, cancel1 := context.WithCancel(context.Background())

    finished := make(chan struct{})

    go worker(ctx1, finished)

    cancel1() //This signals that the worker should stop.

    <-finished //This ensures the main function doesn't exit before the worker finishes. The worker closes finished when it returns.

    fmt.Println("Main finished")

    /*
    cancel() tells the worker to stop. Waiting on finished tells the main function when the worker has actually stopped.
    */



    ctx, cancel := context.WithTimeout(
        context.Background(),
       2 * time.Second,
    )
    defer cancel()

    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Task completed")

    case <-ctx.Done():
        fmt.Println("Error:", ctx.Err())
    }
}
