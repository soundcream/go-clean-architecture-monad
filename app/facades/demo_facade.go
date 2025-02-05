package facades

import (
	"n4a3/clean-architecture/app/base"
	stringutil "n4a3/clean-architecture/app/base/util/string"
	"n4a3/clean-architecture/app/domain/entity"
)

type DemoFacade interface {
	Validate(u *entity.User) base.Either[entity.User, base.ErrContext]
}

type demoFacade struct {
}

func NewDemoFacade() DemoFacade {
	return &demoFacade{}
}

func (f *demoFacade) Validate(u *entity.User) base.Either[entity.User, base.ErrContext] {
	return checkNull(u).
		Then(checkUsername).
		DoNext(checkName).
		DoNext(checkEmail)
}

func checkNull[T any](model *T) base.Either[T, base.ErrContext] {
	if model == nil {
		return base.LeftEither[T, base.ErrContext](base.NewErrorCode(base.BadRequest))
	}
	return base.RightEither[T, base.ErrContext](*model)
}

func checkUsername(u entity.User) base.Either[entity.User, base.ErrContext] {
	if stringutil.IsNullOrEmpty(u.Username) {
		return base.NewEither(&u, base.NewInvalidateError("Username", base.ValueNotInScope))
	}
	return base.RightEither[entity.User, base.ErrContext](u)
}

func checkName(u *entity.User, err *base.ErrContext) base.Either[entity.User, base.ErrContext] {
	if stringutil.IsNullOrEmpty(u.Name) {
		return base.NewEither(u, base.NewInvalidateError("Name", base.ValueIsRequired).AppendExt(err))
	}
	return base.NewEither(u, err)
}

func checkEmail(u *entity.User, err *base.ErrContext) base.Either[entity.User, base.ErrContext] {
	if stringutil.IsNullOrEmpty(u.Email) {
		return base.NewEither(u, base.NewInvalidateError("Email", base.ValueInvalidFormat).AppendExt(err))
	}
	return base.NewEither(u, err)
}
