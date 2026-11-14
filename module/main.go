package main

import (
	"fmt"
	"net/http"
	"example.com/api/routes"
)

func main(){
     
	fmt.Println("Server is runing on http://localost:8000")

	err:=http.ListenAndServe(":8000",routes.AppRouter())

	if err!=nil {
		fmt.Println("Server Error: "+err.Error())
	}
}