package db

import (
	"gorm.io/gorm"
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/global"
)

type QueryUnitOfWork interface {
	DB() *gorm.DB
	Query(dest interface{}, sql string, values ...interface{})
}
type UnitOfWorkReader struct {
	DbContext Context
	//Db        *gorm.DB
}

func (uow *UnitOfWorkReader) DB() *gorm.DB {
	return uow.DbContext.Query()
}

func (uow *UnitOfWorkReader) Query(dest interface{}, sql string, values ...interface{}) {
	uow.DB().Raw(sql, values...).Scan(dest)
}

type CommandUnitOfWork interface {
	BeginReadCommitTx() TransactionContext
	DoTransaction(func(*TransactionContext) error) error
	//Transaction(func(*Context) error) error
}
type UnitOfWorkWrite struct {
	DbContext Context
	//Db        *gorm.DB
}

func (uow UnitOfWorkWrite) BeginReadCommitTx() TransactionContext {
	return uow.BeginReadCommitTx()
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
		DbContext: dbContext,
	}
	return base.RightEither[QueryUnitOfWork, error](uow)
}

func NewUnitOfWork(config *global.Config) base.Either[CommandUnitOfWork, error] {
	dbContext, err := NewDbContextFromConfig(config)
	if err != nil {
		return base.LeftEither[CommandUnitOfWork, error](err)
	}
	uow := UnitOfWorkWrite{
		DbContext: dbContext,
	}
	return base.RightEither[CommandUnitOfWork, error](uow)
}

//
//func ConnectDb() {
//	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Error),
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var tx = db.Begin(&sql.TxOptions{
//		Isolation: sql.LevelReadCommitted,
//		ReadOnly:  false,
//	})
//	tx.SavePoint("SaveUser")
//	tx.RollbackTo("SaveUser")
//	tx.Commit()
//	tx.Rollback()
//
//	// transaction
//
//	// auto transaction
//	_ = db.Transaction(func(tx *gorm.DB) error {
//		tx.SavePoint("SaveUser")
//		tx.RollbackTo("SaveUser")
//		return nil
//	}, &sql.TxOptions{
//		Isolation: sql.LevelReadCommitted,
//	})
//}
