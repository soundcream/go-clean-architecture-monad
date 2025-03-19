package facades

import (
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/either"
	"n4a3/clean-architecture/app/base/util"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/integrates/dto"
	"n4a3/clean-architecture/app/integrates/repository"
)

type CommandFacade interface {
	Insert(user dto.UserDto) base.Either[dto.UserDto, base.ErrContext]
	Update(user dto.CommandDto[dto.UserDto]) base.Either[dto.CommandDto[dto.UserDto], base.ErrContext]
	UpdateTest() base.Either[base.Unit, base.ErrContext]
	UpdateWhere(id int) base.Either[int, base.ErrContext]
	UpdateFieldWhere(id int) base.Either[int, base.ErrContext]
	Delete(id int) base.Either[int, base.ErrContext]
}

type commandFacade struct {
	userRepository repository.UserRepository
}

func NewCommandFacade(repo repository.UserRepository) CommandFacade {
	return &commandFacade{
		userRepository: repo,
	}
}

func (c commandFacade) Insert(user dto.UserDto) base.Either[dto.UserDto, base.ErrContext] {

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

func (c commandFacade) Update(user dto.CommandDto[dto.UserDto]) base.Either[dto.CommandDto[dto.UserDto], base.ErrContext] {
	u := util.MapFrom[entity.User](user.Model)
	u.BaseEntity = entity.NewBaseUpdateWithId(user.Id, "system_update")
	result := c.userRepository.Update(user.Id, *u)
	//result := c.userRepository.UpdateAllFields(u)
	return MapToResult[dto.CommandDto[dto.UserDto]](result, user)
}

func (c commandFacade) UpdateTest() base.Either[base.Unit, base.ErrContext] {
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
	return either.Map[int64, base.Unit, base.ErrContext](result, toUnit)
}

func (c commandFacade) Delete(id int) base.Either[int, base.ErrContext] {
	result := c.userRepository.DeleteById(id)
	return either.Map[int64, int, base.ErrContext](result, toInt)
}

func (c commandFacade) UpdateWhere(id int) base.Either[int, base.ErrContext] {
	u := entity.User{
		Email: "email@email.com",
	}
	result := c.userRepository.UpdatesWhere(u, "id = ?", id)
	return either.Map[int64, int, base.ErrContext](result, toInt)
}

func (c commandFacade) UpdateFieldWhere(id int) base.Either[int, base.ErrContext] {
	result := c.userRepository.UpdatesFieldsWhere(map[string]interface{}{"email": "email@email-test.com", "point": 200}, "id = ?", id)
	return either.Map[int64, int, base.ErrContext](result, toInt)
}

func toUnit(err *base.ErrContext, input *int64) base.Either[base.Unit, base.ErrContext] {
	if err != nil {
		return base.LeftEither[base.Unit, base.ErrContext](*err)
	} else if input != nil && *input <= 0 {
		return base.LeftEither[base.Unit, base.ErrContext](base.NewErrorCode(base.Invalid))
	}
	return base.RightEither[base.Unit, base.ErrContext](base.Unit{})
}

func toInt(err *base.ErrContext, input *int64) base.Either[int, base.ErrContext] {
	if err != nil {
		return base.LeftEither[int, base.ErrContext](*err)
	} else if input != nil && *input <= 0 {
		return base.LeftEither[int, base.ErrContext](base.NewErrorCode(base.Invalid))
	}
	return base.RightEither[int, base.ErrContext](1)
}

func MapToResult[T any](commandResult base.Either[int64, base.ErrContext], result T) base.Either[T, base.ErrContext] {
	if commandResult.IsRight() {
		return base.RightEither[T, base.ErrContext](result)
	} else if commandResult.IsLeft() {
		return base.LeftEither[T, base.ErrContext](*commandResult.Left)
	}
	return base.LeftEither[T, base.ErrContext](base.NewErrorCode(base.Invalid))
}
