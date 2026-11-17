package main

import (
	"fmt"
	"net/http"
	"example.com/api/routes"
	"example.com/api/middleware"
	// "os"
)
// https://chatgpt.com/g/g-p-6adfc6759ad08191832519296f77eeb8-learn-go/shared/c/6acc0439-53e8-83eb-9c73-60d542c9315d?owner_user_id=user-ynXtNEfvC2Vc9lAcFVNANt6u
func main(){
   
	// port := os.Getenv("PORT")
    // if port == "" {
    //     port = "8080" // Default fallback
    // }
    port:=8000
	addr:=fmt.Sprintf(":%d",port)

     
	fmt.Printf("Server is runing on http://localost%s\n",addr)
	handler:= middleware.LoggingMiddleware(routes.AppRouter())
 


	server:=&http.Server{
		Addr:addr,
        Handler:handler,
	}

	// err:=http.ListenAndServe(":8000",handler)
	err:= server.ListenAndServe()


	if err!=nil {
		fmt.Println("Server Error: "+err.Error())
	}
}