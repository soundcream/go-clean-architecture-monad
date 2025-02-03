package facades

import (
	"fmt"
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/interfaces/db"
	"n4a3/clean-architecture/app/interfaces/repository"
)

type QueryFacade struct {
	//userRepo       db.ReadOnlyRepository[entity.User]
	userRepository repository.UserRepository
}

func NewQueryFacade(userRepository2 repository.UserRepository, userRepository db.ReadOnlyRepository[entity.User]) QueryFacade {
	return QueryFacade{
		userRepository: userRepository2,
		//userRepo:       userRepository,
	}
}

func (f QueryFacade) GetUser() base.Either[entity.User, base.ErrContext] {
	user := f.userRepository.GetById(1)
	user2 := f.userRepository.FindById(1)
	up := f.userRepository.Count("point > 1")
	up2 := f.userRepository.CountWith("point > ?", 1)
	w1 := f.userRepository.Where("length(name) > 4")
	w2 := f.userRepository.WhereWith("name LIKE ?", "%te%")

	ug := f.userRepository.FindByIdIncludes(3, "UserGroup")

	wt := f.userRepository.WhereTest()
	//ug := f.userRepo.FindByIdLoad(3)

	fmt.Println(user2, up, up2, w1, w2, ug, wt)
	return base.RightEither[entity.User, base.ErrContext](*user)
}
