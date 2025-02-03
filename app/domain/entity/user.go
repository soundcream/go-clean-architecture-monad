package entity

type UserTx struct {
	BaseEntity
}

// User
// @Description A representation of a user.
// @ID User
type User struct {
	BaseEntity `table-name:"users"`
	Name       string `column:"name" json:"name" example:"John Doe"`
	Username   string `column:"username" json:"username" example:"JohnDoe"`
	Email      string `column:"email" json:"email" example:"john.doe@example.com"`
	Point      *int   `column:"point" json:"point" example:"0"`
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
