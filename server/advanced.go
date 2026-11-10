package main

import (
	"fmt"
	"net/http"
)

func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/",func(res http.ResponseWriter,req *http.Request){
		fmt.Fprintln(res,"Hello God from your child")
	})
    fmt.Println("Server is running in http://localhost:8000")
	err:=http.ListenAndServe(":8000",mux)
	if err !=nil {
		fmt.Println("Server error: ",err)
	}
}