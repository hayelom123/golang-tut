package main

import "fmt"

type User struct {
   	firstName string
    lastName string
    email string
    tickets uint
}
func main(){
	const conferenceName= "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	var bookings =[]User{}

	fmt.Printf("Wellcome to %v booking application\n",conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend.")

      var firstName string
	  var lastName string
	  var email string
	  var userTickets uint
	  // ask user for their name
       
	   fmt.Printf("First Name: ")
	   fmt.Scan(&firstName)
	
	   fmt.Printf("Last name: ")
	   fmt.Scan(&lastName)
     
	   fmt.Printf("Email: ")
	   fmt.Scan(&email)

	   fmt.Printf("Enter number of tickets: ")
	   fmt.Scan(&userTickets)

	  remainingTickets -= userTickets
	var user = User{
		firstName,
		lastName,
		email,
		userTickets,
	  }
	  user.Name=firstName
	bookings=  append(bookings,user)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conformation email at %v\n",firstName,lastName,userTickets,email)
	//   fmt.Printf("User: %v\n",user)
	
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)

	 fmt.Printf("total bookings: %v \n",len(bookings))
	
}