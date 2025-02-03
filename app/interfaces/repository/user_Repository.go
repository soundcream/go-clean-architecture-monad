package repository

import (
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/interfaces/db"
)

type UserRepository interface {
	db.ReadOnlyRepository[entity.User]
	GetById(id int) *entity.User
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

func (r *userRepository) GetById(id int) *entity.User {
	return r.FindById(id)
}

func (r *userRepository) GetByIdIncludeUserGroup(id int) *entity.User {
	return r.FindByIdIncludes(id, "UserGroup")
}
