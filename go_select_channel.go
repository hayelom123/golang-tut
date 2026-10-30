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
	default:
        fmt.Println("No value available")
    
    }

	ch3:=make(chan int,1)
	ch4:=make(chan int,1)

	ch3<-100
	ch4<-200

	select {
	case message:= <-ch3:
		 fmt.Println("Received from channel 3:", message)
	case message:= <-ch4:
		 fmt.Println("Received from channel 4:", message)
	}

}