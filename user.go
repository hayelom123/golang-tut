package main

import ( "errors"
"fmt")

type User struct{
	id int
	name string
	age  int
}


func withdraw(balance float64,amount float64) (float64,error) {

    if amount <=0{
        return balance,errors.New("amount must be positive")
    }

    if amount > balance {
        return balance,fmt.Errorf("Cannot withdraw %.2f: available balance is %.2f",amount,balance)
    }

    return balance-amount, nil
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


     remaining, err := withdraw(100, 300)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Remaining balance:", remaining)
}