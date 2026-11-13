package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

type StudentPatch struct {
	Name  *string `json:"name"`
	Age   *int    `json:"age"`
	Grade *string `json:"grade"`
}

var students = []Student{}
var nextID int
var mu sync.RWMutex

// Call this only while holding mu.Lock() or mu.RLock().
func findStudentIndex(id int) int {
	for i := range students {
		if students[i].ID == id {
			return i
		}
	}
	return -1
}

func parseStudentID(res http.ResponseWriter, req *http.Request) (int, bool) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil || id <= 0 {
		http.Error(res, "Invalid student ID", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func readJSON(res http.ResponseWriter, req *http.Request, destination any) bool {
	mediaType, _, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		http.Error(res, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return false
	}

	req.Body = http.MaxBytesReader(res, req.Body, 1024)
	body, err := io.ReadAll(req.Body)
	if err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			http.Error(res, "Request body too large", http.StatusRequestEntityTooLarge)
		} else {
			http.Error(res, "Cannot read request body", http.StatusBadRequest)
		}
		return false
	}

	if err := json.Unmarshal(body, destination); err != nil {
		http.Error(res, "Invalid JSON", http.StatusBadRequest)
		return false
	}
	return true
}

func writeJSON(res http.ResponseWriter, status int, value any) {
	data, err := json.Marshal(value)
	if err != nil {
		http.Error(res, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	if _, err := res.Write(data); err != nil {
		fmt.Println("Response write error:", err)
	}
}

func validStudent(s *Student) bool {
	s.Name = strings.TrimSpace(s.Name)
	s.Grade = strings.TrimSpace(s.Grade)
	return s.Name != "" && s.Age > 0 && s.Grade != ""
}

func getAllStudents(res http.ResponseWriter, req *http.Request) {
	mu.RLock()
	result := slices.Clone(students)
	mu.RUnlock()
	writeJSON(res, http.StatusOK, result)
}

func getStudentByID(res http.ResponseWriter, req *http.Request) {
	id, ok := parseStudentID(res, req)
	if !ok {
		return
	}

	mu.RLock()
	index := findStudentIndex(id)
	if index == -1 {
		mu.RUnlock()
		http.NotFound(res, req)
		return
	}
	student := students[index]
	mu.RUnlock()

	writeJSON(res, http.StatusOK, student)
}

func createStudent(res http.ResponseWriter, req *http.Request) {
	var student Student
	if !readJSON(res, req, &student) {
		return
	}
	if !validStudent(&student) {
		http.Error(res, "Invalid student data", http.StatusBadRequest)
		return
	}

	mu.Lock()
	nextID++
	student.ID = nextID // Never trust an ID supplied by the client.
	students = append(students, student)
	mu.Unlock()

	res.Header().Set("Location", fmt.Sprintf("/students/%d", student.ID))
	writeJSON(res, http.StatusCreated, student)
}

func patchStudent(res http.ResponseWriter, req *http.Request) {
	id, ok := parseStudentID(res, req)
	if !ok {
		return
	}
	var changes StudentPatch
	if !readJSON(res, req, &changes) {
		return
	}

	if changes.Name == nil && changes.Age == nil && changes.Grade == nil {
		http.Error(res, "No changes provided", http.StatusBadRequest)
		return
	}
	if changes.Name != nil {
		*changes.Name = strings.TrimSpace(*changes.Name)
		if *changes.Name == "" {
			http.Error(res, "Invalid name", http.StatusBadRequest)
			return
		}
	}
	if changes.Age != nil && *changes.Age <= 0 {
		http.Error(res, "Invalid age", http.StatusBadRequest)
		return
	}
	if changes.Grade != nil {
		*changes.Grade = strings.TrimSpace(*changes.Grade)
		if *changes.Grade == "" {
			http.Error(res, "Invalid grade", http.StatusBadRequest)
			return
		}
	}

	mu.Lock()
	index := findStudentIndex(id)
	if index == -1 {
		mu.Unlock()
		http.NotFound(res, req)
		return
	}
	if changes.Name != nil {
		students[index].Name = *changes.Name
	}
	if changes.Age != nil {
		students[index].Age = *changes.Age
	}
	if changes.Grade != nil {
		students[index].Grade = *changes.Grade
	}
	updated := students[index]
	mu.Unlock()

	writeJSON(res, http.StatusOK, updated)
}

func updateStudent(res http.ResponseWriter, req *http.Request) {
	id, ok := parseStudentID(res, req)
	if !ok {
		return
	}

	var updated Student
	if !readJSON(res, req, &updated) {
		return
	}
	if !validStudent(&updated) {
		http.Error(res, "Invalid student data", http.StatusBadRequest)
		return
	}
	updated.ID = id

	mu.Lock()
	index := findStudentIndex(id)
	if index == -1 {
		mu.Unlock()
		http.NotFound(res, req)
		return
	}
	students[index] = updated
	mu.Unlock()

	writeJSON(res, http.StatusOK, updated)
}

func deleteStudent(res http.ResponseWriter, req *http.Request) {
	id, ok := parseStudentID(res, req)
	if !ok {
		return
	}

	mu.Lock()
	index := findStudentIndex(id)
	if index == -1 {
		mu.Unlock()
		http.NotFound(res, req)
		return
	}
	students = append(students[:index], students[index+1:]...)
	mu.Unlock()

	res.WriteHeader(http.StatusNoContent)
}

func newRouter() *http.ServeMux {
	app := http.NewServeMux()
	app.HandleFunc("GET /students", getAllStudents)
	app.HandleFunc("GET /students/{id}", getStudentByID)
	app.HandleFunc("POST /students", createStudent)
	app.HandleFunc("PUT /students/{id}", updateStudent)
	app.HandleFunc("PATCH /students/{id}", patchStudent)
	app.HandleFunc("DELETE /students/{id}", deleteStudent)
	return app
}

func main() {
	fmt.Println("Server running at http://localhost:8000")
	if err := http.ListenAndServe(":8000", newRouter()); err != nil {
		fmt.Println("Server error:", err)
	}
}
