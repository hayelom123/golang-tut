package main

import "fmt"


func main(){
  result := sum(7,8,9,10)
  fmt.Printf("Sum of 7 and 8 is %d\n",result)

  sum,diff := calculate(9,6)

  fmt.Println("Sum:", sum)
  fmt.Println("Difference:", diff)

  func(){
    fmt.Println("I am immidately invoked function")
  }()// it will be used for asyn and defer function in the future

  mul:=func(numbers ...int)(result int){
        result=1
        for _,num:= range numbers{
          result*=num
        }
        return
  }

  fmt.Println("Multplication Result:",mul(1,2,3,4))


}

func sum(a ...int) (result int) {
   for _,number:= range a {
       result+=number
   }
	return 
}

// multiple return

func calculate(a,b int)(int,int){
	sum := a+b
	difference := a-b
	return sum,difference
}