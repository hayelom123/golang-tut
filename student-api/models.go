
package main

import "sync"

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
