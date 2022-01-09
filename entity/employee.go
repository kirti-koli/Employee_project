package entity

//Person object for REST(CRUD)
type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}
type Employee struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Salary     string `json:"salary"`
	Department string `json:"department"`
}
