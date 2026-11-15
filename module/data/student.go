package store


type StudentStore struct {
    students []models.Student
    nextID   int
    mu       sync.RWMutex
}
func NewStudentStore() *StudentStore {
    return &StudentStore{
        students: []models.Student{},
    }
}