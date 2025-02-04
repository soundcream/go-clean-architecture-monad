package entity

import (
	"time"
)

type UserGroup struct {
	BaseEntity  `table-name:"customer_groups"`
	Name        string    `json:"name" example:"GOLD"`
	Level       int       `json:"level" example:"1"`
	IsActive    bool      `json:"isActive" example:"false"`
	CreatedDate time.Time `json:"createdDate" example:"2020-01-01"`
	Users       []User    `json:"users" gorm:"foreignKey:user_group_id"`
}
