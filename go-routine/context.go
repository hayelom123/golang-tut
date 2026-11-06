package main

import (
	"fmt"
	// "time"
	"context"
	"sync"
)

func checkDrawing(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Drawing canceled")
	default:
		fmt.Println("Drawing is ready")
	}

}

func processDrawing(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Drawing processing started")

	<-ctx.Done()

	fmt.Println("Drawing processing canceled: ", ctx.Err())
}
func main() {

	//  ctx,cancel:=context.WithCancel(context.Background())

	//   fmt.Println("Before:", ctx.Err())//Err() is a method of the context. It tells us why the context ended.
	//    //Before cancellation, it returns nil, meaning there is no cancellation error.
	//    processDrawing(ctx)

	//   cancel()
	//    fmt.Println("After:", ctx.Err()) //After cancel() is called, it returns context.Canceled.

	// fmt.Println("Cancellation detected")
	// fmt.Println("Reason:", ctx.Err())

	// ctx, cancel := context.WithCancel(
	//     context.Background(),
	// )

	//  cancel()
	// checkDrawing(ctx)

	// <-ctx.Done()

	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	var wg sync.WaitGroup

	wg.Add(1)

	go processDrawing(ctx, &wg)

	cancel()

	wg.Wait()

	fmt.Println("Main finished")

}
