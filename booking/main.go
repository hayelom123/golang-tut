package main

import "fmt"


func main(){
	const conferenceName= "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 20

	fmt.Printf("Wellcome to %v booking application\n",conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend.")

      var userName string
	  var userTickets int
	  // ask user for their name

	   fmt.Scan(&userName)
	 
	//   userName="Tom"
	  userTickets=2
	  fmt.Printf("User %v booked %v tickets.\n",userName,userTickets)
}