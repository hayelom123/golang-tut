
package main

import "net/http"

func newRouter() *http.ServeMux {
    app := http.NewServeMux()

    app.HandleFunc("GET /students", getAllStudents)
    app.HandleFunc("POST /students", createStudent)

    return app
}
