package main

import "fmt"


func main(){
  result := sum(7,8)
  fmt.Printf("Sum of 7 and 8 is %d\n",result)

  sum,diff := calculate(9,6)

      fmt.Println("Sum:", sum)
    fmt.Println("Difference:", diff)
}

func sum(a int, b int) int {
	return a + b
}

// multiple return

func calculate(a,b int)(int,int){
	sum := a+b
	difference := a-b
	return sum,difference
}