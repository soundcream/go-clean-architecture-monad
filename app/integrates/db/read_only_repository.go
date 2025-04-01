package db

import (
	"fmt"
	"gorm.io/gorm"
	"math/big"
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/core/either"
	"n4a3/clean-architecture/app/core/util"
	stringutil "n4a3/clean-architecture/app/core/util/string"
	"n4a3/clean-architecture/app/domain/entity"
)

type ReadOnlyRepository[Entity entity.Entity] interface {
	Query() QueryContext[Entity]
	BuildQueryPagination() QueryContext[Entity]

	Count(query interface{}, args ...interface{}) core.Either[int, core.ErrContext]
	CountBig(query interface{}, args ...interface{}) core.Either[int64, core.ErrContext]
	Sum(query interface{}, args ...interface{}) core.Either[int, core.ErrContext]
	SumBig(query interface{}, args ...interface{}) core.Either[big.Float, core.ErrContext]

	FindById(id int) core.Either[Entity, core.ErrContext]
	FindBy(query interface{}, args ...interface{}) core.Either[Entity, core.ErrContext]
	FindOrderBy(order interface{}, query interface{}, args ...interface{}) core.Either[Entity, core.ErrContext]

	// FindByIdPreload Ex: (3, util.Map("UserGroup"))
	FindByIdPreload(id int, preloads *map[string][]interface{}) core.Either[Entity, core.ErrContext]
	// FindByIdPreloadInclude Ex: (1, entity.User{}.UserGroup) or (1, entity.User{}.UserGroup, "is_active = ?", true)
	FindByIdPreloadInclude(id int, field interface{}, args ...interface{}) core.Either[Entity, core.ErrContext]
	// Where
	Where(query interface{}, args ...interface{}) core.Either[[]Entity, core.ErrContext]
	WhereOrderBy(order interface{}, query interface{}, args ...interface{}) core.Either[[]Entity, core.ErrContext]
}

type readOnlyRepository[Entity entity.Entity] struct {
	UoW       QueryUnitOfWork
	TableName string
}

func NewReadOnlyRepository[Entity entity.Entity](uow *QueryUnitOfWork) ReadOnlyRepository[Entity] {
	//tableName := generic.GetFieldTagByName[Entity]("BaseEntity", "table-name")
	var e = *new(Entity)
	return &readOnlyRepository[Entity]{
		UoW:       *uow,
		TableName: e.TableName(),
	}
}

func (repo *readOnlyRepository[Entity]) Query() QueryContext[Entity] {
	return NewQueryContext[Entity](repo.UoW.DB())
}

func (repo *readOnlyRepository[Entity]) BuildQueryPagination() QueryContext[Entity] {
	var e Entity
	return NewQueryContext[Entity](repo.UoW.DB().Model(&e).Session(&gorm.Session{}))
}

func (repo *readOnlyRepository[Entity]) FindByIdPreload(id int, preloads *map[string][]interface{}) core.Either[Entity, core.ErrContext] {
	var e Entity
	query := repo.UoW.DB().Model(&e)
	if preloads != nil {
		for k, preload := range *preloads {
			query = query.Preload(k, preload...)
		}
	}
	result := query.Take(&e, id)
	if result.RowsAffected == 0 {
		return core.NewRightEither[Entity, core.ErrContext](nil)
	} else if result.Error != nil {
		return core.LeftEither[Entity, core.ErrContext](core.NewErrorWithMsg(core.Invalid, fmt.Sprintf("%s FindByIdPreload %d", util.New[Entity]().TableName(), id), result.Error))
	}
	return core.NewRightEither[Entity, core.ErrContext](&e)
}

func (repo *readOnlyRepository[Entity]) FindByIdPreloadInclude(id int, field interface{}, args ...interface{}) core.Either[Entity, core.ErrContext] {
	var e Entity
	query := repo.UoW.DB().Model(&e)
	fieldName := util.GetFieldName(field)
	if !stringutil.IsNullOrEmpty(fieldName) {
		query = query.Preload(fieldName, args...)
	}
	result := query.Take(&e, id)
	if result.RowsAffected == 0 {
		return core.NewRightEither[Entity, core.ErrContext](nil)
	} else if result.Error != nil {
		return core.LeftEither[Entity, core.ErrContext](core.NewErrorWithMsg(core.Invalid, fmt.Sprintf("%s FindByIdPreloadInclude", util.New[Entity]().TableName()), result.Error))
	}
	return core.NewRightEither[Entity, core.ErrContext](&e)
}

func (repo *readOnlyRepository[Entity]) FindByIdLeftJoins(id int, joins ...string) core.Either[Entity, core.ErrContext] {
	var e Entity
	query := repo.UoW.DB().Model(&e)
	for _, join := range joins {
		query = query.Joins(join)
	}
	result := query.Take(&e, id)
	if result.RowsAffected == 0 {
		return core.NewRightEither[Entity, core.ErrContext](nil)
	} else if result.Error != nil {
		return core.LeftEither[Entity, core.ErrContext](core.NewErrorWithMsg(core.Invalid, fmt.Sprintf("%s FindByIdLeftJoins", util.New[Entity]().TableName()), result.Error))
	}
	return core.NewRightEither[Entity, core.ErrContext](&e)
}

func (repo *readOnlyRepository[Entity]) FindByIdInnerJoins(id int, joins ...string) core.Either[Entity, core.ErrContext] {
	var e Entity
	query := repo.UoW.DB().Model(&e)
	for _, join := range joins {
		query = query.InnerJoins(join)
	}
	result := query.Take(&e, id)
	if result.RowsAffected == 0 {
		return core.NewRightEither[Entity, core.ErrContext](nil)
	} else if result.Error != nil {
		return core.LeftEither[Entity, core.ErrContext](core.NewErrorWithMsg(core.Invalid, fmt.Sprintf("%s FindByIdInnerJoins", util.New[Entity]().TableName()), result.Error))
	}
	return core.NewRightEither[Entity, core.ErrContext](&e)
}

func (repo *readOnlyRepository[Entity]) FindById(id int) core.Either[Entity, core.ErrContext] {
	var e Entity
	result := repo.init().First(&e, id)
	if result.RowsAffected == 0 {
		return core.NewRightEither[Entity, core.ErrContext](nil)
	} else if result.Error != nil {
		return core.LeftEither[Entity, core.ErrContext](core.NewErrorWithMsg(core.Invalid, fmt.Sprintf("%s FindById", util.New[Entity]().TableName()), result.Error))
	}
	return core.NewRightEither[Entity, core.ErrContext](&e)
}

func (repo *readOnlyRepository[Entity]) FindBy(query interface{}, args ...interface{}) core.Either[Entity, core.ErrContext] {
	if args == nil {
		panic("method FindBy, args cannot be nil")
	}
	var e Entity
	result := repo.init().Where(query, args...).First(&e)
	if result.RowsAffected == 0 {
		return core.NewRightEither[Entity, core.ErrContext](nil)
	} else if result.Error != nil {
		return core.LeftEither[Entity, core.ErrContext](core.NewErrorWithMsg(core.Invalid, fmt.Sprintf("%s FindBy", util.New[Entity]().TableName()), result.Error))
	}
	return core.NewRightEither[Entity, core.ErrContext](&e)
}

func (repo *readOnlyRepository[Entity]) FindOrderBy(order interface{}, query interface{}, args ...interface{}) core.Either[Entity, core.ErrContext] {
	if args == nil {
		panic("method FindOrderBy, args cannot be nil")
	}
	var e Entity
	result := repo.init().Where(query, args...).Order(order).First(&e)
	if result.RowsAffected == 0 {
		return core.NewRightEither[Entity, core.ErrContext](nil)
	} else if result.Error != nil {
		return core.LeftEither[Entity, core.ErrContext](core.NewErrorWithMsg(core.Invalid, fmt.Sprintf("%s FindOrderBy", util.New[Entity]().TableName()), result.Error))
	}
	return core.NewRightEither[Entity, core.ErrContext](&e)
}

func (repo *readOnlyRepository[Entity]) Where(query interface{}, args ...interface{}) core.Either[[]Entity, core.ErrContext] {
	var e []Entity
	result := repo.init().Where(query, args...).Find(&e)
	if result.RowsAffected == 0 {
		return core.NewRightEither[[]Entity, core.ErrContext](nil)
	} else if result.Error != nil {
		return core.LeftEither[[]Entity, core.ErrContext](core.NewErrorWithMsg(core.Invalid, fmt.Sprintf("%s Where", util.New[Entity]().TableName()), result.Error))
	}
	return core.NewRightEither[[]Entity, core.ErrContext](&e)
}

func (repo *readOnlyRepository[Entity]) WhereOrderBy(order interface{}, query interface{}, args ...interface{}) core.Either[[]Entity, core.ErrContext] {
	var e []Entity
	result := repo.init().Find(&e).Where(query, args...).Order(order)
	if result.RowsAffected == 0 {
		return core.NewRightEither[[]Entity, core.ErrContext](nil)
	} else if result.Error != nil {
		return core.LeftEither[[]Entity, core.ErrContext](core.NewErrorWithMsg(core.Invalid, fmt.Sprintf("%s WhereOrderBy", util.New[Entity]().TableName()), result.Error))
	}
	return core.NewRightEither[[]Entity, core.ErrContext](&e)
}

func (repo *readOnlyRepository[Entity]) Count(query interface{}, args ...interface{}) core.Either[int, core.ErrContext] {
	count := repo.CountBig(query, args...)
	return either.Bind(count, func(i int64) core.Either[int, core.ErrContext] {
		return core.NewRightEither[int, core.ErrContext](util.ToPtr(int(i)))
	})
}

func (repo *readOnlyRepository[Entity]) CountBig(query interface{}, args ...interface{}) core.Either[int64, core.ErrContext] {
	e := new(int64)
	result := repo.init().Where(query, args...).Count(e)
	if result.RowsAffected == 0 {
		return core.NewRightEither[int64, core.ErrContext](nil)
	} else if result.Error != nil {
		return core.LeftEither[int64, core.ErrContext](core.NewErrorWithMsg(core.Invalid, fmt.Sprintf("%s Count or CountBig", util.New[Entity]().TableName()), result.Error))
	}
	return core.NewRightEither[int64, core.ErrContext](e)
}

func (repo *readOnlyRepository[Entity]) Sum(query interface{}, args ...interface{}) core.Either[int, core.ErrContext] {
	e := new(int)
	result := repo.init().Where(query, args...).Scan(&e)
	if result.RowsAffected == 0 {
		return core.NewRightEither[int, core.ErrContext](nil)
	} else if result.Error != nil {
		return core.LeftEither[int, core.ErrContext](core.NewErrorWithMsg(core.Invalid, fmt.Sprintf("%s Sum", util.New[Entity]().TableName()), result.Error))
	}
	return core.NewRightEither[int, core.ErrContext](e)
}

func (repo *readOnlyRepository[Entity]) SumBig(query interface{}, args ...interface{}) core.Either[big.Float, core.ErrContext] {
	e := new(big.Float)
	result := repo.init().Where(query, args...).Scan(&e)
	if result.RowsAffected == 0 {
		return core.NewRightEither[big.Float, core.ErrContext](nil)
	} else if result.Error != nil {
		return core.LeftEither[big.Float, core.ErrContext](core.NewErrorWithMsg(core.Invalid, fmt.Sprintf("%s SumBig", util.New[Entity]().TableName()), result.Error))
	}
	return core.NewRightEither[big.Float, core.ErrContext](e)
}

func (repo *readOnlyRepository[Entity]) init() *gorm.DB {
	return repo.UoW.DB().Table(repo.TableName)
}
