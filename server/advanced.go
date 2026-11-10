package main

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)


type Student struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Grade string `json:"grade"`
}


var students = []Student{
    {ID: 1, Name: "Hayelom", Age: 25, Grade: "A"},
    {ID: 2, Name: "Sara", Age: 23, Grade: "B"},
    {ID: 3, Name: "Abel", Age: 24, Grade: "A"},
}


func studentHandler(res http.ResponseWriter, req *http.Request) {
	   data, err := json.Marshal(students)

    if err != nil {
        http.Error(res, "Failed to encode students",
            http.StatusInternalServerError)
        return
    }

    res.Header().Set("Content-Type", "application/json")

    if _, err := res.Write(data); err != nil {
        fmt.Println("Response error:", err)
    }
}
func createStudentHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusCreated)
    fmt.Fprintln(res, "Welcome create to Students API")
}
func getStudentByIdHandler(res http.ResponseWriter, req *http.Request) {
	 idString:=req.PathValue("id")
	 id,err:=strconv.Atoi(idString)

    // fmt.Fprintln(res, "student id: ", id)
	if err != nil || id <= 0 {
    http.Error(res, "Invalid student ID", http.StatusBadRequest)
    return
		}
	for _,student:= range students {
		if student.ID==id{
			res.Header().Set("Content-Type","application/json")
			json.NewEncoder(res).Encode(student)
			return
		}
	}
	http.NotFound(res, req)
}

func coursesHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Welcome to Courses API")
}
func teachersHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Welcome to Teachers API")
}

func productHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Welcome to Products API")
}

func createProductHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Welcome to Create Products API")
}
func main(){
	mux:=http.NewServeMux()
	
	// Register routes
	mux.HandleFunc("/{$}",func(res http.ResponseWriter,req *http.Request){
		fmt.Fprintln(res,"Hello God from your child")
	})
    mux.HandleFunc("GET /students", studentHandler)
	mux.HandleFunc("POST /students", createStudentHandler)
	mux.HandleFunc("GET /students/{id}", getStudentByIdHandler)
	mux.HandleFunc("GET /courses", coursesHandler)
	mux.HandleFunc("GET /teachers", teachersHandler)
    mux.HandleFunc("GET /products", productHandler)
	mux.HandleFunc("POST /products/create",createProductHandler)

    fmt.Println("Server is running in http://localhost:8000")
	err:=http.ListenAndServe(":8000",mux)
	if err !=nil {
		fmt.Println("Server error: ",err)
	}
}