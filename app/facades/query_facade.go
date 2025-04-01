package facades

import (
	"fmt"
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/core/global"
	"n4a3/clean-architecture/app/core/util"
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

func (f QueryFacade) GetUser() core.Either[entity.User, core.ErrContext] {
	user01 := f.userRepository.FindById(1)
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
	return user01
}

func (f QueryFacade) GetUserById(id int) core.Either[entity.User, core.ErrContext] {
	result := f.userRepository.FindById(id)
	return result
}

func (f QueryFacade) SearchUsers(keyword string, limit, offset int) core.Either[global.PagingModel[entity.User], core.ErrContext] {
	result := f.userRepository.BuildQueryPagination().
		PreloadWith(entity.User{}.UserGroup, "is_active = ?", true).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
		Where("username LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
		Order("point desc").
		ToPaging(limit, offset)
	return core.NewRightEither[global.PagingModel[entity.User], core.ErrContext](&result)
}
