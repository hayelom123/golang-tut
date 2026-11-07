package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		 if req.Method == http.MethodGet {
			fmt.Fprintf(w,"Get request received!!!\n")
		 }else if req.Method==http.MethodPost {
			fmt.Fprintf(w,"Post request received!!!\n")
		 }else {
			fmt.Fprintf(w,"Unkown request received!!!\n")
		 }
	})


	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
