package dto

type UserDto struct {
	Name        string `json:"name" example:"John Doe"`
	Username    string `json:"username" example:"JohnDoe"`
	Email       string `json:"email" example:"john.doe@example.com"`
	Point       *int   `json:"point" example:"0"`
	UserGroupId *int   `json:"userGroupId" example:"0"`
}
