
package main

import (
    "fmt"
    "net/http"
    "slices"
)

func getAllStudents(res http.ResponseWriter, req *http.Request) {
    mu.RLock()
    result := slices.Clone(students)
    mu.RUnlock()

    writeJSON(res, http.StatusOK, result)
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
    student.ID = nextID
    students = append(students, student)

    mu.Unlock()

    res.Header().Set("Location", fmt.Sprintf("/students/%d", student.ID))
    writeJSON(res, http.StatusCreated, student)
}
