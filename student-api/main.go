
package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Server running at http://localhost:8000")

    err := http.ListenAndServe(":8000", newRouter())

    if err != nil {
        fmt.Println("Server error:", err)
    }
}
