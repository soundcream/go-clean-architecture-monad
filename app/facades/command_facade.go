package facades

import (
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/core/either"
	"n4a3/clean-architecture/app/core/util"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/integrates/dto"
	"n4a3/clean-architecture/app/integrates/repository"
)

type CommandFacade interface {
	Insert(user dto.UserDto) core.Either[dto.UserDto, core.ErrContext]
	Update(user dto.CommandDto[dto.UserDto]) core.Either[dto.CommandDto[dto.UserDto], core.ErrContext]
	UpdateTest() core.Either[core.Unit, core.ErrContext]
	UpdateWhere(id int) core.Either[int, core.ErrContext]
	UpdateFieldWhere(id int) core.Either[int, core.ErrContext]
	Delete(id int) core.Either[int, core.ErrContext]
}

type commandFacade struct {
	userRepository repository.UserRepository
}

func NewCommandFacade(repo repository.UserRepository) CommandFacade {
	return &commandFacade{
		userRepository: repo,
	}
}

func (c commandFacade) Insert(user dto.UserDto) core.Either[dto.UserDto, core.ErrContext] {

	u := entity.User{
		BaseEntity:  entity.NewBase(),
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		Point:       user.Point,
		UserGroupId: user.UserGroupId,
		UserGroup:   nil,
	}
	u.SetInserter("system")
	res := c.userRepository.Insert(&u)
	return MapToResult[dto.UserDto](res, user)
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

func (c commandFacade) Update(user dto.CommandDto[dto.UserDto]) core.Either[dto.CommandDto[dto.UserDto], core.ErrContext] {
	u := util.MapFrom[entity.User](user.Model)
	u.BaseEntity = entity.NewBaseUpdateWithId(user.Id, "system_update")
	result := c.userRepository.Update(user.Id, *u)
	return MapToResult[dto.CommandDto[dto.UserDto]](result, user)
}

func (c commandFacade) UpdateTest() core.Either[core.Unit, core.ErrContext] {
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
	result := c.userRepository.Insert(&u)
	return either.Map[int64, core.Unit, core.ErrContext](result, toUnit)
}

func (c commandFacade) Delete(id int) core.Either[int, core.ErrContext] {
	result := c.userRepository.DeleteById(id)
	return either.Map[int64, int, core.ErrContext](result, toInt)
}

func (c commandFacade) UpdateWhere(id int) core.Either[int, core.ErrContext] {
	u := entity.User{
		Email: "email@email.com",
	}
	result := c.userRepository.UpdatesWhere(u, "id = ?", id)
	return either.Map[int64, int, core.ErrContext](result, toInt)
}

func (c commandFacade) UpdateFieldWhere(id int) core.Either[int, core.ErrContext] {
	result := c.userRepository.UpdatesFieldsWhere(map[string]interface{}{"email": "email@email-test.com", "point": 200}, "id = ?", id)
	return either.Map[int64, int, core.ErrContext](result, toInt)
}

func toUnit(err *core.ErrContext, input *int64) core.Either[core.Unit, core.ErrContext] {
	if err != nil {
		return core.LeftEither[core.Unit, core.ErrContext](*err)
	} else if input != nil && *input <= 0 {
		return core.LeftEither[core.Unit, core.ErrContext](core.NewErrorCode(core.Invalid))
	}
	return core.RightEither[core.Unit, core.ErrContext](core.Unit{})
}

func toInt(err *core.ErrContext, input *int64) core.Either[int, core.ErrContext] {
	if err != nil {
		return core.LeftEither[int, core.ErrContext](*err)
	} else if input != nil && *input <= 0 {
		return core.LeftEither[int, core.ErrContext](core.NewErrorCode(core.Invalid))
	}
	return core.RightEither[int, core.ErrContext](1)
}

func MapToResult[T any](commandResult core.Either[int64, core.ErrContext], result T) core.Either[T, core.ErrContext] {
	if commandResult.IsRight() {
		return core.RightEither[T, core.ErrContext](result)
	} else if commandResult.IsLeft() {
		return core.LeftEither[T, core.ErrContext](*commandResult.Left)
	}
	return core.LeftEither[T, core.ErrContext](core.NewErrorCode(core.Invalid))
}
