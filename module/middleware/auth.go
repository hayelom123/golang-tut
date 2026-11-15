package middleware

import ("net/http")

func AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(
		res http.ResponseWriter,
		req *http.Request,
	){

		// get header token
		token:= req.Header.Get("Authorization")

		if token!="Bearer parctice-token" {
			res.Header().Set("WWW-Authenticate", "Bearer")
			http.Error(res,"Unauthorized",http.StatusUnauthorized)
		}
		
		// Continue only when token is valid
		next.ServeHTTP(res,req)
	})
}