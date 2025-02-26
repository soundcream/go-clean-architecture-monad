package facades

import (
	"fmt"
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/global"
	"n4a3/clean-architecture/app/base/util"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/integrates/repository"
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
	user01 := f.userRepository.FindById(1)
	user01.SetUpdater("abc")
	count01 := f.userRepository.Count("point > ?", 1)
	user02 := f.userRepository.Where("name LIKE ?", "%te%")
	user03 := f.userRepository.FindByIdPreloadInclude(3, entity.User{}.UserGroup, "is_active = ?", true)
	user04 := f.userRepository.FindByIdPreload(3, util.Map("UserGroup", "is_active = ?"))
	// Query
	user05 := f.userRepository.Query().
		Preload("UserGroup").
		Where("point > ?", 1).
		Order("point desc").
		FetchAll()
	// Paging
	user06 := f.userRepository.BuildQueryPagination().
		PreloadWith(entity.User{}.UserGroup, "is_active = ?", true).
		Where("point > ?", 1).
		Order("point desc").
		ToPaging(1, 0)
	fmt.Println(user01, user02, user03, user04, user05, user06, count01)
	return base.NewRightEither[entity.User, base.ErrContext](user01)
}

func (f QueryFacade) GetUserById(id int) base.Either[entity.User, base.ErrContext] {
	result := f.userRepository.FindById(id)
	return base.NewRightEither[entity.User, base.ErrContext](result)
}

func (f QueryFacade) SearchUsers(keyword string, limit, offset int) base.Either[global.PagingModel[entity.User], base.ErrContext] {
	result := f.userRepository.BuildQueryPagination().
		PreloadWith(entity.User{}.UserGroup, "is_active = ?", true).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
		Order("point desc").
		ToPaging(limit, offset)
	return base.NewRightEither[global.PagingModel[entity.User], base.ErrContext](&result)
}
