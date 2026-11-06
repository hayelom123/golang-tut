package main

import (
	"fmt"
	// "time"
	"context"
)

func processDrawing(ctx context.Context){
	select{
	case <-ctx.Done():
		 fmt.Println("Drawing canceled")
	default:
        fmt.Println("Drawing processing started")
	}
	    
}

func main(){

	//  ctx,cancel:=context.WithCancel(context.Background())
	 
	//   fmt.Println("Before:", ctx.Err())//Err() is a method of the context. It tells us why the context ended.
	//    //Before cancellation, it returns nil, meaning there is no cancellation error.
	//    processDrawing(ctx)

	//   cancel()
    //    fmt.Println("After:", ctx.Err()) //After cancel() is called, it returns context.Canceled.

 
    // fmt.Println("Cancellation detected")
    // fmt.Println("Reason:", ctx.Err())

	ctx, cancel := context.WithCancel(
        context.Background(),
    )
    cancel()
    processDrawing(ctx)
    // <-ctx.Done()

    fmt.Println("Finished")
}