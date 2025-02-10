package facades

import (
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/either"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/integrates/repository"
)

type CommandFacade interface {
	Insert() base.Either[base.Unit, base.ErrContext]
}

type commandFacade struct {
	userRepository repository.UserRepository
}

func NewCommandFacade(repo repository.UserRepository) CommandFacade {
	return &commandFacade{
		userRepository: repo,
	}
}

func (c commandFacade) Insert() base.Either[base.Unit, base.ErrContext] {

	u := entity.User{
		BaseEntity:  entity.NewBase(),
		Name:        "aaa",
		Username:    "user_aa",
		Email:       "user_aa@email",
		Point:       nil,
		UserGroupId: nil,
		UserGroup:   nil,
	}
	u.SetInserter("system")
	res := c.userRepository.Insert(&u)
	return either.Map[int64, base.Unit, base.ErrContext](res, toUnit)

	//users := []entity.User{
	//	entity.User{},
	//	entity.User{},
	//}
	//c.userRepository.BulkInsert(&users)

	// Delete
	// Delete ById

	//c.userRepository.Update()
	//c.userRepository.Updates()
	//c.userRepository.UpdateWhere()
}

func toUnit(err *base.ErrContext, input *int64) base.Either[base.Unit, base.ErrContext] {
	if err != nil {
		return base.LeftEither[base.Unit, base.ErrContext](*err)
	}
	return base.RightEither[base.Unit, base.ErrContext](base.Unit{})
}
