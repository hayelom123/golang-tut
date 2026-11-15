package main

import (
	"fmt"
	"net/http"
	"example.com/api/routes"
	"example.com/api/middleware"
)

func main(){
     
	fmt.Println("Server is runing on http://localost:8000")
	handler:= middleware.LoggingMiddleware(routes.AppRouter())

	err:=http.ListenAndServe(":8000",handler)

	if err!=nil {
		fmt.Println("Server Error: "+err.Error())
	}
}