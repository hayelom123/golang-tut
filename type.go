package main

import "fmt"

type UserId int

func main() {
	var id UserId = 10
	var number int =101

	id=UserId(number)

	fmt.Println(id)
	fmt.Printf("%T\n", id)
}
