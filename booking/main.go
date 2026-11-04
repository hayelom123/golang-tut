package main

import ("fmt"
 "booking/utils"
 "booking/models"
)

type User struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

func main() {
	const conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	var bookings = []User{}

	fmt.Printf("Wellcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	user,err := getUserInput()
	if err!=nil {
		fmt.Println("Failed to register user")
		return
	}
	remainingTickets -= user.tickets

	bookings = append(bookings, user)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conformation email at %v\n", user.firstName, user.lastName, user.tickets, user.email)
	//   fmt.Printf("User: %v\n",user)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("total bookings: %v \n", len(bookings))

	fmt.Println("\n\n========================Bank==============================")

    acc,err := models.NewBankAccount(user.firstName,50000.0) 

	if err !=nil {
       fmt.Println("Error:", err) // Output: Error: initial deposit cannot be negative
	}
   fmt.Printf("User balance: %.2f\n",acc.GetBalance())

	acc.Deposit(20000)
	fmt.Printf("User balance after Deposit: %.2f\n",acc.GetBalance())
 
	acc.Withdraw(30000)
	fmt.Printf("User balance after Withdraw: %.2f\n",acc.GetBalance())


}

func getUserInput() (User,error) {
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

	isValidName, isValidEmail, isValidTicket := utils.ValidateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicket {
		return User{
			firstName,
			lastName,
			email,
			userTickets,
		},nil
	} else {
		if !isValidName {
			fmt.Println("Invalid first name or last name and each must be 2 letters")
		}
		if !isValidEmail {
			fmt.Println("Invalid email. Email must contain @")
		}
		if !isValidTicket {
			fmt.Println("Invalid tickets. Ticket must be 1 or more")
		}
		return User{},fmt.Errorf("Invalid Input")
	}

}


