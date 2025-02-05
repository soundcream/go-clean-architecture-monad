package db

import "gorm.io/gorm"

type TransactionContext struct {
	dbTx *gorm.DB
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
