package repository

import (
	"n4a3/clean-architecture/app/base/util"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/integrates/db"
)

type UserRepository interface {
	db.ReadOnlyRepository[entity.User]
	GetSpecialLogicUser(id int) *entity.User
}

type userRepository struct {
	db.Repository[entity.User]
	db.ReadOnlyRepository[entity.User]
}

func NewUserRepository(rUoW *db.QueryUnitOfWork, uow *db.CommandUnitOfWork) UserRepository {
	readOnlyRepository := db.NewReadOnlyRepository[entity.User](rUoW)
	repository := db.NewRepository[entity.User](uow)
	return &userRepository{
		Repository:         repository,
		ReadOnlyRepository: readOnlyRepository,
	}
}

func (r *userRepository) GetSpecialLogicUser(id int) *entity.User {
	return r.FindByIdPreload(id, util.Map("UserGroup"))
}
