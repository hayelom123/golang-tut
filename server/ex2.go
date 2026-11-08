
package main

import (
    "fmt"
    "net/http"
)

func taskHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Println("Method:", req.Method)
    fmt.Println("Path:", req.URL.Path)

    if req.Method == "GET" {
        fmt.Fprintln(res, "Here are your tasks")

    } else if req.Method == "POST" {
        fmt.Fprintln(res, "New task received")

    } else {
        res.Header().Set("Allow", "GET, POST")
        http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/tasks", taskHandler)

    fmt.Println("Server running at http://localhost:8000")

    err := http.ListenAndServe(":8000", nil)

    if err != nil {
        fmt.Println("Server error:", err)
    }
}
