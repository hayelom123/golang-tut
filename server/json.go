package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Task struct {
	Title     string `json:"title"`
	Priority  int    `json:"priority"`
	Completed bool   `json:"completed"`
}

var allTasks = []Task{}

func routeHandler(res http.ResponseWriter, req *http.Request) {
	method := req.Method

	if method == "GET" {

		// Set Content-Type header to tell the client it's receiving JSON
		res.Header().Set("Content-Type", "application/json")

		// FIX 3: Capture the encoded JSON bytes and error
		jsonData, err := json.Marshal(allTasks)
		if err != nil {
			http.Error(res, "Failed to encode tasks", http.StatusInternalServerError)
			return
		}

		// Write the JSON response directly to the client
		res.Write(jsonData)

	} else if method == "POST" {

		body, err := io.ReadAll(req.Body)

		if err != nil {

			http.Error(res, "Failed to read request body", http.StatusBadRequest)
			return
		}

		var task Task
		err = json.Unmarshal(body, &task)
		if err != nil {
			http.Error(res, "Invalid JSON", http.StatusBadRequest)
			return
		}
		// FIX 2: Append the newly created task to your global slice
		allTasks = append(allTasks, task)
		fmt.Println("Title:", task.Title)
		fmt.Println("Priority:", task.Priority)
		fmt.Println("Completed:", task.Completed)

		 // Send success response back to client
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(map[string]string{"message": "Task received successfully"})

	} else {
		res.Header().Set("Allow", "GET,POST")
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/tasks", routeHandler)

	fmt.Println("Server running at http://localhost:8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
