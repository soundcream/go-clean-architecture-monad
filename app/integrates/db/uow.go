package db

import (
	"gorm.io/gorm"
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/global"
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

	SavePoint(name string) base.Either[base.Unit, base.ErrContext]
	RollbackTo(name string) base.Either[base.Unit, base.ErrContext]
	Commit() base.Either[base.Unit, base.ErrContext]
	Rollback() base.Either[base.Unit, base.ErrContext]
}
type UnitOfWorkWrite struct {
	Context
}

func (uow UnitOfWorkWrite) BeginReadCommitTx() TransactionContext {
	return uow.BeginReadCommitTx()
}

func (uow UnitOfWorkWrite) BeginSerializableTx() TransactionContext {
	return uow.BeginSerializableTx()
}

func (uow UnitOfWorkWrite) DoTransaction(fn func(*TransactionContext) error) error {
	return uow.DoTransaction(fn)
}

func NewQueryUnitOfWork(config *global.Config) base.Either[QueryUnitOfWork, error] {
	dbContext, err := NewDbContextFromConfig(config)
	if err != nil {
		return base.LeftEither[QueryUnitOfWork, error](err)
	}
	uow := &UnitOfWorkReader{
		Context: dbContext,
	}
	return base.RightEither[QueryUnitOfWork, error](uow)
}

func NewUnitOfWork(config *global.Config) base.Either[CommandUnitOfWork, error] {
	dbContext, err := NewDbContextFromConfig(config)
	if err != nil {
		return base.LeftEither[CommandUnitOfWork, error](err)
	}
	uow := UnitOfWorkWrite{
		Context: dbContext,
	}
	return base.RightEither[CommandUnitOfWork, error](uow)
}
