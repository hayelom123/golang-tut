package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"sync"
	"io"
	"strconv"
)

type Student struct{
	ID int      `json:"id"`
	Name string `json:"name"`
	Age int 	`json:"age"`
	Grade string `json:"grade"`
}

// for optional fields
type StudentPatch struct {
    Name  *string `json:"name"`
    Age   *int    `json:"age"`
    Grade *string `json:"grade"`
}


var students = []Student{}
var mux  sync.RWMutex
var nextID int=0

func createSutdent(res http.ResponseWriter,req *http.Request){
	 // Limit incoming body to 1024 bytes
	 req.Body= http.MaxBytesReader(res,req.Body,1024)

     body,err:=io.ReadAll(req.Body) //read 

	if err!=nil {
		http.Error( res,"Failed to read body",http.StatusInternalServerError)
		return
	}

	var student Student
	err= json.Unmarshal(body,&student)

	fmt.Println(string(body))

	if err!=nil {
			http.Error(res,"failed to parse data",http.StatusInternalServerError)
			return
	}
	mux.Lock()

	 
	if student.ID == 0 {
		nextID++
		student.ID=nextID
	}
	students =append(students,student)
   mux.Unlock()

   res.Header().Set("Content-Type","application/json")
   res.WriteHeader(http.StatusCreated)
    

   if err=json.NewEncoder(res).Encode(student);err!=nil{
			http.Error(res,"Failed to send data",http.StatusInternalServerError)
	}
     

}
func getAllSutdents(res http.ResponseWriter,req *http.Request){
	
	res.Header().Set("Content-Type","application/json")
	res.WriteHeader(http.StatusOK)
	if err:=json.NewEncoder(res).Encode(students);err!=nil{
		http.Error(res,"Failed to fetch students",http.StatusInternalServerError)
			return
	}
}
func findStudent(id int)(Student,int,error){
	for index,student:=range students{
       if student.ID == id {
		return student,index,nil
	   }
	}
	return Student{},-1,fmt.Errorf("Student not found")
}
func getStudentById(res http.ResponseWriter,req *http.Request){

	idString:=req.PathValue("id")
	id,err:=strconv.Atoi(idString)

	if err != nil || id <= 0 {
         http.Error(res, "Invalid student ID", http.StatusBadRequest)
        return
		}
	mux.RLock()
	student,_,err := findStudent(id)
	if err !=nil {
		http.Error(res, err.Error(),http.StatusNotFound)
		return
	}
      mux.RUnlock()
	if err=json.NewEncoder(res).Encode(student);err!=nil{
		http.Error(res, err.Error(),http.StatusNotFound)
	}
}
func replaceSutdentInformation(res http.ResponseWriter,req *http.Request){
	idString:=req.PathValue("id")
	id,err:=strconv.Atoi(idString)

	if err != nil || id <= 0 {
         http.Error(res, "Invalid student ID", http.StatusBadRequest)
        return
		}
	 
	_,index,err := findStudent(id)
	if err !=nil {
		http.Error(res, err.Error(),http.StatusNotFound)
		return
	}

	req.Body =http.MaxBytesReader(res,req.Body,1024)
	body,err:=io.ReadAll(req.Body)

	if err!=nil {
		http.Error( res,"Failed to read body",http.StatusInternalServerError)
		return
	}
	 var newInfo StudentPatch
	 

	if err = json.Unmarshal(body, &newInfo); err != nil {
		http.Error(res, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}
  mux.Lock()
	if newInfo.Name !=nil{
		students[index].Name= *newInfo.Name
	}
	if newInfo.Grade != nil{
		students[index].Grade=*newInfo.Grade
	}
	if newInfo.Age !=nil{
		students[index].Age= *newInfo.Age
	}
mux.Unlock()
	if err=json.NewEncoder(res).Encode(students[index]);err!=nil{
		http.Error(res,"Failed to update:"+err.Error(),http.StatusInternalServerError)
	}
}
func updateStudent(res http.ResponseWriter,req *http.Request){
	 idString:=req.PathValue("id")
	id,err:=strconv.Atoi(idString)

	if err != nil || id <= 0 {
         http.Error(res, "Invalid student ID", http.StatusBadRequest)
        return
		}
	 
	student,index,err := findStudent(id)
	if err !=nil {
		http.Error(res, err.Error(),http.StatusNotFound)
		return
	}

	req.Body =http.MaxBytesReader(res,req.Body,1024)
	body,err:=io.ReadAll(req.Body)

	if err!=nil {
		http.Error( res,"Failed to read body",http.StatusInternalServerError)
		return
	}
	 var newInfo Student
	 

	if err = json.Unmarshal(body, &newInfo); err != nil {
		http.Error(res, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}
    mux.Lock()
    newInfo.ID=student.ID
	 students[index] = newInfo
	mux.Unlock()
   res.Header().Set("Content-Type","application/json")
    
	if err=json.NewEncoder(res).Encode(students[index]);err!=nil{
		http.Error(res,"Failed to update:"+err.Error(),http.StatusInternalServerError)
	}
}
func deleteSutdent(res http.ResponseWriter,req *http.Request){
	idString:=req.PathValue("id")
	id,err:=strconv.Atoi(idString)

	if err != nil || id <= 0 {
         http.Error(res, "Invalid student ID", http.StatusBadRequest)
        return
		}
	 
	student,index,err := findStudent(id)
	if err !=nil {
		http.Error(res, err.Error(),http.StatusNotFound)
		return
	}
    mux.Lock()
	students= append(students[:index],students[index+1:]...)
    mux.Unlock()
	 res.Header().Set("Content-Type","application/json")
  
    if err=json.NewEncoder(res).Encode(student);err!=nil{
		http.Error(res,"Failed to update:"+err.Error(),http.StatusInternalServerError)
	}
}

func main(){
  app:=http.NewServeMux()

  // routes
  app.HandleFunc("GET /students",getAllSutdents)
  app.HandleFunc("GET /students/{id}",getStudentById)
  app.HandleFunc("POST /students",createSutdent)

  app.HandleFunc("PUT /students/{id}",updateStudent)
  app.HandleFunc("PATCH /students/{id}",replaceSutdentInformation)
  app.HandleFunc("DELETE /students/{id}",deleteSutdent)
 
  fmt.Println("Server is running in http://localhost:8000")

  // listen
  err:=http.ListenAndServe(":8000",app)

  if err !=nil {
	fmt.Println("server error: ",err)
  }

}