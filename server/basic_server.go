package main

import (
	"fmt"
	"net/http"
)

func routeHandler(res http.ResponseWriter,req *http.Request){
    method:=req.Method
	path:=req.URL.Path
     fmt.Printf("%v\t%v\n",method,path)
	if path =="/" {
fmt.Fprintln(res, "Welcome to our architecture website!")
	}else if path=="/about"{
fmt.Fprintln(res, "About our architecture company")
	}else if path =="/projects"{
 fmt.Fprintln(res, "Our architectural projects")
	}else{
		http.NotFound(res,req)
	}
}

func main(){

	http.HandleFunc("/",routeHandler)

	fmt.Println("Server is listening on http://localhost:8000")

	err :=http.ListenAndServe(":8000",nil)

	if err!=nil{
		fmt.Println("Server error",err)
	}
}