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

func (m *MockRepository[Entity]) FindByIdPreload(id int, preloads *map[string][]interface{}) base.Either[Entity, base.ErrContext] {
	arguments := m.Called(id, preloads)
	result := arguments.Get(0).(*Entity)
	return base.RightEither[Entity, base.ErrContext](*result)
}

func (m *MockRepository[Entity]) FindByIdPreloadInclude(id int, field interface{}, args ...interface{}) base.Either[Entity, base.ErrContext] {
	arguments := m.Called(id, field, args)
	result := arguments.Get(0).(*Entity)
	return base.RightEither[Entity, base.ErrContext](*result)
}

func (m *MockRepository[Entity]) FindByIdLeftJoins(id int, joins ...string) base.Either[Entity, base.ErrContext] {
	arguments := m.Called(id, joins)
	result := arguments.Get(0).(*Entity)
	return base.RightEither[Entity, base.ErrContext](*result)
}

func (m *MockRepository[Entity]) FindByIdInnerJoins(id int, joins ...string) base.Either[Entity, base.ErrContext] {
	arguments := m.Called(id, joins)
	result := arguments.Get(0).(*Entity)
	return base.RightEither[Entity, base.ErrContext](*result)
}

func (m *MockRepository[Entity]) FindById(id int) base.Either[Entity, base.ErrContext] {
	arguments := m.Mock.Called(id)
	result := arguments.Get(0).(*Entity)
	return base.RightEither[Entity, base.ErrContext](*result)
}

func (m *MockRepository[Entity]) FindBy(query interface{}, args ...interface{}) base.Either[Entity, base.ErrContext] {
	arguments := m.Called(query, args)
	result := arguments.Get(0).(*Entity)
	return base.RightEither[Entity, base.ErrContext](*result)
}

func (m *MockRepository[Entity]) FindOrderBy(order interface{}, query interface{}, args ...interface{}) base.Either[Entity, base.ErrContext] {
	arguments := m.Called(order, query, args)
	result := arguments.Get(0).(*Entity)
	return base.RightEither[Entity, base.ErrContext](*result)
}

func (m *MockRepository[Entity]) Where(query interface{}, args ...interface{}) base.Either[[]Entity, base.ErrContext] {
	arguments := m.Called(query, args)
	result := arguments.Get(0).([]Entity)
	return base.RightEither[[]Entity, base.ErrContext](result)
}

func (m *MockRepository[Entity]) WhereOrderBy(order interface{}, query interface{}, args ...interface{}) base.Either[[]Entity, base.ErrContext] {
	arguments := m.Called(order, query, args)
	result := arguments.Get(0).([]Entity)
	return base.RightEither[[]Entity, base.ErrContext](result)
}

func (m *MockRepository[Entity]) Count(query interface{}, args ...interface{}) base.Either[int, base.ErrContext] {
	arguments := m.Called(query, args)
	result := arguments.Get(0).(*int)
	return base.RightEither[int, base.ErrContext](*result)
}

func (m *MockRepository[Entity]) CountBig(query interface{}, args ...interface{}) base.Either[int64, base.ErrContext] {
	arguments := m.Called(query, args)
	result := arguments.Get(0).(*int64)
	return base.RightEither[int64, base.ErrContext](*result)
}

func (m *MockRepository[Entity]) Sum(query interface{}, args ...interface{}) base.Either[int, base.ErrContext] {
	arguments := m.Called(query, args)
	result := arguments.Get(0).(*int)
	return base.RightEither[int, base.ErrContext](*result)
}

func (m *MockRepository[Entity]) SumBig(query interface{}, args ...interface{}) base.Either[big.Float, base.ErrContext] {
	arguments := m.Called(query, args)
	result := arguments.Get(0).(*big.Float)
	return base.RightEither[big.Float, base.ErrContext](*result)
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
