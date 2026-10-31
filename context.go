
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

    fmt.Println("Worker stopped:", ctx.Err())
}

func main() {

	 ctx1, cancel1 := context.WithCancel(context.Background())

    finished := make(chan struct{})

    go worker(ctx1, finished)

    cancel1()

    <-finished

    fmt.Println("Main finished")

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
