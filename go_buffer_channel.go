package main

import ("fmt")

func sendNumbers(ch chan int){
	for i:=1; i <=5;i++ {
		ch<-i
	}
	close(ch)
}

func main(){

	ch:=make(chan int, 2)

	go sendNumbers(ch)

	for number:= range ch{
		fmt.Println("Received",number)
	}
     fmt.Println("Finished!")

// 	Even though the channel capacity is only 2, we can send five numbers.
// Why? The receiver removes values as the sender produces them, freeing buffer space. 
// The sender may block temporarily when the buffer is full.

//  exercise
ch1:=make(chan int,3)
   ch1<-100
   ch1<-200
   ch1<-300

   fmt.Println(<-ch1)
   fmt.Println(<-ch1)
   fmt.Println(<-ch1)

//    for num:= range ch1 {
// 	fmt.Println(num)
//    }
}