package main

import (  
"fmt")

type User struct{
	id int
	name string
	age  int
}


func addUser(user User, users []User) []User {
    
	 return append(users,user)
}

func deleteUser(id int, users []User) ([]User, error) {
	 
	for index,user := range users{
		if(user.id==id){
		 users = append(users[:index],users[index+1:]...)
           return users,nil
		}
	}

    return users,fmt.Errorf("User not found with id %d",id)
}

func updateUser(id int, updatedUser User, users []User) ([]User, error) {
	  for index,user := range users{
		if(user.id==id){
			users[index] =updatedUser
			return users,nil;
		}

	  }
       return users,fmt.Errorf("User not found with ID %d",id)
}

func getUser(id int, users []User) (User,error){

    for _, user := range users{
      if(user.id==id){
		return user,nil
	  }
	}
	return User{},fmt.Errorf("user with ID %d not found", id)
}

func main() {
    users := []User{
        {id: 0, name: "Hayelom Kiros", age: 29},
        {id: 1, name: "Kassa Kiros", age: 32},
    }

    // CREATE
    users = addUser(
        User{id: 2, name: "Kinfe michael Kiros", age: 27},
        users,
    )

    fmt.Println("After adding:", users)

    // UPDATE
    var err error

    users, err = updateUser(
        2,
        User{id: 2, name: "Kinfe Michael Kiros", age: 28},
        users,
    )

    if err != nil {
        fmt.Println("Update failed:", err)
    } else {
        fmt.Println("After updating:", users)
    }

    // DELETE
    users, err = deleteUser(2, users)

    if err != nil {
        fmt.Println("Delete failed:", err)
    } else {
        fmt.Println("After deleting:", users)
    }

    // READ
    user, err := getUser(2, users)

    if err != nil {
        fmt.Println("Get user failed:", err)
    } else {
        fmt.Println("User found:", user)
    }

    fmt.Println("Total users:", len(users))
}