package main

import "fmt"


func main(){
	var myVal any = 67
	
	newval,ok:=myVal.(int)

	if !ok {
		fmt.Println("value is not int")
          return
	}
   
	fmt.Println(newval)
}