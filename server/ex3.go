package main

import (
	"fmt"
	"io"
	"errors"
	"net/http"
	"encoding/json"
)

type Student struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Grade string `json:"grade"`
}

var students=[]Student{}

func routeHandler(res http.ResponseWriter,req *http.Request){

	method:=req.Method
	path:=req.URL.Path

	fmt.Printf("%s\t%s\n",method,path)

	fmt.Println("===========================")
	 
 

    if method=="GET" {
		res.Header().Set("Content-type","application/json")
		if err:=json.NewEncoder(res).Encode(students);err!=nil{
			
			http.Error(res,"Failed to students",http.StatusInternalServerError)
			return
		}

	}else if method=="POST"  {
		// Limit incoming body to 1024 bytes
        req.Body = http.MaxBytesReader(res, req.Body, 1024)

		body,err:=io.ReadAll(req.Body)
		if err !=nil {
			 var maxErr *http.MaxBytesError

            if errors.As(err, &maxErr) {
                http.Error(res, "Request body too large",
                    http.StatusRequestEntityTooLarge)
                return
            }

			http.Error(res,"Failed to read data",http.StatusInternalServerError)
			return
		}
		var student Student
		err=json.Unmarshal(body,&student)
		fmt.Println(string(body))

		if err!=nil {
			http.Error(res,"failed to parse data",http.StatusInternalServerError)
			return
		}

		students=append(students,student)
		
		res.Header().Set("Content-Type","application/json")
		res.WriteHeader(http.StatusCreated)
		if err=json.NewEncoder(res).Encode(student);err!=nil{
			http.Error(res,"Failed to send data",http.StatusInternalServerError)
		}
	}else{
		http.Error(res,"Page not found",http.StatusBadRequest)
	}

}

func main(){
	http.HandleFunc("/students",routeHandler)
	fmt.Println("Server is listening on http://localhost:8000")

	err := http.ListenAndServe(":8000",nil)

	if err !=nil {
		fmt.Println("Server Error: ",err)
	} 
}