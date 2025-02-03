package db

import (
	"n4a3/clean-architecture/app/base/generic"
	"n4a3/clean-architecture/app/domain/entity"
)

type ReadOnlyRepository[Entity entity.IBaseEntity] struct {
	UoW       QueryUnitOfWork
	TableName string
}

func NewReadOnlyRepository[Entity entity.IBaseEntity](uow *QueryUnitOfWork) ReadOnlyRepository[Entity] {
	tableName := generic.GetTagByName[Entity]("table-name")
	return ReadOnlyRepository[Entity]{
		UoW:       *uow,
		TableName: tableName,
	}
}

func (repo ReadOnlyRepository[Entity]) FindById(id int) *Entity {
	var result Entity
	repo.UoW.DB().Table(repo.TableName).Take(&result, id)
	return &result
}

func (repo ReadOnlyRepository[Entity]) FindBy(query interface{}, args ...interface{}) *Entity {
	var result Entity
	repo.UoW.DB().Table(repo.TableName).Where(query, args).First(&result)
	return &result
}

func (repo ReadOnlyRepository[Entity]) Where(query interface{}) []Entity {
	var result []Entity
	repo.UoW.DB().Table(repo.TableName).Where(query).Find(&result)
	return result
}

func (repo ReadOnlyRepository[Entity]) WhereWith(query interface{}, args ...interface{}) []Entity {
	var result []Entity
	repo.UoW.DB().Table(repo.TableName).Where(query, args).Find(&result)
	return result
}

func (repo ReadOnlyRepository[Entity]) FindOrderBy(query interface{}, order interface{}, args ...interface{}) *Entity {
	var result Entity
	repo.UoW.DB().Table(repo.TableName).Where(query, args).Order(order).First(&result)
	return &result
}

func (repo ReadOnlyRepository[Entity]) WhereOrderBy(query interface{}, order interface{}) []Entity {
	var result []Entity
	repo.UoW.DB().Table(repo.TableName).Find(&result).Where(query).Order(order)
	return result
}

func (repo ReadOnlyRepository[Entity]) WhereWithOrderBy(query interface{}, order interface{}, args ...interface{}) []Entity {
	var result []Entity
	repo.UoW.DB().Table(repo.TableName).Find(&result).Where(query, args).Order(order)
	return result
}

func (repo ReadOnlyRepository[Entity]) Count(query interface{}, args ...interface{}) int {
	var result int
	repo.UoW.DB().Table(repo.TableName).Where(query, args).First(&result)
	return result
}

func (repo ReadOnlyRepository[Entity]) BigCount(query interface{}, args ...interface{}) int64 {
	var result int64
	repo.UoW.DB().Table(repo.TableName).Where(query, args).First(&result)
	return result
}
