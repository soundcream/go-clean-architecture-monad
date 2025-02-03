package db

import (
	"gorm.io/gorm"
	"n4a3/clean-architecture/app/domain/entity"
)

type QueryContext[Entity entity.IBaseEntity] interface {
	Where(query interface{}, args ...interface{}) QueryContext[Entity]
	Find(query interface{}, args ...interface{}) QueryContext[Entity]
	Execute() *Entity
	Order(value interface{}) QueryContext[Entity]
}

type queryContext[Entity entity.IBaseEntity] struct {
	db *gorm.DB
}

func NewQueryContext[Entity entity.IBaseEntity](db *gorm.DB) QueryContext[Entity] {
	return &queryContext[Entity]{
		db: db,
	}
}

func (q *queryContext[Entity]) Next(db *gorm.DB) QueryContext[Entity] {
	return NewQueryContext[Entity](db)
}

func (q *queryContext[Entity]) Where(query interface{}, args ...interface{}) QueryContext[Entity] {
	return q.Next(q.db.Where(query, args))
}

func (q *queryContext[Entity]) Find(query interface{}, args ...interface{}) QueryContext[Entity] {
	return q.Next(q.db.Find(query, args))
}

func (q *queryContext[Entity]) Execute() *Entity {
	var result Entity
	q.db.Find(&result)
	return &result
}

func (q *queryContext[Entity]) Order(value interface{}) QueryContext[Entity] {
	return q.Next(q.db.Order(value))
}
