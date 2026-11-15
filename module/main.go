package main

import (
	"fmt"
	"net/http"
	"example.com/api/routes"
	"example.com/api/middleware"
)
// https://chatgpt.com/g/g-p-6adfc6759ad08191832519296f77eeb8-learn-go/shared/c/6acc0439-53e8-83eb-9c73-60d542c9315d?owner_user_id=user-ynXtNEfvC2Vc9lAcFVNANt6u
func main(){
     
	fmt.Println("Server is runing on http://localost:8000")
	handler:= middleware.LoggingMiddleware(routes.AppRouter())

	err:=http.ListenAndServe(":8000",handler)

	if err!=nil {
		fmt.Println("Server Error: "+err.Error())
	}
}