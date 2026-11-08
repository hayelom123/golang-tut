package main

import (
	"fmt"
	"io"
	"net/http"
)

func routeHandler(res http.ResponseWriter, req *http.Request) {

	method := req.Method

	fmt.Println("Method: ", method)

	if method == "POST" {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "Failed to read request body", http.StatusBadRequest)
			return
		}
		task := string(body)
		fmt.Println("body: ", task)

		fmt.Fprintln(res, "Message received:", task)
	} else if method == "GET" {
		fmt.Fprintln(res, "Send a POST request to submit a message")
	} else {
		res.Header().Set("Allow", "GET,POST")
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/messages", routeHandler)

	fmt.Println("server is listening on http://localhost:8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Server error: ", err)
	}
}
