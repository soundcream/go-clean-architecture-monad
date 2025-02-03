package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"n4a3/clean-architecture/app/base/global"
)

type Context interface {
	Query() *gorm.DB
	BeginReadCommitTx() TransactionContext
	DoTransaction(func(*TransactionContext) error) error
}

type dbContext struct {
	db *gorm.DB
}

type TransactionContext struct {
	dbTx *gorm.DB
}

func (c dbContext) Query() *gorm.DB {
	return c.db
}

func (c dbContext) BeginReadCommitTx() TransactionContext {
	tx := c.db.Begin(&sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	return TransactionContext{
		dbTx: tx,
	}
}

func (c dbContext) DoTransaction(fn func(*TransactionContext) error) error {
	err := c.db.Transaction(func(db *gorm.DB) error {
		tx := TransactionContext{
			dbTx: db,
		}
		return fn(&tx)
	})
	return err
}

func (t *TransactionContext) SavePoint(name string) {
	t.dbTx.SavePoint(name)
}

func (t *TransactionContext) RollbackTo(name string) {
	t.dbTx.RollbackTo(name)
}

func (t *TransactionContext) Commit() {
	t.dbTx.Commit()
}

func (t *TransactionContext) Rollback() {
	t.dbTx.Rollback()
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
