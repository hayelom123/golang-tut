package main

import (
	"fmt"
	// "time"
	"context"
)

func processDrawing(ctx context.Context){
	    fmt.Println("Drawing processing started")
}

func main(){

	 ctx:=context.Background()
	   processDrawing(ctx)

}