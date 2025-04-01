package facades

import (
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/core/either"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/integrates/repository"
)

type UserFacade interface {
	CreateUser(name, email string) (*entity.User, error)
	GetUserById(id int) (*entity.User, error)
	ListUsers() ([]entity.User, error)
	TestThen() *entity.Txn
	TestValidate(u *entity.User) *core.ErrContext
}

type userFacade struct {
	repo repository.UserRepository
}

func NewUserFacade(repo repository.UserRepository) UserFacade {
	return &userFacade{repo: repo}
}

func (f *userFacade) CreateUser(name, email string) (*entity.User, error) {
	return &entity.User{Name: name, Email: email}, nil
}

func (f *userFacade) ListUsers() ([]entity.User, error) {
	tx := f.TestThen()
	return []entity.User{
		tx.User,
	}, nil
}

func (f *userFacade) GetUserById(id int) (*entity.User, error) {
	u := f.repo.FindById(id)
	return u.Right, nil
}

func (f *userFacade) RemoveUserById(id int) core.Either[entity.User, error] {
	user := entity.User{}
	return core.RightEither[entity.User, error](user)
}

func (f *userFacade) RemoveUserById2(id int) *entity.User {
	return either.Bind(
		either.Bind(
			doSome1(id).Then(doSome4).Then(doSome5),
			doSome3),
		doSome2).Right
}

func newTx() core.Either[entity.Txn, error] {
	return core.RightEither[entity.Txn, error](entity.Txn{})
}

func setUser(tx entity.Txn) core.Either[entity.Txn, error] {
	tx.User = entity.User{}
	return core.RightEither[entity.Txn, error](tx)
}

func setUserGroup(tx entity.Txn) core.Either[entity.Txn, error] {
	tx.UserGroup = entity.UserGroup{Name: "ArmGroup"}
	return core.RightEither[entity.Txn, error](tx)
}

func (f *userFacade) TestThen() *entity.Txn {
	tx := newTx().
		Then(setUser).
		Then(setUserGroup)
	return tx.Right
}

func (f *userFacade) RemoveUserById_After(id int) core.Either[entity.User, core.ErrContext] {
	s1 := doSome1(id)
	s2 := either.Bind(s1, doSome4)
	s3 := either.Bind(s2, doSome5)
	s4 := either.Bind(s3, doSome3)
	result := either.Bind(s4, doSome2)
	return result.BindErrContext(handleUserError)
}

func (f *userFacade) RemoveUserById_After2(id int) core.Either[entity.User, core.ErrContext] {
	r1 := doSome1(id).
		Then(doSome4).
		Then(doSome5)
	r2 := either.Bind(r1, doSome3)
	r3 := either.Bind(r2, doSome2)
	return r3.BindErrContext(handleUserError)
}

func handleUserError(input error) core.Either[entity.User, core.ErrContext] {
	return core.LeftEither[entity.User, core.ErrContext](core.NewErrorWithCode(core.UnHandleError, input))
}

func (f *userFacade) RemoveUserByIdFinal(id int) *entity.User {
	return either.Bind(either.Bind(
		doSome1(id).Then(doSome4).Then(doSome5), doSome3), doSome2).Right
}

func (f *userFacade) RemoveUserByIdFinal22(id int) *entity.User {
	b1 := doSome1(id).
		Then(doSome4).
		Then(doSome5)
	b2 := either.Bind(b1, doSome3)
	return either.Bind(b2, doSome2).Right
}

func (f *userFacade) TestValidate(u *entity.User) *core.ErrContext {
	result := vStep1(u).DoNext(vStep2).DoNext(f.vStep3)
	return result.Left
}

func (f *userFacade) TestValidate2(u *entity.User) *core.ErrContext {
	result := validateUserStep1(u).
		DoNext(validateUserStep2).
		DoNext(f.validateUserStep3)
	return result.Left
}

func validateUserStep1(u *entity.User) core.Either[entity.User, core.ErrContext] {
	if u.Id >= 0 {
		return core.NewEither(u, core.NewInvalidateError("ID", core.ValueNotInScope))
	}
	return core.RightEither[entity.User, core.ErrContext](*u)
}

func validateUserStep2(u *entity.User, err *core.ErrContext) core.Either[entity.User, core.ErrContext] {
	if u.Name != "" {
		return core.NewEither(u, core.NewInvalidateError("Name", core.ValueIsRequired).AppendExt(err))
	}
	return core.NewEither(u, err)
}

func (f *userFacade) validateUserStep3(u *entity.User, err *core.ErrContext) core.Either[entity.User, core.ErrContext] {
	if u.Email == "" {
		return core.NewEither(u, core.NewInvalidateError("Email", core.ValueInvalidFormat).AppendExt(err))
	}
	return core.NewEither(u, err)
}

func vStep1(u *entity.User) core.Either[entity.User, core.ErrContext] {
	return core.Validate(u, nil, func(u *entity.User) bool {
		return u.Id > 0
	}, core.NewInvalidateError("ID", core.ValueNotInScope))
}

func vStep2(u *entity.User, err *core.ErrContext) core.Either[entity.User, core.ErrContext] {
	return core.Validate(u, err, func(u *entity.User) bool {
		return u.Name == ""
	}, core.NewInvalidateError("Name", core.ValueIsRequired))
}

func (f *userFacade) vStep3(u *entity.User, err *core.ErrContext) core.Either[entity.User, core.ErrContext] {
	return core.Validate(u, err, func(u *entity.User) bool {
		return u.Email != ""
	}, core.NewInvalidateError("Email", core.ValueInvalidFormat))
}

func doSome1(id int) core.Either[entity.User, error] {
	user := entity.User{}
	return core.RightEither[entity.User, error](user)
}

func doSome2(u entity.UserGroup) core.Either[entity.User, error] {
	return core.RightEither[entity.User, error](entity.User{})
}

func doSome3(u entity.User) core.Either[entity.UserGroup, error] {
	return core.RightEither[entity.UserGroup, error](entity.UserGroup{})
}

func doSome4(u entity.User) core.Either[entity.User, error] {
	return core.RightEither[entity.User, error](u)
}

func doSome5(u entity.User) core.Either[entity.User, error] {
	return core.RightEither[entity.User, error](u)
}
