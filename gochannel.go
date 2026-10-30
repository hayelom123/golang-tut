package main

import ("fmt")

func sendMessage(ch chan string){
	ch <- "Hello from the channel"
}

func main(){

	ch:= make(chan string)

	go	sendMessage(ch)

	message := <- ch

	fmt.Println(message)

}