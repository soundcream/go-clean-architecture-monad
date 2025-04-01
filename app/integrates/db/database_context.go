package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/core/global"
)

type Context interface {
	GetDb() *gorm.DB
	BeginReadCommitTx() TransactionContext
	BeginSerializableTx() TransactionContext
	DoTransaction(func(*TransactionContext) error) error

	SavePoint(name string) core.Either[core.Unit, core.ErrContext]
	RollbackTo(name string) core.Either[core.Unit, core.ErrContext]
	Commit() core.Either[core.Unit, core.ErrContext]
	Rollback() core.Either[core.Unit, core.ErrContext]
}

type dbContext struct {
	db *gorm.DB
}

func (c *dbContext) GetDb() *gorm.DB {
	return c.db
}

func (c *dbContext) BeginReadCommitTx() TransactionContext {
	tx := c.db.Begin(&sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	return TransactionContext{
		dbTx: tx,
	}
}

func (c *dbContext) BeginSerializableTx() TransactionContext {
	tx := c.db.Begin(&sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	return TransactionContext{
		dbTx: tx,
	}
}

func (c *dbContext) SavePoint(name string) core.Either[core.Unit, core.ErrContext] {
	result := c.db.SavePoint(name)
	return core.NewEither(core.NewUnitPtr(), core.NewIfError(result.Error))
}

func (c *dbContext) RollbackTo(name string) core.Either[core.Unit, core.ErrContext] {
	result := c.db.RollbackTo(name)
	return core.NewEither(core.NewUnitPtr(), core.NewIfError(result.Error))
}

func (c *dbContext) Commit() core.Either[core.Unit, core.ErrContext] {
	result := c.db.Commit()
	return core.NewEither(core.NewUnitPtr(), core.NewIfError(result.Error))
}

func (c *dbContext) Rollback() core.Either[core.Unit, core.ErrContext] {
	result := c.db.Rollback()
	return core.NewEither(core.NewUnitPtr(), core.NewIfError(result.Error))
}

// DoTransaction Example use
// c := *cmdUoW.Right
//
//	err := c.DoTransaction(func(context *db.TransactionContext) error {
//		return nil
//	})
func (c *dbContext) DoTransaction(fn func(*TransactionContext) error) error {
	err := c.db.Transaction(func(db *gorm.DB) error {
		tx := TransactionContext{
			dbTx: db,
		}
		er := fn(&tx)
		if er != nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
		return er
	})
	return err
}

func getDSN(config *global.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DbConfig.Host,
		config.DbConfig.Username,
		config.DbConfig.Password,
		config.DbConfig.DbName,
		config.DbConfig.Port)
}

func NewDbContextFromConfig(config *global.Config) (Context, error) {
	db, err := gorm.Open(postgres.Open(getDSN(config)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}
	context := &dbContext{
		db: db,
	}
	return context, nil
}

func NewDbContext(db *gorm.DB) (Context, error) {
	return &dbContext{
		db: db,
	}, nil
}
