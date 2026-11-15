package middleware

import (
	"log"
	"net/http"
	"time"
	"fmt"
)


func LoggingMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(res http.ResponseWriter,req *http.Request){
		start := time.Now()
		
		// 1. Pre-processing logic runs here (before the main handler)
		log.Printf("Started %s %s", req.Method, req.URL.Path)


		// 2. Pass control to the next handler in the chain
		next.ServeHTTP(res,req)

		// 3. Post-processing logic runs here (after the main handler finishes)
		log.Printf("\n%s %s completed in %v\n", req.Method, req.URL.Path, time.Since(start))
	})
}