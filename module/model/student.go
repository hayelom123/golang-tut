package model

 

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
