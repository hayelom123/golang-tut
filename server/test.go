package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)


type User struct{
	Name string `json:"name"`
	Id  string `json:"id"`
}

var users =[]User{}
func routeHandler(res http.ResponseWriter, req *http.Request){

	method:=req.Method
	path:=req.URL.Path

	fmt.Printf("%s\t%s\n",method,path)

	if  method=="GET" && path=="/" {
		 
		res.Header().Set("Content-type","application/json")
		res.WriteHeader(http.StatusOK)
		if err:=json.NewEncoder(res).Encode(users);err!=nil{
			http.Error(res,"Failed to fetch data",http.StatusInternalServerError)
		}
		return
	}else if method=="POST" && path=="/"{
      body,err:=io.ReadAll(req.Body)
	  if err!=nil{
         http.Error(res,"failed to read data",http.StatusInternalServerError)
		return
	  }
     var user User
	 err=json.Unmarshal(body,&user)
	if err!=nil{
         http.Error(res,"validatio error",http.StatusInternalServerError)
		return
	  }
	users= append(users,user)
	res.Header().Set("Content-Type","application/json")

	  res.WriteHeader(http.StatusCreated)
	//   data,_err:=json.Marshal(users)
	 if err=json.NewEncoder(res).Encode(user); err!=nil{
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	 }

	}else{
		http.Error(res,"Page not found",http.StatusBadRequest)
	}
}

func main(){
      
	  http.HandleFunc("/",routeHandler)
      fmt.Println("http://localhost:8000")
	  err:=http.ListenAndServe(":8000",nil)

	  if err!=nil {
		fmt.Println("Server Error: ",err)
	  }


}