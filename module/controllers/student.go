package controllers

import (
	"fmt"
	"net/http"
	"example.com/api/model"
	"example.com/api/helpers"
	"example.com/api/middleware"
)

// // An HTTP handler that depends on a StudentStore.
// type StudentHandler struct {
//     store *StudentStore
// }

// // Inject the store into the handler.
// func NewStudentHandler(store *StudentStore) *StudentHandler {
//     return &StudentHandler{
//         store: store,
//     }
// }

func  CreateStudent(res http.ResponseWriter,req *http.Request){
      
	var student model.Student
	
	sucess:= helpers.ReadJson(res,req,&student)
 
    
	 
	if !sucess {
		return
	}
	 
	fmt.Printf("Student: %\n",student)

    helpers.SendJson(res,200,student)

  
	 
 
}

func ProfileHandler(res http.ResponseWriter, req *http.Request) {
    userID, ok := middleware.UserIDFromContext(req.Context())

    if !ok {
        http.Error(res, "User information missing",
            http.StatusInternalServerError)
        return
    }

    fmt.Fprintln(res, "Authenticated user:", userID)
}

func UpdateStudent(res http.ResponseWriter,req *http.Request){

	
}


func PatchStudent(res http.ResponseWriter,req *http.Request){

	
}

func GetStudentById(res http.ResponseWriter,req *http.Request){

	
}

func GetAllStudents(res http.ResponseWriter,req *http.Request){

	
}