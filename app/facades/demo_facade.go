package facades

import (
	"fmt"
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/core/either"
	"n4a3/clean-architecture/app/core/global"
	"n4a3/clean-architecture/app/core/util/json"
	stringutil "n4a3/clean-architecture/app/core/util/string"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/integrates/services"
)

type DemoFacade interface {
	Validate(u *entity.User) core.Either[entity.User, core.ErrContext]
	RequestHttp() core.Either[core.Unit, core.ErrContext]
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

func (f *demoFacade) Validate(u *entity.User) core.Either[entity.User, core.ErrContext] {
	return core.CheckNull[entity.User](u).
		Then(checkUsername).
		DoNext(checkName).
		DoNext(checkEmail)
}

func (f *demoFacade) RequestHttp() core.Either[core.Unit, core.ErrContext] {
	httpRequest := f.httpService.GetHttpRequest(f.config.Service.PgwUrl, nil)
	result := either.Bind(httpRequest, json.Unmarshal[[]any])
	fmt.Println(result)
	return core.RightEither[core.Unit, core.ErrContext](core.Unit{})
}

func checkUsername(u entity.User) core.Either[entity.User, core.ErrContext] {
	if stringutil.IsNullOrEmpty(u.Username) {
		return core.NewEither(&u, core.NewInvalidateError("Username", core.ValueNotInScope))
	}
	return core.RightEither[entity.User, core.ErrContext](u)
}

func checkName(u *entity.User, err *core.ErrContext) core.Either[entity.User, core.ErrContext] {
	if stringutil.IsNullOrEmpty(u.Name) {
		return core.NewEither(u, core.NewInvalidateError("Name", core.ValueIsRequired).AppendExt(err))
	}
	return core.NewEither(u, err)
}

func checkEmail(u *entity.User, err *core.ErrContext) core.Either[entity.User, core.ErrContext] {
	if stringutil.IsNullOrEmpty(u.Email) {
		return core.NewEither(u, core.NewInvalidateError("Email", core.ValueInvalidFormat).AppendExt(err))
	}
	return core.NewEither(u, err)
}
