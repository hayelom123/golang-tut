package controllers

import (
	"fmt"
	"net/http"
	"example.com/api/model"
	"example.com/api/helpers"
)

func CreateStudent(res http.ResponseWriter,req *http.Request){
      
	var student model.Student
	
	sucess:= helpers.ReadJson(res,req,&student)

	if !sucess {
		return
	}
	 
	fmt.Printf("Student: %",student)

    helpers.SendJson(res,200,student)

  
	 
 
}

func UpdateStudent(res http.ResponseWriter,req *http.Request){

	
}

func PatchStudent(res http.ResponseWriter,req *http.Request){

	
}

func GetStudentById(res http.ResponseWriter,req *http.Request){

	
}

func GetAllStudents(res http.ResponseWriter,req *http.Request){

	
}