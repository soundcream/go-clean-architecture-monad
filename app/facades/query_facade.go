package facades

import (
	"fmt"
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/util"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/interfaces/repository"
)

type QueryFacade struct {
	userRepository repository.UserRepository
}

func NewQueryFacade(userRepository2 repository.UserRepository) QueryFacade {
	return QueryFacade{
		userRepository: userRepository2,
	}
}

func (f QueryFacade) GetUser() base.Either[entity.User, base.ErrContext] {
	user := f.userRepository.GetById(1)
	user2 := f.userRepository.FindById(1)
	up := f.userRepository.Count("point > 1")
	up2 := f.userRepository.CountWith("point > ?", 1)
	w1 := f.userRepository.Where("length(name) > 4")
	w2 := f.userRepository.WhereWith("name LIKE ?", "%te%")

	ug := f.userRepository.FindByIdPreload(3, util.Map("UserGroup"))

	r1 := f.userRepository.Query().Where("point > ?", "1").Order("point desc").Execute()

	fud := f.userRepository.FindByIdPreload(3, util.Map("UserGroup", "is_active = ?", true))
	//ug := f.userRepo.FindByIdLoad(3)

	fmt.Println(user2, up, up2, w1, w2, ug, fud, r1)
	return base.RightEither[entity.User, base.ErrContext](*user)
}
