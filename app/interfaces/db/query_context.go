package db

import "gorm.io/gorm"

type QueryContext interface {
	Where(query interface{}, args ...interface{}) QueryContext
	Find(query interface{}, args ...interface{}) QueryContext
	Order(value interface{}) QueryContext
	Query() *gorm.DB
}

type queryContext struct {
	db *gorm.DB
}

func NewQueryContext(db *gorm.DB) QueryContext {
	return &queryContext{
		db: db,
	}
}

func (q *queryContext) Query() *gorm.DB {
	return q.db
}

func (q *queryContext) Where(query interface{}, args ...interface{}) QueryContext {
	next := q.db.Where(query, args)
	return NewQueryContext(next)
}

func (q *queryContext) Find(query interface{}, args ...interface{}) QueryContext {
	next := q.db.Find(query, args)
	return NewQueryContext(next)
}

func (q *queryContext) Order(value interface{}) QueryContext {
	next := q.db.Order(value)
	return NewQueryContext(next)
}
