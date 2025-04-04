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
	return m.Called().Get(0).(db.QueryContext[Entity])
}

func (m *MockRepository[Entity]) BuildQueryPagination() db.QueryContext[Entity] {
	return m.Called().Get(0).(db.QueryContext[Entity])
}

func (m *MockRepository[Entity]) FindByIdPreload(id int, preloads *map[string][]interface{}) core.Either[Entity, core.ErrContext] {
	return m.Called(id, preloads).Get(0).(core.Either[Entity, core.ErrContext])
}

func (m *MockRepository[Entity]) FindByIdPreloadInclude(id int, field interface{}, args ...interface{}) core.Either[Entity, core.ErrContext] {
	return m.Called(id, field, args).Get(0).(core.Either[Entity, core.ErrContext])
}

func (m *MockRepository[Entity]) FindByIdLeftJoins(id int, joins ...string) core.Either[Entity, core.ErrContext] {
	return m.Called(id, joins).Get(0).(core.Either[Entity, core.ErrContext])
}

func (m *MockRepository[Entity]) FindByIdInnerJoins(id int, joins ...string) core.Either[Entity, core.ErrContext] {
	return m.Called(id, joins).Get(0).(core.Either[Entity, core.ErrContext])
}

func (m *MockRepository[Entity]) FindById(id int) core.Either[Entity, core.ErrContext] {
	return m.Mock.Called(id).Get(0).(core.Either[Entity, core.ErrContext])
}

func (m *MockRepository[Entity]) FindBy(query interface{}, args ...interface{}) core.Either[Entity, core.ErrContext] {
	return m.Called(query, args).Get(0).(core.Either[Entity, core.ErrContext])
}

func (m *MockRepository[Entity]) FindOrderBy(order interface{}, query interface{}, args ...interface{}) core.Either[Entity, core.ErrContext] {
	return m.Called(order, query, args).Get(0).(core.Either[Entity, core.ErrContext])
}

func (m *MockRepository[Entity]) Where(query interface{}, args ...interface{}) core.Either[[]Entity, core.ErrContext] {
	return m.Called(query, args).Get(0).(core.Either[[]Entity, core.ErrContext])
}

func (m *MockRepository[Entity]) WhereOrderBy(order interface{}, query interface{}, args ...interface{}) core.Either[[]Entity, core.ErrContext] {
	return m.Called(order, query, args).Get(0).(core.Either[[]Entity, core.ErrContext])
}

func (m *MockRepository[Entity]) Count(query interface{}, args ...interface{}) core.Either[int, core.ErrContext] {
	return m.Called(query, args).Get(0).(core.Either[int, core.ErrContext])
}

func (m *MockRepository[Entity]) CountBig(query interface{}, args ...interface{}) core.Either[int64, core.ErrContext] {
	return m.Called(query, args).Get(0).(core.Either[int64, core.ErrContext])
}

func (m *MockRepository[Entity]) Sum(query interface{}, args ...interface{}) core.Either[int, core.ErrContext] {
	return m.Called(query, args).Get(0).(core.Either[int, core.ErrContext])
}

func (m *MockRepository[Entity]) SumBig(query interface{}, args ...interface{}) core.Either[big.Float, core.ErrContext] {
	return m.Called(query, args).Get(0).(core.Either[big.Float, core.ErrContext])
}

func (m *MockRepository[Entity]) init() *gorm.DB {
	return m.Called().Get(0).(*gorm.DB)
}

func (m *MockRepository[Entity]) Insert(entity *Entity) core.Either[int64, core.ErrContext] {
	return m.Mock.Called(entity).Get(0).(core.Either[int64, core.ErrContext])
}

func (m *MockRepository[Entity]) BulkInsert(entities *[]Entity) core.Either[int64, core.ErrContext] {
	return m.Mock.Called(entities).Get(0).(core.Either[int64, core.ErrContext])
}

func (m *MockRepository[Entity]) Delete(entity Entity) core.Either[int64, core.ErrContext] {
	return m.Mock.Called(entity).Get(0).(core.Either[int64, core.ErrContext])
}

func (m *MockRepository[Entity]) DeleteById(id int) core.Either[int64, core.ErrContext] {
	return m.Mock.Called(id).Get(0).(core.Either[int64, core.ErrContext])
}

func (m *MockRepository[Entity]) UpdateAllFields(entity *Entity) core.Either[int64, core.ErrContext] {
	return m.Mock.Called(entity).Get(0).(core.Either[int64, core.ErrContext])
}

func (m *MockRepository[Entity]) Update(id int, entity Entity) core.Either[int64, core.ErrContext] {
	return m.Mock.Called(id, entity).Get(0).(core.Either[int64, core.ErrContext])
}

func (m *MockRepository[Entity]) UpdateWhere(column string, value interface{}, query interface{}, args ...interface{}) core.Either[int64, core.ErrContext] {
	return m.Mock.Called(column, value, query, args).Get(0).(core.Either[int64, core.ErrContext])
}

func (m *MockRepository[Entity]) UpdatesWhere(entity Entity, query interface{}, args ...interface{}) core.Either[int64, core.ErrContext] {
	return m.Mock.Called(entity, query, args).Get(0).(core.Either[int64, core.ErrContext])
}

func (m *MockRepository[Entity]) UpdatesFieldsWhere(fields map[string]interface{}, query interface{}, args ...interface{}) core.Either[int64, core.ErrContext] {
	return m.Mock.Called(fields, query, args).Get(0).(core.Either[int64, core.ErrContext])
}
