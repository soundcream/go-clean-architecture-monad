package entity

import (
	"time"
)

type UserGroup struct {
	BaseEntity
	Activation
	Name  string `column:"name" json:"name" example:"GOLD"`
	Level int    `column:"level" json:"level" example:"1"`
	Users []User `json:"users" gorm:"foreignKey:user_group_id"`
}

func (UserGroup) TableName() string {
	return "user_groups"
}

type UserGroup2 struct {
	BaseEntity
	Name        string    `column:"name" json:"name" example:"GOLD"`
	Level       int       `column:"level" json:"level" example:"1"`
	IsActive    bool      `column:"is_active" json:"isActive" example:"false"`
	CreatedDate time.Time `column:"created_date" json:"createdDate" example:"2020-01-01"`
	Users       []User    `json:"users" gorm:"foreignKey:user_group_id"`
}

func (UserGroup2) TableName() string {
	return "user_groups"
}
