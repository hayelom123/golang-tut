package main

import "fmt"

func main(){
	basicSalary := 5000.0
	allowance := 1500.0
	deduction := 300.0

	totalSalary := basicSalary + allowance
	netSalary := totalSalary - deduction

	fmt.Printf("Basic Salary: %.2f AED\n", basicSalary)
    fmt.Printf("Total Salary: %.2f AED\n", totalSalary)
    fmt.Printf("Net Salary: %.2f AED\n", netSalary)
}