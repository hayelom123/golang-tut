package main

import "fmt"

type UserId int

func main() {
	var id UserId = 10
	fmt.Println(id)
	fmt.Printf("%T\n", id)

}
