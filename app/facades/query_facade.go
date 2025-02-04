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

func NewQueryFacade(repo repository.UserRepository) QueryFacade {
	return QueryFacade{
		userRepository: repo,
	}
}

func (f QueryFacade) GetUser() base.Either[entity.User, base.ErrContext] {
	user := f.userRepository.GetSpecialLogicUser(1)
	user2 := f.userRepository.FindById(1)
	up2 := f.userRepository.Count("point > ?", 1)
	w2 := f.userRepository.Where("name LIKE ?", "%te%")

	ug1 := f.userRepository.FindByIdPreloadInclude(3, entity.User{}.UserGroup, "is_active = ?", true)

	ug := f.userRepository.FindByIdPreload(3, util.Map("UserGroup"))
	fud := f.userRepository.FindByIdPreload(3, util.Map("UserGroup", "is_active = ?", true))
	//ug := f.userRepo.FindByIdLoad(3)

	// Use
	r0 := f.userRepository.Query().Where("point > 1").Order("point desc").FetchAll()
	r1 := f.userRepository.Query().Where("point > ?", 1).Order("point desc").FetchAll()
	r2 := f.userRepository.Query().Where("point > ?", 1).Order("point desc").Fetch()
	r3 := f.userRepository.Query().
		Preload("UserGroup").
		Where("point > ?", 1).
		Order("point desc").
		FetchAll()

	p1 := f.userRepository.BuildQueryPagination().
		PreloadWith(entity.User{}.UserGroup).
		Where("point > ?", 1).
		Order("point desc").
		ToPaging(1, 0)

	p2 := f.userRepository.BuildQueryPagination().
		PreloadWith(entity.User{}.UserGroup, "is_active = ?", true).
		Where("point > ?", 1).
		Order("point desc").
		ToPaging(1, 0)

	fmt.Println(user2, up2, w2, ug, fud, r1, r2, r3, r0, ug1, p1, p2)
	return base.RightEither[entity.User, base.ErrContext](*user)
}
