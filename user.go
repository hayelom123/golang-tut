package main

import ("errors" 
"fmt")

type User struct{
	id int
	name string
	age  int
}


func getUser(id int, users []User) (User,error){

    for _, user := range users{
      if(user.id==id){
		return user,nil
	  }
	}
	return User{},errors.New("User not found")
}

func main(){
   users := []User{}

	user1 := User{
		id: 0,
		name: "Hayelom Kiros",
		age: 29,
	}


	users = append(users,user1)
     users = append(users,User{
		id:1,
		name: "Kassa Kiros",
		age:32,
	 })
	fmt.Printf("%s age %d \n",user1.name,user1.age)

	userLen:= len(users)

	fmt.Printf("total users: %d \n",userLen)
	fmt.Println(users)

	user,err := getUser(0,users)

	if err != nil {
		fmt.Println("User not found")
	}else{
		 fmt.Println(user)
	}
	 
}