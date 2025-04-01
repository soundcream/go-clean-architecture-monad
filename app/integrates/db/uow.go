package db

import (
	"gorm.io/gorm"
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/core/global"
)

type QueryUnitOfWork interface {
	// DB chain to query
	DB() *gorm.DB
	// Query in case raw sql query map to dest
	Query(dest interface{}, sql string, values ...interface{})
}
type UnitOfWorkReader struct {
	Context
}

func (uow *UnitOfWorkReader) DB() *gorm.DB {
	return uow.Context.GetDb()
}

func (uow *UnitOfWorkReader) Query(dest interface{}, sql string, values ...interface{}) {
	uow.DB().Raw(sql, values...).Scan(dest)
}

type CommandUnitOfWork interface {
	GetDb() *gorm.DB

	BeginReadCommitTx() TransactionContext
	BeginSerializableTx() TransactionContext
	DoTransaction(func(*TransactionContext) error) error

	SavePoint(name string) core.Either[core.Unit, core.ErrContext]
	RollbackTo(name string) core.Either[core.Unit, core.ErrContext]
	Commit() core.Either[core.Unit, core.ErrContext]
	Rollback() core.Either[core.Unit, core.ErrContext]
}
type UnitOfWorkWrite struct {
	Context
}

func (uow UnitOfWorkWrite) BeginReadCommitTx() TransactionContext {
	return uow.Context.BeginReadCommitTx()
}

func (uow UnitOfWorkWrite) BeginSerializableTx() TransactionContext {
	return uow.Context.BeginSerializableTx()
}

func (uow UnitOfWorkWrite) DoTransaction(fn func(*TransactionContext) error) error {
	return uow.Context.DoTransaction(fn)
}

func NewQueryUnitOfWork(config *global.Config) core.Either[QueryUnitOfWork, error] {
	dbContext, err := NewDbContextFromConfig(config)
	if err != nil {
		return core.LeftEither[QueryUnitOfWork, error](err)
	}
	uow := &UnitOfWorkReader{
		Context: dbContext,
	}
	return core.RightEither[QueryUnitOfWork, error](uow)
}

func NewUnitOfWork(config *global.Config) core.Either[CommandUnitOfWork, error] {
	dbContext, err := NewDbContextFromConfig(config)
	if err != nil {
		return core.LeftEither[CommandUnitOfWork, error](err)
	}
	uow := UnitOfWorkWrite{
		Context: dbContext,
	}
	return core.RightEither[CommandUnitOfWork, error](uow)
}
