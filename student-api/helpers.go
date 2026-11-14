
package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "mime"
    "net/http"
    "strings"
)

func readJSON(res http.ResponseWriter, req *http.Request, destination any) bool {
    mediaType, _, err := mime.ParseMediaType(req.Header.Get("Content-Type"))

    if err != nil || mediaType != "application/json" {
        http.Error(res, "JSON required", http.StatusUnsupportedMediaType)
        return false
    }

    req.Body = http.MaxBytesReader(res, req.Body, 1024)

    body, err := io.ReadAll(req.Body)

    if err != nil {
        var maxErr *http.MaxBytesError

        if errors.As(err, &maxErr) {
            http.Error(res, "Body too large", http.StatusRequestEntityTooLarge)
        } else {
            http.Error(res, "Cannot read body", http.StatusBadRequest)
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
        http.Error(res, "JSON encoding failed", http.StatusInternalServerError)
        return
    }

    res.Header().Set("Content-Type", "application/json")
    res.WriteHeader(status)

    if _, err := res.Write(data); err != nil {
        fmt.Println("Response error:", err)
    }
}

func validStudent(s *Student) bool {
    s.Name = strings.TrimSpace(s.Name)
    s.Grade = strings.TrimSpace(s.Grade)

    return s.Name != "" && s.Age > 0 && s.Grade != ""
}
