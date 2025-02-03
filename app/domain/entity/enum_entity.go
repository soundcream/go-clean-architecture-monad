package entity

type CustomerStatus Enum
type EmployeeStatus Enum

type Enum struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (e Enum) ToString() string {
	return e.Name
}
