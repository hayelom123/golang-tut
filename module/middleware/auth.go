package middleware

import ("net/http")

// Private context key type
type userIDKey struct{}
// Retrieve the authenticated user ID
func UserIDFromContext(ctx context.Context) (string, bool) {
    id, ok := ctx.Value(userIDKey{}).(string)
    return id, ok
}


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
			return
		}
		
		// Demo identity associated with our practice token
        userID := "student-42"

		// Add user ID to the request context
        ctx := context.WithValue(
            req.Context(),
            userIDKey{},
            userID,
        )


		// Continue only when token is valid
		// next.ServeHTTP(res,req)

		// Forward the request with its updated context
        next.ServeHTTP(res, req.WithContext(ctx))
	})
}