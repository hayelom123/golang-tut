package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func processHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing started...")

	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Fprintln(w, "Drawing processed!")
		fmt.Println("Processing completed")

	case <-r.Context().Done():
		fmt.Println("Processing canceled:", r.Context().Err())
		return
	}
}

func main() {
	http.HandleFunc("/process", processHandler)

	log.Println("Server running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
