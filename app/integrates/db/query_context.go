package db

import (
	"gorm.io/gorm"
	"math"
	_ "math"
	"n4a3/clean-architecture/app/base/global"
	"n4a3/clean-architecture/app/base/util"
	stringutil "n4a3/clean-architecture/app/base/util/string"
	"n4a3/clean-architecture/app/domain/entity"
)

type QueryContext[Entity entity.Entity] interface {
	// Join LeftJoin ex: ("query = ?", "value")
	Join(query string, args ...interface{}) QueryContext[Entity]
	// Preload ex: ("query = ?", "value")
	Preload(query string, args ...interface{}) QueryContext[Entity]
	// PreloadWith ex: (User{}.Name) or (User{}.Name, "name = ?", "A")
	PreloadWith(field interface{}, args ...interface{}) QueryContext[Entity]
	// Where ex: ("query = 2")
	Where(query interface{}, args ...interface{}) QueryContext[Entity]
	// Find ex: ("query = ?", "value")
	Find(query interface{}, args ...interface{}) QueryContext[Entity]
	// Order ex: ("name") or ("name desc")
	Order(value interface{}) QueryContext[Entity]
	// Group ex: ("name")
	Group(name string) QueryContext[Entity]
	// Having ex:
	Having(query string, args ...interface{}) QueryContext[Entity]
	// Fetch execute data as *Entity
	Fetch() *Entity
	// FetchAll execute data as []Entity
	FetchAll() []Entity
	// ToPaging use for BuildQueryPaging
	ToPaging(limit int, offset int) global.PagingModel[Entity]
}

type queryContext[Entity entity.Entity] struct {
	db *gorm.DB
}

func NewQueryContext[Entity entity.Entity](db *gorm.DB) QueryContext[Entity] {
	return &queryContext[Entity]{
		db: db,
	}
}

func (q *queryContext[Entity]) Where(query interface{}, args ...interface{}) QueryContext[Entity] {
	return next[Entity](q.db.Where(query, args...))
}

func (q *queryContext[Entity]) Find(query interface{}, args ...interface{}) QueryContext[Entity] {
	return next[Entity](q.db.Find(query, args...))
}

func (q *queryContext[Entity]) Fetch() *Entity {
	var result Entity
	q.db.Find(&result)
	return &result
}

func (q *queryContext[Entity]) FetchAll() []Entity {
	var result []Entity
	q.db.Find(&result)
	return result
}

func (q *queryContext[Entity]) ToPaging(limit int, offset int) global.PagingModel[Entity] {
	var result []Entity
	var c int64
	var total = 0
	q.db.Count(&c)
	q.db.Limit(limit).Offset(offset).Find(&result)
	total = int(c)
	//totalPages := int(math.Ceil((float64(total) * 1.0) / float64(limit)))
	page := int(math.Ceil((float64(offset)*1.0)/float64(limit)) + 1)
	return global.PagingModel[Entity]{
		Data:   result,
		Total:  total,
		Limit:  limit,
		Offset: offset,
		Page:   page,
	}
}

func (q *queryContext[Entity]) Order(value interface{}) QueryContext[Entity] {
	return next[Entity](q.db.Order(value))
}

func (q *queryContext[Entity]) Preload(query string, args ...interface{}) QueryContext[Entity] {
	return next[Entity](q.db.Preload(query, args...))
}

func (q *queryContext[Entity]) PreloadWith(field interface{}, args ...interface{}) QueryContext[Entity] {
	fieldName := util.GetFieldName(field)
	if !stringutil.IsNullOrEmpty(fieldName) {
		return next[Entity](q.db.Preload(fieldName, args...))
	}
	return next[Entity](q.db)
}

func (q *queryContext[Entity]) Join(query string, args ...interface{}) QueryContext[Entity] {
	return next[Entity](q.db.Joins(query, args...))
}

func (q *queryContext[Entity]) Group(name string) QueryContext[Entity] {
	return next[Entity](q.db.Group(name))
}

func (q *queryContext[Entity]) Having(query string, args ...interface{}) QueryContext[Entity] {
	return next[Entity](q.db.Having(query, args...))
}

func next[Entity entity.Entity](db *gorm.DB) QueryContext[Entity] {
	return NewQueryContext[Entity](db)
}
