package models

// import "errors"


// User is capitalized, so it is exported.
type User struct{
  FirstName string // Exported field
  LastName string // Exported field
  role string // Unexported field (lowercase) - other packages cannot see this!
}

//NewUser acts as a constructor. 
//// It validates inputs before returning a successfully initialized pointer.
func NewUser(firstName,string,lastName string,)(User,error){
	   
       return User{
		FirstName:firstName,
		LastName:lastName,
		role:"User",
	   },nil
}


func (user *User) setRole(role string) {
     user.role=role
}

