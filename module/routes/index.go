package routes

import (
	"example.com/api/controllers"
	"net/http"
)

func AppRouter() *http.ServeMux {
	app:=http.NewServeMux()
    app.HandleFunc("POST /students/{$}",controllers.CreateStudent)
	return app
}