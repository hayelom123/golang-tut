package main

import "fmt"
// https://chatgpt.com/share/6adfc9ec-6dd8-83eb-b4b8-19b8888d78c9
// https://chatgpt.com/g/g-p-6adfc6759ad08191832519296f77eeb8-learn-go/project
func main(){

	fmt.Println("Hello, God! sorry for everyhting!!!")

	name:="Hayelom"
	age := 29
	salary := 20000.0

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(salary)

	var name1 string = "Daniel" //Method 1: Explicit type
	var age1 =27  // Method 2: Type inference
	salary1 :=  6000.50 // Method 3: Short declaration
    isEmployed := true
	fmt.Println(name1,age1,salary1)

	fmt.Printf("Name: %s\n",name1)
	fmt.Printf("Age: %d\n",age1)
	fmt.Printf("Salary: %.2f\n",salary1)
	fmt.Printf("Employed: %t\n",isEmployed)
}