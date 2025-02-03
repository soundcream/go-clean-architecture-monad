package facades

import (
	"fmt"
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/interfaces/db"
	"n4a3/clean-architecture/app/interfaces/repository"
)

type QueryFacade struct {
	userRepository  db.ReadOnlyRepository[entity.User]
	userRepository2 repository.UserRepository
}

func NewQueryFacade(userRepository2 repository.UserRepository, userRepository db.ReadOnlyRepository[entity.User]) QueryFacade {
	return QueryFacade{
		userRepository2: userRepository2,
		userRepository:  userRepository,
	}
}

func (f QueryFacade) GetUser() base.Either[entity.User, base.ErrContext] {
	user := f.userRepository2.GetById(1)
	user2 := f.userRepository.FindById(1)
	up := f.userRepository.Count("point > ?", 1)
	up2 := f.userRepository.BigCount("point > ?", 1)
	w1 := f.userRepository.Where("length(name) > 4")
	w2 := f.userRepository.WhereWith("name LIKE ?", "%te%")
	fmt.Println(user2, up, up2, w1, w2)
	return base.RightEither[entity.User, base.ErrContext](*user)
}
