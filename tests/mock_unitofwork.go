package tests

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/integrates/db"
)

type MockQueryUnitOfWork struct {
	mock.Mock
}

type MockCommandUnitOfWork struct {
	mock.Mock
}

func (m *MockQueryUnitOfWork) DB() *gorm.DB {
	return m.Called().Get(0).(*gorm.DB)
}

func (m *MockQueryUnitOfWork) Query(dest interface{}, sql string, values ...interface{}) (tx *gorm.DB) {
	return m.Called(dest, sql, values).Get(0).(*gorm.DB)
}

func (m *MockCommandUnitOfWork) GetDb() *gorm.DB {
	return m.Called().Get(0).(*gorm.DB)
}
func (m *MockCommandUnitOfWork) BeginReadCommitTx() db.TransactionContext {
	return m.Called().Get(0).(db.TransactionContext)
}
func (m *MockCommandUnitOfWork) BeginSerializableTx() db.TransactionContext {
	return m.Called().Get(0).(db.TransactionContext)
}
func (m *MockCommandUnitOfWork) DoTransaction(func(*db.TransactionContext) error) error {
	return m.Called(nil).Error(0)
}
func (m *MockCommandUnitOfWork) SavePoint(name string) core.Either[core.Unit, core.ErrContext] {
	return m.Called(name).Get(0).(core.Either[core.Unit, core.ErrContext])
}
func (m *MockCommandUnitOfWork) RollbackTo(name string) core.Either[core.Unit, core.ErrContext] {
	return m.Called(name).Get(0).(core.Either[core.Unit, core.ErrContext])
}
func (m *MockCommandUnitOfWork) Commit() core.Either[core.Unit, core.ErrContext] {
	return m.Called().Get(0).(core.Either[core.Unit, core.ErrContext])
}
func (m *MockCommandUnitOfWork) Rollback() core.Either[core.Unit, core.ErrContext] {
	return m.Called().Get(0).(core.Either[core.Unit, core.ErrContext])
}
