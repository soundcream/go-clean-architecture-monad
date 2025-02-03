package db

import (
	"math/big"
	"n4a3/clean-architecture/app/base/generic"
	"n4a3/clean-architecture/app/domain/entity"
)

type ReadOnlyRepository[Entity entity.IBaseEntity] interface {
	FindByIdIncludes(id int, preloads ...string) *Entity
	FindById(id int) *Entity
	FindBy(query interface{}, args ...interface{}) *Entity
	FindOrderBy(query interface{}, order interface{}, args ...interface{}) *Entity
	Where(query interface{}) []Entity
	WhereWith(query interface{}, args ...interface{}) []Entity
	WhereOrderBy(query interface{}, order interface{}) []Entity
	WhereWithOrderBy(query interface{}, order interface{}, args ...interface{}) []Entity
	Count(query interface{}) *int64
	CountWith(query interface{}, args ...interface{}) *int64
	Sum(query interface{}) *int
	SumWith(query interface{}, args ...interface{}) *int
	SumBig(query interface{}) *big.Float
	SumBigWith(query interface{}, args ...interface{}) *big.Float

	WhereTest() []Entity
}

type readOnlyRepository[Entity entity.IBaseEntity] struct {
	UoW       QueryUnitOfWork
	TableName string
}

func NewReadOnlyRepository[Entity entity.IBaseEntity](uow *QueryUnitOfWork) ReadOnlyRepository[Entity] {
	tableName := generic.GetTagByName[Entity]("table-name")
	return &readOnlyRepository[Entity]{
		UoW:       *uow,
		TableName: tableName,
	}
}

func (repo *readOnlyRepository[Entity]) FindByIdIncludes(id int, preloads ...string) *Entity {
	var result Entity
	query := repo.UoW.DB().Model(&result)
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	query.Take(&result, id)
	return &result
}

func (repo *readOnlyRepository[Entity]) FindByIdLeftJoins(id int, joins ...string) *Entity {
	var result Entity
	query := repo.UoW.DB().Model(&result)
	for _, join := range joins {
		query = query.Joins(join)
	}
	query.Take(&result, id)
	return &result
}

func (repo *readOnlyRepository[Entity]) FindByIdInnerJoins(id int, joins ...string) *Entity {
	var result Entity
	query := repo.UoW.DB().Model(&result)
	for _, join := range joins {
		query = query.InnerJoins(join)
	}
	query.Take(&result, id)
	return &result
}

func (repo *readOnlyRepository[Entity]) FindById(id int) *Entity {
	var result Entity
	repo.UoW.DB().Table(repo.TableName).Take(&result, id)
	return &result
}

func (repo *readOnlyRepository[Entity]) FindBy(query interface{}, args ...interface{}) *Entity {
	if args == nil {
		panic("method FindBy, args cannot be nil")
	}
	var result Entity
	repo.UoW.DB().Table(repo.TableName).Where(query, args).First(&result)
	return &result
}

func (repo *readOnlyRepository[Entity]) FindOrderBy(query interface{}, order interface{}, args ...interface{}) *Entity {
	if args == nil {
		panic("method FindOrderBy, args cannot be nil")
	}
	var result Entity
	repo.UoW.DB().Table(repo.TableName).Where(query, args).Order(order).First(&result)
	return &result
}

func (repo *readOnlyRepository[Entity]) Where(query interface{}) []Entity {
	var result []Entity
	repo.UoW.DB().Table(repo.TableName).Where(query).Find(&result)
	return result
}

func (repo *readOnlyRepository[Entity]) WhereWith(query interface{}, args ...interface{}) []Entity {
	if args == nil {
		panic("method WhereWith, args cannot be nil")
	}
	var result []Entity
	repo.UoW.DB().Table(repo.TableName).Where(query, args).Find(&result)
	return result
}

func (repo *readOnlyRepository[Entity]) WhereOrderBy(query interface{}, order interface{}) []Entity {
	var result []Entity
	repo.UoW.DB().Table(repo.TableName).Find(&result).Where(query).Order(order)
	return result
}

func (repo *readOnlyRepository[Entity]) WhereWithOrderBy(query interface{}, order interface{}, args ...interface{}) []Entity {
	if args == nil {
		panic("method WhereWithOrderBy, args cannot be nil")
	}
	var result []Entity
	repo.UoW.DB().Table(repo.TableName).Find(&result).Where(query, args).Order(order)
	return result
}

func (repo *readOnlyRepository[Entity]) Query() QueryContext {
	return NewQueryContext(repo.UoW.DB())
}

func (repo *readOnlyRepository[Entity]) WhereTest() []Entity {
	var result []Entity
	repo.Query().Where("point > ?", "1").Order("point desc").Find(&result)
	return result
}

func (repo *readOnlyRepository[Entity]) Count(query interface{}) *int64 {
	result := new(int64)
	repo.UoW.DB().Table(repo.TableName).Where(query).Count(result)
	return result
}

func (repo *readOnlyRepository[Entity]) CountWith(query interface{}, args ...interface{}) *int64 {
	if args == nil {
		panic("method CountWith, args cannot be nil")
	}
	result := new(int64)
	repo.UoW.DB().Table(repo.TableName).Where(query, args).Count(result)
	return result
}

func (repo *readOnlyRepository[Entity]) Sum(query interface{}) *int {
	result := new(int)
	repo.UoW.DB().Table(repo.TableName).Where(query).Scan(&result)
	return result
}

func (repo *readOnlyRepository[Entity]) SumWith(query interface{}, args ...interface{}) *int {
	if args == nil {
		panic("method SumWith, args cannot be nil")
	}
	result := new(int)
	repo.UoW.DB().Table(repo.TableName).Where(query, args).Scan(&result)
	return result
}

func (repo *readOnlyRepository[Entity]) SumBig(query interface{}) *big.Float {
	result := new(big.Float)
	repo.UoW.DB().Table(repo.TableName).Where(query).Scan(&result)
	return result
}

func (repo *readOnlyRepository[Entity]) SumBigWith(query interface{}, args ...interface{}) *big.Float {
	if args == nil {
		panic("method SumBigWith, args cannot be nil")
	}
	result := new(big.Float)
	repo.UoW.DB().Table(repo.TableName).Where(query, args).Scan(&result)
	return result
}
