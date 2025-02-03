package db

import "n4a3/clean-architecture/app/domain/entity"

type IRepository interface {
}

type Repository[Entity entity.IBaseEntity] struct {
	UoW CommandUnitOfWork
}
