package db

import (
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/domain/entity"
)

type Repository[Entity entity.Entity] interface {
	// Insert insert entity
	Insert(entity *Entity) core.Either[int64, core.ErrContext]
	// BulkInsert insert entities
	BulkInsert(entities *[]Entity) core.Either[int64, core.ErrContext]
	// Delete delete entity
	Delete(entity Entity) core.Either[int64, core.ErrContext]
	// DeleteById delete entity by id
	DeleteById(id int) core.Either[int64, core.ErrContext]
	// UpdateAllFields update all entity
	UpdateAllFields(entity *Entity) core.Either[int64, core.ErrContext]
	// Update attributes with `struct`, will only update non-zero fields
	Update(id int, entity Entity) core.Either[int64, core.ErrContext]
	// UpdateWhere update 1 field where condition
	// EX UpdateWhere("name", "My Name", "id = ?", 12)
	UpdateWhere(column string, value interface{}, query interface{}, args ...interface{}) core.Either[int64, core.ErrContext]
	// UpdatesWhere updates where condition
	// EX UpdatesWhere(Entity{}, "id = ?", 12)
	UpdatesWhere(entity Entity, query interface{}, args ...interface{}) core.Either[int64, core.ErrContext]
	// UpdatesFieldsWhere
	// EX UpdatesFieldsWhere(map[string]interface{}{"name": "hello", "age": 18, "active": false}, "id = ?", 12)
	UpdatesFieldsWhere(fields map[string]interface{}, query interface{}, args ...interface{}) core.Either[int64, core.ErrContext]
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

func (repo *repository[Entity]) Insert(entity *Entity) core.Either[int64, core.ErrContext] {

	result := repo.uow.GetDb().Create(entity)
	return core.NewEither(&result.RowsAffected, core.NewIfError(result.Error))
}

func (repo *repository[Entity]) BulkInsert(entities *[]Entity) core.Either[int64, core.ErrContext] {
	result := repo.uow.GetDb().Create(entities)
	return core.NewEither(&result.RowsAffected, core.NewIfError(result.Error))
}

func (repo *repository[Entity]) Delete(entity Entity) core.Either[int64, core.ErrContext] {
	result := repo.uow.GetDb().Delete(&entity)
	return core.NewEither(&result.RowsAffected, core.NewIfError(result.Error))
}

func (repo *repository[Entity]) DeleteById(id int) core.Either[int64, core.ErrContext] {
	var e Entity
	result := repo.uow.GetDb().Where("id = ?", id).Delete(&e)
	return core.NewEither(&result.RowsAffected, core.NewIfError(result.Error))
}

func (repo *repository[Entity]) UpdateAllFields(entity *Entity) core.Either[int64, core.ErrContext] {
	result := repo.uow.GetDb().Save(entity)
	return core.NewEither(&result.RowsAffected, core.NewIfError(result.Error))
}

func (repo *repository[Entity]) Update(id int, entity Entity) core.Either[int64, core.ErrContext] {
	var e = new(Entity)
	result := repo.uow.GetDb().Model(&e).Where("id = ?", id).Updates(entity)
	return core.NewEither(&result.RowsAffected, core.NewIfError(result.Error))
}

func (repo *repository[Entity]) UpdateWhere(column string, value interface{}, query interface{}, args ...interface{}) core.Either[int64, core.ErrContext] {
	var e Entity
	result := repo.uow.GetDb().Model(&e).Where(query, args...).Update(column, value)
	return core.NewEither(&result.RowsAffected, core.NewIfError(result.Error))
}

func (repo *repository[Entity]) UpdatesWhere(entity Entity, query interface{}, args ...interface{}) core.Either[int64, core.ErrContext] {
	var e Entity
	result := repo.uow.GetDb().Model(&e).Where(query, args...).Updates(entity)
	return core.NewEither(&result.RowsAffected, core.NewIfError(result.Error))
}

func (repo *repository[Entity]) UpdatesFieldsWhere(fields map[string]interface{}, query interface{}, args ...interface{}) core.Either[int64, core.ErrContext] {
	var e Entity
	result := repo.uow.GetDb().Model(&e).Where(query, args...).Updates(fields)
	return core.NewEither(&result.RowsAffected, core.NewIfError(result.Error))
}
