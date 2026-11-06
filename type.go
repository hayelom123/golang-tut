package main

import "fmt"

type UserId int

type Email string
type Username string

// alias
type ID=int

func main() {
	var id UserId = 10
	var number int = 101

	id = UserId(number)

	fmt.Println(id)
	fmt.Printf("%T\n", id)

	var uid ID = 201
	uid = number

	fmt.Println(uid)
	fmt.Printf("%T\n", uid)

	var email Email = "hayelomkiros20@gmail.com"
	sendEmail(email)


}

// they guard safty
func sendEmail(to Email) {
	fmt.Printf("Email sent to %v\n", to)
}
