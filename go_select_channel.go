package main

import ("fmt")

func main(){
	ch1 := make(chan string, 1)
    ch2 := make(chan string, 1)

	ch1 <- "Hello from channel 1"
    ch2 <- "Hello from channel 2"

	select {
		
    case message := <-ch1:
        fmt.Println(message)

    case message := <-ch2:
        fmt.Println(message)
    }

}