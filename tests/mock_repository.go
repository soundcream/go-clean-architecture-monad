package tests

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"math/big"
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/integrates/db"
)

type MockRepository[Entity entity.Entity] struct {
	mock.Mock
}

func (m *MockRepository[Entity]) Query() db.QueryContext[Entity] {
	arguments := m.Called()
	return arguments.Get(0).(db.QueryContext[Entity])
}

func (m *MockRepository[Entity]) BuildQueryPagination() db.QueryContext[Entity] {
	arguments := m.Called()
	return arguments.Get(0).(db.QueryContext[Entity])
}

func (m *MockRepository[Entity]) FindByIdPreload(id int, preloads *map[string][]interface{}) *Entity {
	arguments := m.Called(id, preloads)
	return arguments.Get(0).(*Entity)
}

func (m *MockRepository[Entity]) FindByIdPreloadInclude(id int, field interface{}, args ...interface{}) *Entity {
	arguments := m.Called(id, field, args)
	return arguments.Get(0).(*Entity)
}

func (m *MockRepository[Entity]) FindByIdLeftJoins(id int, joins ...string) *Entity {
	arguments := m.Called(id, joins)
	return arguments.Get(0).(*Entity)
}

func (m *MockRepository[Entity]) FindByIdInnerJoins(id int, joins ...string) *Entity {
	arguments := m.Called(id, joins)
	return arguments.Get(0).(*Entity)
}

func (m *MockRepository[Entity]) FindById(id int) *Entity {
	arguments := m.Mock.Called(id)
	return arguments.Get(0).(*Entity)
}

func (m *MockRepository[Entity]) FindBy(query interface{}, args ...interface{}) *Entity {
	arguments := m.Called(query, args)
	return arguments.Get(0).(*Entity)
}

func (m *MockRepository[Entity]) FindOrderBy(order interface{}, query interface{}, args ...interface{}) *Entity {
	arguments := m.Called(order, query, args)
	return arguments.Get(0).(*Entity)
}

func (m *MockRepository[Entity]) Where(query interface{}, args ...interface{}) []Entity {
	arguments := m.Called(query, args)
	return arguments.Get(0).([]Entity)
}

func (m *MockRepository[Entity]) WhereOrderBy(order interface{}, query interface{}, args ...interface{}) []Entity {
	arguments := m.Called(order, query, args)
	return arguments.Get(0).([]Entity)
}

func (m *MockRepository[Entity]) Count(query interface{}, args ...interface{}) *int {
	arguments := m.Called(query, args)
	return arguments.Get(0).(*int)
}

func (m *MockRepository[Entity]) CountBig(query interface{}, args ...interface{}) *int64 {
	arguments := m.Called(query, args)
	return arguments.Get(0).(*int64)
}

func (m *MockRepository[Entity]) Sum(query interface{}, args ...interface{}) *int {
	arguments := m.Called(query, args)
	return arguments.Get(0).(*int)
}

func (m *MockRepository[Entity]) SumBig(query interface{}, args ...interface{}) *big.Float {
	arguments := m.Called(query, args)
	return arguments.Get(0).(*big.Float)
}

func (m *MockRepository[Entity]) init() *gorm.DB {
	arguments := m.Called()
	return arguments.Get(0).(*gorm.DB)
}

func (m *MockRepository[Entity]) Insert(entity *Entity) base.Either[int64, base.ErrContext] {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository[Entity]) BulkInsert(entities *[]Entity) base.Either[int64, base.ErrContext] {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository[Entity]) Delete(entity Entity) base.Either[int64, base.ErrContext] {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository[Entity]) DeleteById(id int) base.Either[int64, base.ErrContext] {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository[Entity]) Update(entity *Entity) base.Either[int64, base.ErrContext] {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository[Entity]) Updates(entity Entity) base.Either[int64, base.ErrContext] {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository[Entity]) UpdateWhere(column string, value interface{}, query interface{}, args ...interface{}) base.Either[int64, base.ErrContext] {
	//TODO implement me
	panic("implement me")
}
