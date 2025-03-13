package db

import (
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/domain/entity"
)

type Repository[Entity entity.Entity] interface {
	// Insert insert entity
	Insert(entity *Entity) base.Either[int64, base.ErrContext]
	// BulkInsert insert entities
	BulkInsert(entities *[]Entity) base.Either[int64, base.ErrContext]
	// Delete delete entity
	Delete(entity Entity) base.Either[int64, base.ErrContext]
	// DeleteById delete entity by id
	DeleteById(id int) base.Either[int64, base.ErrContext]
	// Update update all entity
	UpdateAll(entity *Entity) base.Either[int64, base.ErrContext]
	// Updates Update attributes with `struct`, will only update non-zero fields
	Update(id int, entity Entity) base.Either[int64, base.ErrContext]
	// UpdateWhere update 1 field where condition
	UpdateWhere(column string, value interface{}, query interface{}, args ...interface{}) base.Either[int64, base.ErrContext]
}

type repository[Entity entity.Entity] struct {
	uow       CommandUnitOfWork
	tableName string
}

func NewRepository[Entity entity.Entity](uow *CommandUnitOfWork) Repository[Entity] {
	var e = *new(Entity)
	return &repository[Entity]{
		uow:       *uow,
		tableName: e.TableName(),
	}
}

func (repo *repository[Entity]) Insert(entity *Entity) base.Either[int64, base.ErrContext] {

	result := repo.uow.GetDb().Create(entity)
	return base.NewEither(&result.RowsAffected, base.NewIfError(result.Error))
}

func (repo *repository[Entity]) BulkInsert(entities *[]Entity) base.Either[int64, base.ErrContext] {
	result := repo.uow.GetDb().Create(entities)
	return base.NewEither(&result.RowsAffected, base.NewIfError(result.Error))
}

func (repo *repository[Entity]) Delete(entity Entity) base.Either[int64, base.ErrContext] {
	result := repo.uow.GetDb().Delete(&entity)
	return base.NewEither(&result.RowsAffected, base.NewIfError(result.Error))
}

func (repo *repository[Entity]) DeleteById(id int) base.Either[int64, base.ErrContext] {
	var e Entity
	result := repo.uow.GetDb().Where("id = ?", id).Delete(&e)
	return base.NewEither(&result.RowsAffected, base.NewIfError(result.Error))
}

func (repo *repository[Entity]) UpdateAll(entity *Entity) base.Either[int64, base.ErrContext] {
	result := repo.uow.GetDb().Save(entity)
	return base.NewEither(&result.RowsAffected, base.NewIfError(result.Error))
}

func (repo *repository[Entity]) Update(id int, entity Entity) base.Either[int64, base.ErrContext] {
	var e = new(Entity)

	result := repo.uow.GetDb().Model(&e).Where("id = ?", id).Updates(entity)
	return base.NewEither(&result.RowsAffected, base.NewIfError(result.Error))
}

func (repo *repository[Entity]) UpdateWhere(column string, value interface{}, query interface{}, args ...interface{}) base.Either[int64, base.ErrContext] {
	var e Entity
	result := repo.uow.GetDb().Model(&e).Where(query, args...).Update(column, value)
	return base.NewEither(&result.RowsAffected, base.NewIfError(result.Error))
}
