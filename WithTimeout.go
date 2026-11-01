package main

import (
	"context"
	"fmt"
	"time"
)

// slowDatabaseQuery simulates a database call that listens for cancellation
func slowDatabaseQuery(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second): // Simulates a 3-second query
		fmt.Println("Query finished successfully!")
	case <-ctx.Done(): // Triggered if the context is canceled or times out
		fmt.Println("Query aborted:", ctx.Err())
	}
}

func main() {
	// Create a root context
	rootCtx := context.Background()

	// Create a derived context that times out automatically after 1 second
	ctx, cancel := context.WithTimeout(rootCtx, 1*time.Second)
	
	// Always call cancel() to release system resources when finished
	defer cancel() 

	// Start the database query
	slowDatabaseQuery(ctx)
}
