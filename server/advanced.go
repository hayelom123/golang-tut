package main

import (
	"fmt"
	"net/http"
)

func studentHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Welcome to Students API")
}

func coursesHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Welcome to Courses API")
}
func teachersHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Welcome to Teachers API")
}

func productHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Welcome to Products API")
}

func main(){
	mux:=http.NewServeMux()
	
	// Register routes
	mux.HandleFunc("/",func(res http.ResponseWriter,req *http.Request){
		fmt.Fprintln(res,"Hello God from your child")
	})
    mux.HandleFunc("/students", studentHandler)
	 mux.HandleFunc("/courses", coursesHandler)
	  mux.HandleFunc("/teachers", teachersHandler)
    mux.HandleFunc("/products", productHandler)

    fmt.Println("Server is running in http://localhost:8000")
	err:=http.ListenAndServe(":8000",mux)
	if err !=nil {
		fmt.Println("Server error: ",err)
	}
}