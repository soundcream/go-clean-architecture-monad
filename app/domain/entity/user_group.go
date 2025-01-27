package entity

type UserGroup struct {
	Name string `json:"name" example:"GOLD"`
}

type UserTest struct {
	Name string `validate:"required,min=5,max=20"` // Required field, min 5 char long max 20
	Age  int    `validate:"required,teener"`       // Required field, and client needs to implement our 'teener' tag format which we'll see later

}
