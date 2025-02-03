package entity

type IBaseEntity interface {
	GetId() int
}

const PK = "primaryKey"

type BaseEntity struct {
	Id int `gorm:"primaryKey" column:"id" json:"id"`
}

func (e BaseEntity) GetId() int {
	return e.Id
}
