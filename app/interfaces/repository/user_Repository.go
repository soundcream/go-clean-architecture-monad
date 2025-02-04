package repository

import (
	"n4a3/clean-architecture/app/base/util"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/interfaces/db"
)

type UserRepository interface {
	db.ReadOnlyRepository[entity.User]
	GetSpecialLogicUser(id int) *entity.User
}

type userRepository struct {
	db.ReadOnlyRepository[entity.User]
	repository *db.ReadOnlyRepository[entity.User]
}

func NewUserRepository(uow *db.QueryUnitOfWork) UserRepository {
	repository := db.NewReadOnlyRepository[entity.User](uow)
	return &userRepository{
		repository:         &repository,
		ReadOnlyRepository: repository,
	}
}

func (r *userRepository) GetSpecialLogicUser(id int) *entity.User {
	return r.FindByIdPreload(id, util.Map("UserGroup"))
}
