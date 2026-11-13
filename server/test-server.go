package main

import (
	"fmt"
	"net/http"
)

func main(){
	app:=http.NewServeMux()

	app.HandleFunc("GET /{$}",func(res http.ResponseWriter,req *http.Request){
		path:=req.URL.Path
		method:=req.Method
		fmt.Printf("%s\t%s\n",method,path)
		fmt.Fprintln(res,"Home page")
	})
	app.HandleFunc("DELETE /{$}",func(res http.ResponseWriter,req *http.Request){
		path:=req.URL.Path
		method:=req.Method
		fmt.Printf("%s\t%s\n",method,path)
		fmt.Fprintln(res,"Delete page")
	})
    fmt.Println("Server is running in http://localhost:8000")
	err:=http.ListenAndServe(":8000",app)

	if err != nil {
		fmt.Println("server error: ",err)
	}
}