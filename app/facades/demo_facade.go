package facades

import (
	"fmt"
	"n4a3/clean-architecture/app/base"
	"n4a3/clean-architecture/app/base/either"
	"n4a3/clean-architecture/app/base/global"
	"n4a3/clean-architecture/app/base/util/json"
	stringutil "n4a3/clean-architecture/app/base/util/string"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/integrates/services"
)

type DemoFacade interface {
	Validate(u *entity.User) base.Either[entity.User, base.ErrContext]
	RequestHttp() base.Either[base.Unit, base.ErrContext]
}

type demoFacade struct {
	httpService services.HttpService
	config      global.Config
}

func NewDemoFacade(config global.Config) DemoFacade {
	return &demoFacade{
		config:      config,
		httpService: services.NewHttpService(),
	}
}

func (f *demoFacade) Validate(u *entity.User) base.Either[entity.User, base.ErrContext] {
	return base.CheckNull[entity.User](u).
		Then(checkUsername).
		DoNext(checkName).
		DoNext(checkEmail)
}

func (f *demoFacade) RequestHttp() base.Either[base.Unit, base.ErrContext] {
	httpRequest := f.httpService.GetHttpRequest(f.config.Service.PgwUrl, nil)
	result := either.Bind(httpRequest, json.Unmarshal[[]any])
	fmt.Println(result)
	return base.RightEither[base.Unit, base.ErrContext](base.Unit{})
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
