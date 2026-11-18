package routes

import (
	"example.com/api/controllers"
	"example.com/api/middleware"
	"example.com/api/helpers"
	"net/http"
)

func AppRouter() *http.ServeMux {
	app:=http.NewServeMux()
    app.Handle("POST /students/{$}",middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateStudent)))
	app.Handle("GET /students/profile/{$}",middleware.AuthMiddleware(http.HandlerFunc(controllers.ProfileHandler)))
	app.HandleFunc("/students/", func(
    res http.ResponseWriter,
    req *http.Request,
) {
    helpers.WriteJSONError(
        res,
        http.StatusNotFound,
        "Route not found",
    )
})
	return app
}