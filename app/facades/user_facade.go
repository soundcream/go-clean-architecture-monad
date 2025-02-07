package facades

import (
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/either"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/integrates/repository"
)

type UserFacade interface {
	CreateUser(name, email string) (*entity.User, error)
	GetUserById(id int) (*entity.User, error)
	ListUsers() ([]entity.User, error)
	TestThen() *entity.Txn
	TestValidate(u *entity.User) *base.ErrContext
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
	return u, nil
}

func (f *userFacade) RemoveUserById(id int) base.Either[entity.User, error] {
	user := entity.User{}
	return base.RightEither[entity.User, error](user)
}

func (f *userFacade) RemoveUserById2(id int) *entity.User {
	return either.Bind(
		either.Bind(
			doSome1(id).Then(doSome4).Then(doSome5),
			doSome3),
		doSome2).Right
}

func newTx() base.Either[entity.Txn, error] {
	return base.RightEither[entity.Txn, error](entity.Txn{})
}

func setUser(tx entity.Txn) base.Either[entity.Txn, error] {
	tx.User = entity.User{}
	return base.RightEither[entity.Txn, error](tx)
}

func setUserGroup(tx entity.Txn) base.Either[entity.Txn, error] {
	tx.UserGroup = entity.UserGroup{Name: "ArmGroup"}
	return base.RightEither[entity.Txn, error](tx)
}

func (f *userFacade) TestThen() *entity.Txn {
	tx := newTx().
		Then(setUser).
		Then(setUserGroup)
	return tx.Right
}

func (f *userFacade) RemoveUserById_After(id int) base.Either[entity.User, base.ErrContext] {
	s1 := doSome1(id)
	s2 := either.Bind(s1, doSome4)
	s3 := either.Bind(s2, doSome5)
	s4 := either.Bind(s3, doSome3)
	result := either.Bind(s4, doSome2)
	return result.BindErrContext(handleUserError)
}

func (f *userFacade) RemoveUserById_After2(id int) base.Either[entity.User, base.ErrContext] {
	r1 := doSome1(id).
		Then(doSome4).
		Then(doSome5)
	r2 := either.Bind(r1, doSome3)
	r3 := either.Bind(r2, doSome2)
	return r3.BindErrContext(handleUserError)
}

func handleUserError(input error) base.Either[entity.User, base.ErrContext] {
	return base.LeftEither[entity.User, base.ErrContext](base.NewErrorWithCode(base.UnHandleError, input))
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

func (f *userFacade) TestValidate(u *entity.User) *base.ErrContext {
	result := vStep1(u).DoNext(vStep2).DoNext(f.vStep3)
	return result.Left
}

func (f *userFacade) TestValidate2(u *entity.User) *base.ErrContext {
	result := validateUserStep1(u).
		DoNext(validateUserStep2).
		DoNext(f.validateUserStep3)
	return result.Left
}

func validateUserStep1(u *entity.User) base.Either[entity.User, base.ErrContext] {
	if u.Id >= 0 {
		return base.NewEither(u, base.NewInvalidateError("ID", base.ValueNotInScope))
	}
	return base.RightEither[entity.User, base.ErrContext](*u)
}

func validateUserStep2(u *entity.User, err *base.ErrContext) base.Either[entity.User, base.ErrContext] {
	if u.Name != "" {
		return base.NewEither(u, base.NewInvalidateError("Name", base.ValueIsRequired).AppendExt(err))
	}
	return base.NewEither(u, err)
}

func (f *userFacade) validateUserStep3(u *entity.User, err *base.ErrContext) base.Either[entity.User, base.ErrContext] {
	if u.Email == "" {
		return base.NewEither(u, base.NewInvalidateError("Email", base.ValueInvalidFormat).AppendExt(err))
	}
	return base.NewEither(u, err)
}

func vStep1(u *entity.User) base.Either[entity.User, base.ErrContext] {
	return base.Validate(u, nil, func(u *entity.User) bool {
		return u.Id > 0
	}, base.NewInvalidateError("ID", base.ValueNotInScope))
}

func vStep2(u *entity.User, err *base.ErrContext) base.Either[entity.User, base.ErrContext] {
	return base.Validate(u, err, func(u *entity.User) bool {
		return u.Name == ""
	}, base.NewInvalidateError("Name", base.ValueIsRequired))
}

func (f *userFacade) vStep3(u *entity.User, err *base.ErrContext) base.Either[entity.User, base.ErrContext] {
	return base.Validate(u, err, func(u *entity.User) bool {
		return u.Email != ""
	}, base.NewInvalidateError("Email", base.ValueInvalidFormat))
}

func doSome1(id int) base.Either[entity.User, error] {
	user := entity.User{}
	return base.RightEither[entity.User, error](user)
}

func doSome2(u entity.UserGroup) base.Either[entity.User, error] {
	return base.RightEither[entity.User, error](entity.User{})
}

func doSome3(u entity.User) base.Either[entity.UserGroup, error] {
	return base.RightEither[entity.UserGroup, error](entity.UserGroup{})
}

func doSome4(u entity.User) base.Either[entity.User, error] {
	return base.RightEither[entity.User, error](u)
}

func doSome5(u entity.User) base.Either[entity.User, error] {
	return base.RightEither[entity.User, error](u)
}
