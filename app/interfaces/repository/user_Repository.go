package repository

import (
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/interfaces/db"
)

type UserRepository interface {
	GetById(id int) *entity.User
}

type userRepository struct {
	repository *db.ReadOnlyRepository[entity.User]
}

func NewUserRepository(uow *db.QueryUnitOfWork) UserRepository {
	repository := db.NewReadOnlyRepository[entity.User](uow)
	return &userRepository{
		repository: &repository,
	}
}

func (r userRepository) GetById(id int) *entity.User {
	return r.repository.FindById(id)
}
