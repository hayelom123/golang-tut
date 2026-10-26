package main

import "fmt"


func main(){

	 age := 27

	 if(age >= 18){
		        fmt.Println("You are an adult")
	 }else{
		 fmt.Println("You are a minor")
	 }



score := 95

if score >= 90 {
    fmt.Println("Excellent")
} else if score >= 75 {
    fmt.Println("Very Good")
} else if score >= 50 {
    fmt.Println("Pass")
} else {
    fmt.Println("Fail")
}

   day := "Monday"
  switch day {
    case "Monday":
        fmt.Println("Start of work week")
    case "Friday":
        fmt.Println("Almost weekend")
    case "Saturday", "Sunday":
        fmt.Println("Weekend")
    default:
        fmt.Println("Normal working day")
    }


    username := "alex"
    age2 := 25
    isActive := true

    if username == "" {
        fmt.Println("Username is required")
    } else if age2 < 18 {
        fmt.Println("User must be 18 or older")
    } else if !isActive {
        fmt.Println("Account is inactive")
    } else {
        fmt.Println("Access granted")
    }


}