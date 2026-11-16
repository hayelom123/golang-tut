package routes

import (
	"example.com/api/controllers"
	"example.com/api/middleware"
	"net/http"
)

func AppRouter() *http.ServeMux {
	app:=http.NewServeMux()
    app.Handle("POST /students/{$}",middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateStudent)))
	app.Handle("GET /students/profile/{$}",middleware.AuthMiddleware(http.HandlerFunc(controllers.ProfileHandler)))
	return app
}