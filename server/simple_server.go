package main

import (
	"fmt"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {

	fmt.Println("Method:", req.Method)
	fmt.Println("Path:", req.URL.Path)
	fmt.Println("headers:",req.Header)

	fmt.Fprintln(res, "Hello God!!!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Server started on http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Server error: ", err)
	}
}
