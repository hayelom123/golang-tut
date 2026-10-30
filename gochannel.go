package main

import ("fmt")

func sendMessage(ch chan string){
	ch <- "Hello from the channel"
}
func sendNumbers(ch chan int){
    for i:=1; i<=3; i++ {
		ch <- i
	}
	close(ch)
}

func main(){

	ch:= make(chan string)

	go	sendMessage(ch)

	message := <- ch

	fmt.Println(message)

	ch1 :=make(chan int)

	go sendNumbers(ch1)

	for number := range ch1 {
		fmt.Println(number)
	}
  fmt.Println("Finished!")
}