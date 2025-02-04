package db

import (
	"gorm.io/gorm"
	"math/big"
	"n4a3/clean-architecture/app/base/generic"
	"n4a3/clean-architecture/app/base/util"
	stringutil "n4a3/clean-architecture/app/base/util/string"
	"n4a3/clean-architecture/app/domain/entity"
)

type ReadOnlyRepository[Entity entity.IBaseEntity] interface {
	Query() QueryContext[Entity]
	BuildQueryPagination() QueryContext[Entity]

	Count(query interface{}, args ...interface{}) *int
	CountBig(query interface{}, args ...interface{}) *int64
	Sum(query interface{}, args ...interface{}) *int
	SumBig(query interface{}, args ...interface{}) *big.Float

	FindById(id int) *Entity
	FindBy(query interface{}, args ...interface{}) *Entity
	FindOrderBy(order interface{}, query interface{}, args ...interface{}) *Entity

	// FindByIdPreload Ex: (3, util.Map("UserGroup"))
	FindByIdPreload(id int, preloads *map[string][]interface{}) *Entity
	// FindByIdPreloadInclude Ex: (1, entity.User{}.UserGroup) or (1, entity.User{}.UserGroup, "is_active = ?", true)
	FindByIdPreloadInclude(id int, field interface{}, args ...interface{}) *Entity
	// Where
	Where(query interface{}, args ...interface{}) []Entity
	WhereOrderBy(order interface{}, query interface{}, args ...interface{}) []Entity
}

type readOnlyRepository[Entity entity.IBaseEntity] struct {
	UoW       QueryUnitOfWork
	TableName string
}

func NewReadOnlyRepository[Entity entity.IBaseEntity](uow *QueryUnitOfWork) ReadOnlyRepository[Entity] {
	tableName := generic.GetFieldTagByName[Entity]("BaseEntity", "table-name")
	return &readOnlyRepository[Entity]{
		UoW:       *uow,
		TableName: tableName,
	}
}

func (repo *readOnlyRepository[Entity]) Query() QueryContext[Entity] {
	return NewQueryContext[Entity](repo.UoW.DB())
}

func (repo *readOnlyRepository[Entity]) BuildQueryPagination() QueryContext[Entity] {
	var e Entity
	return NewQueryContext[Entity](repo.UoW.DB().Model(&e).Session(&gorm.Session{}))
}

func (repo *readOnlyRepository[Entity]) FindByIdPreload(id int, preloads *map[string][]interface{}) *Entity {
	var result Entity
	query := repo.UoW.DB().Model(&result)
	if preloads != nil {
		for k, preload := range *preloads {
			query = query.Preload(k, preload...)
		}
	}

	query.Take(&result, id)
	return &result
}

func (repo *readOnlyRepository[Entity]) FindByIdPreloadInclude(id int, field interface{}, args ...interface{}) *Entity {
	var result Entity
	query := repo.UoW.DB().Model(&result)
	fieldName := util.GetFieldName(field)
	if !stringutil.IsNullOrEmpty(fieldName) {
		query = query.Preload(fieldName, args...)
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
	repo.init().Take(&result, id)
	return &result
}

func (repo *readOnlyRepository[Entity]) FindBy(query interface{}, args ...interface{}) *Entity {
	if args == nil {
		panic("method FindBy, args cannot be nil")
	}
	var result Entity
	repo.init().Where(query, args...).First(&result)
	return &result
}

func (repo *readOnlyRepository[Entity]) FindOrderBy(order interface{}, query interface{}, args ...interface{}) *Entity {
	if args == nil {
		panic("method FindOrderBy, args cannot be nil")
	}
	var result Entity
	repo.init().Where(query, args...).Order(order).First(&result)
	return &result
}

func (repo *readOnlyRepository[Entity]) Where(query interface{}, args ...interface{}) []Entity {
	var result []Entity
	repo.init().Where(query, args...).Find(&result)
	return result
}

func (repo *readOnlyRepository[Entity]) WhereOrderBy(order interface{}, query interface{}, args ...interface{}) []Entity {
	var result []Entity
	repo.init().Find(&result).Where(query, args...).Order(order)
	return result
}

func (repo *readOnlyRepository[Entity]) Count(query interface{}, args ...interface{}) *int {
	return util.PtrToPtr(repo.CountBig(query, args...), func(i int64) int {
		return int(i)
	})
}

func (repo *readOnlyRepository[Entity]) CountBig(query interface{}, args ...interface{}) *int64 {
	result := new(int64)
	repo.init().Where(query, args...).Count(result)
	return result
}

func (repo *readOnlyRepository[Entity]) Sum(query interface{}, args ...interface{}) *int {
	result := new(int)
	repo.init().Where(query, args...).Scan(&result)
	return result
}

func (repo *readOnlyRepository[Entity]) SumBig(query interface{}, args ...interface{}) *big.Float {
	result := new(big.Float)
	repo.init().Where(query, args...).Scan(&result)
	return result
}

func (repo *readOnlyRepository[Entity]) init() *gorm.DB {
	return repo.UoW.DB().Table(repo.TableName)
}
