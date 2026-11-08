package main

import (
	"fmt"
	"net/http"
)

func routeHandler(res http.ResponseWriter, req *http.Request) {

	method := req.Method
	path := req.URL.Path

	fmt.Printf("%v\t%v\n", method, path)

	if path == "/" {
		fmt.Fprintln(res, "Welcome to Mel Architects")
	} else if path == "/services" {
		fmt.Fprintln(res, "Our services: Architecture and Interior Design")
	} else if path == "/portfolio" {
		fmt.Fprintln(res, "Explore our architectural projects")
	} else {
		http.NotFound(res, req)
	}
}


func main(){
	http.HandleFunc("/",routeHandler)

	fmt.Println("server is listening on http://localhost:8000")

	err:=http.ListenAndServe(":8000",nil)

	if err !=nil {
		fmt.Println("Server error: ",err)
	}
}