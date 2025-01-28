package entity

// User
// @Description A representation of a user.
// @ID User
type User struct {
	ID       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"John Doe"`
	Username string `json:"username" example:"JohnDoe"`
	Email    string `json:"email" example:"john.doe@example.com"`
}

type CustomerStatus Enum
type EmployeeStatus Enum

type Enum struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (e Enum) String() string {
	return e.Name
}
