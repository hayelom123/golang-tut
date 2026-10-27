package main

import "fmt"

type User struct{
	name string
	age  int
}



func main(){
   users := []User{}
	user1 := User{
		name: "Hayelom Kiros",
		age: 29,
	}


	users = append(users,user1)
     users = append(users,User{
		name: "Kassa Kiros",
		age:32,
	 })
	fmt.Printf("%s age %d \n",user1.name,user1.age)

	userLen:= len(users)

	fmt.Printf("total users: %d \n",userLen)
	fmt.Println(users)
	 
}