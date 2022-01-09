package entity

//employee 
type Employee struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Salary     string `json:"salary"`
	Department string `json:"department"`
}
