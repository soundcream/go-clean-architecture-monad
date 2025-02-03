package db

import "n4a3/clean-architecture/app/domain/entity"

type Repository[Entity entity.IBaseEntity] interface {
	ReadOnlyRepository[Entity]
}

type repository[Entity entity.IBaseEntity] struct {
	UoW       CommandUnitOfWork
	TableName string
}
