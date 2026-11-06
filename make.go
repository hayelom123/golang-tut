package main

import "fmt"

func main(){
	numbers:=make([]int,3,5)
	fmt.Println(numbers)
	fmt.Println("length: ",len(numbers))
	fmt.Println("capacity: ",cap(numbers))
}