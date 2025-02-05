package entity

// User
// @Description A representation of a user.
// @ID User
type User struct {
	BaseEntity
	Name        string     `column:"name" json:"name" example:"John Doe"`
	Username    string     `column:"username" json:"username" example:"JohnDoe"`
	Email       string     `column:"email" json:"email" example:"john.doe@example.com"`
	Point       *int       `column:"point" json:"point" example:"0"`
	UserGroupId *int       `column:"user_group_id" json:"userGroupId" example:"0"`
	UserGroup   *UserGroup `json:"userGroup" gorm:"foreignKey:user_group_id"`
}

func (User) TableName() string {
	return "users"
}

type UserTx struct {
	BaseEntity
}

type UserTest struct {
	Name string `validate:"required,min=5,max=20"` // Required field, min 5 char long max 20
	Age  int    `validate:"required,teen-person"`  // Required field, and client needs to implement our 'teen-person' tag format which we'll see later
}
