package main

import (
	"fmt"
	"sync"
	"context"
	"time"
)

func processDrawing(ctx context.Context,wg *sync.WaitGroup){
	defer wg.Done()
	for i:=1;i<=5;i++{
		select{
		case <-ctx.Done(): 
				fmt.Println("Failed to draw: ",ctx.Err())
				return
		case <-time.After(2*time.Second):
		      fmt.Println("Drawing: ",i)
		}
	}
}

func main(){

	currentTime:=time.Now()
	fmt.Println("Current time: ", currentTime)

	deadline:=time.Now().Add(5*time.Second)
	fmt.Println("Deadline: ", deadline)

	ctx,cancel:=context.WithDeadline(context.Background(),deadline)

	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go processDrawing(ctx,&wg)

	wg.Wait()

}