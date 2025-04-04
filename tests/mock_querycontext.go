package tests

import (
	"github.com/stretchr/testify/mock"
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/core/global"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/integrates/db"
)

type MockQueryContext[Entity entity.Entity] struct {
	mock.Mock
}

func (m *MockQueryContext[Entity]) Join(query string, args ...interface{}) db.QueryContext[Entity] {
	return m.Called(query, args).Get(0).(db.QueryContext[Entity])
}

func (m *MockQueryContext[Entity]) Preload(query string, args ...interface{}) db.QueryContext[Entity] {
	return m.Called(query, args).Get(0).(db.QueryContext[Entity])
}

func (m *MockQueryContext[Entity]) PreloadWith(field interface{}, args ...interface{}) db.QueryContext[Entity] {
	return m.Called(field, args).Get(0).(db.QueryContext[Entity])
}

func (m *MockQueryContext[Entity]) Where(query interface{}, args ...interface{}) db.QueryContext[Entity] {
	return m.Called(query, args).Get(0).(db.QueryContext[Entity])
}

func (m *MockQueryContext[Entity]) Find(query interface{}, args ...interface{}) db.QueryContext[Entity] {
	return m.Called(query, args).Get(0).(db.QueryContext[Entity])
}

func (m *MockQueryContext[Entity]) Order(value interface{}) db.QueryContext[Entity] {
	return m.Called(value).Get(0).(db.QueryContext[Entity])
}

func (m *MockQueryContext[Entity]) Group(name string) db.QueryContext[Entity] {
	return m.Called(name).Get(0).(db.QueryContext[Entity])
}

func (m *MockQueryContext[Entity]) Having(query string, args ...interface{}) db.QueryContext[Entity] {
	return m.Called(query, args).Get(0).(db.QueryContext[Entity])
}

func (m *MockQueryContext[Entity]) Fetch() core.Either[Entity, core.ErrContext] {
	return m.Called().Get(0).(core.Either[Entity, core.ErrContext])
}

func (m *MockQueryContext[Entity]) FetchAll() core.Either[[]Entity, core.ErrContext] {
	return m.Called().Get(0).(core.Either[[]Entity, core.ErrContext])
}

func (m *MockQueryContext[Entity]) ToPaging(limit int, offset int) core.Either[global.PagingModel[Entity], core.ErrContext] {
	return m.Called(limit, offset).Get(0).(core.Either[global.PagingModel[Entity], core.ErrContext])
}
