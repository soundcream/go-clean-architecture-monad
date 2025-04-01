package tests

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"math/big"
	"n4a3/clean-architecture/app/core"
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

func (m *MockRepository[Entity]) FindByIdPreload(id int, preloads *map[string][]interface{}) core.Either[Entity, core.ErrContext] {
	arguments := m.Called(id, preloads)
	result := arguments.Get(0).(*Entity)
	return core.RightEither[Entity, core.ErrContext](*result)
}

func (m *MockRepository[Entity]) FindByIdPreloadInclude(id int, field interface{}, args ...interface{}) core.Either[Entity, core.ErrContext] {
	arguments := m.Called(id, field, args)
	result := arguments.Get(0).(*Entity)
	return core.RightEither[Entity, core.ErrContext](*result)
}

func (m *MockRepository[Entity]) FindByIdLeftJoins(id int, joins ...string) core.Either[Entity, core.ErrContext] {
	arguments := m.Called(id, joins)
	result := arguments.Get(0).(*Entity)
	return core.RightEither[Entity, core.ErrContext](*result)
}

func (m *MockRepository[Entity]) FindByIdInnerJoins(id int, joins ...string) core.Either[Entity, core.ErrContext] {
	arguments := m.Called(id, joins)
	result := arguments.Get(0).(*Entity)
	return core.RightEither[Entity, core.ErrContext](*result)
}

func (m *MockRepository[Entity]) FindById(id int) core.Either[Entity, core.ErrContext] {
	arguments := m.Mock.Called(id)
	result := arguments.Get(0).(*Entity)
	return core.RightEither[Entity, core.ErrContext](*result)
}

func (m *MockRepository[Entity]) FindBy(query interface{}, args ...interface{}) core.Either[Entity, core.ErrContext] {
	arguments := m.Called(query, args)
	result := arguments.Get(0).(*Entity)
	return core.RightEither[Entity, core.ErrContext](*result)
}

func (m *MockRepository[Entity]) FindOrderBy(order interface{}, query interface{}, args ...interface{}) core.Either[Entity, core.ErrContext] {
	arguments := m.Called(order, query, args)
	result := arguments.Get(0).(*Entity)
	return core.RightEither[Entity, core.ErrContext](*result)
}

func (m *MockRepository[Entity]) Where(query interface{}, args ...interface{}) core.Either[[]Entity, core.ErrContext] {
	arguments := m.Called(query, args)
	result := arguments.Get(0).([]Entity)
	return core.RightEither[[]Entity, core.ErrContext](result)
}

func (m *MockRepository[Entity]) WhereOrderBy(order interface{}, query interface{}, args ...interface{}) core.Either[[]Entity, core.ErrContext] {
	arguments := m.Called(order, query, args)
	result := arguments.Get(0).([]Entity)
	return core.RightEither[[]Entity, core.ErrContext](result)
}

func (m *MockRepository[Entity]) Count(query interface{}, args ...interface{}) core.Either[int, core.ErrContext] {
	arguments := m.Called(query, args)
	result := arguments.Get(0).(*int)
	return core.RightEither[int, core.ErrContext](*result)
}

func (m *MockRepository[Entity]) CountBig(query interface{}, args ...interface{}) core.Either[int64, core.ErrContext] {
	arguments := m.Called(query, args)
	result := arguments.Get(0).(*int64)
	return core.RightEither[int64, core.ErrContext](*result)
}

func (m *MockRepository[Entity]) Sum(query interface{}, args ...interface{}) core.Either[int, core.ErrContext] {
	arguments := m.Called(query, args)
	result := arguments.Get(0).(*int)
	return core.RightEither[int, core.ErrContext](*result)
}

func (m *MockRepository[Entity]) SumBig(query interface{}, args ...interface{}) core.Either[big.Float, core.ErrContext] {
	arguments := m.Called(query, args)
	result := arguments.Get(0).(*big.Float)
	return core.RightEither[big.Float, core.ErrContext](*result)
}

func (m *MockRepository[Entity]) init() *gorm.DB {
	arguments := m.Called()
	return arguments.Get(0).(*gorm.DB)
}

func (m *MockRepository[Entity]) Insert(entity *Entity) core.Either[int64, core.ErrContext] {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository[Entity]) BulkInsert(entities *[]Entity) core.Either[int64, core.ErrContext] {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository[Entity]) Delete(entity Entity) core.Either[int64, core.ErrContext] {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository[Entity]) DeleteById(id int) core.Either[int64, core.ErrContext] {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository[Entity]) Update(entity *Entity) core.Either[int64, core.ErrContext] {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository[Entity]) Updates(entity Entity) core.Either[int64, core.ErrContext] {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository[Entity]) UpdateWhere(column string, value interface{}, query interface{}, args ...interface{}) core.Either[int64, core.ErrContext] {
	//TODO implement me
	panic("implement me")
}
