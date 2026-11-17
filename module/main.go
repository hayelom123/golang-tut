package main

import (
	"fmt"
	"net/http"
	"context"
    "errors"
    "log"
    "os"
    "os/signal"
    "time"
	"example.com/api/routes"
	"example.com/api/middleware"
	// "os"
)
// https://chatgpt.com/g/g-p-6adfc6759ad08191832519296f77eeb8-learn-go/shared/c/6acc0439-53e8-83eb-9c73-60d542c9315d?owner_user_id=user-ynXtNEfvC2Vc9lAcFVNANt6u
func main(){
   
	// port := os.Getenv("PORT")
    // if port == "" {
    //     port = "8080" // Default fallback
    // }
    port:=8000
	addr:=fmt.Sprintf(":%d",port)

     
	fmt.Printf("Server is runing on http://localost%s\n",addr)
	handler:= middleware.LoggingMiddleware(routes.AppRouter())
 


	server:=&http.Server{
		Addr:addr,
        Handler:handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// err:=http.ListenAndServe(":8000",handler)
	// err:= server.ListenAndServe()


	// if err!=nil {
	// 	fmt.Println("Server Error: "+err.Error())
	// }

	// Step 3: Start server concurrently
    go func() {
        err := server.ListenAndServe()

        if err != nil && !errors.Is(err, http.ErrServerClosed) {
            log.Println("Server error:", err)
        }
    }()

    log.Println("Server running on port 8000")

    // Step 4: Wait for Ctrl+C
    shutdownSignal, stop := signal.NotifyContext(
        context.Background(),
        os.Interrupt,
    )
    defer stop()

    <-shutdownSignal.Done()

    log.Println("Shutdown requested")

    // Step 5: Give requests up to 10 seconds
    ctx, cancel := context.WithTimeout(
        context.Background(),
        10*time.Second,
    )
    defer cancel()

    // Step 6: Stop the server gracefully
    if err := server.Shutdown(ctx); err != nil {
        log.Println("Shutdown error:", err)
    }

    log.Println("Server stopped")
}