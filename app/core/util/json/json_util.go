package json

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"n4a3/clean-architecture/app/core"
)

func Unmarshal[T any](b []byte) core.Either[T, core.ErrContext] {
	result := new(T)
	if b != nil {
		if err := json.Unmarshal(b, &result); err != nil {
			log.Error("Error cannot Unmarshalling", err)
			return core.LeftEither[T, core.ErrContext](core.NewErrorWithCode(core.Invalid, err))
		}
		return core.RightEither[T, core.ErrContext](*result)
	}
	return core.RightEither[T, core.ErrContext](*result)
}

func Parse[T any](value string) core.Either[T, core.ErrContext] {
	result := new(T)
	if err := json.Unmarshal([]byte(value), &result); err != nil {
		log.Error("Error cannot Unmarshalling", err)
		return core.LeftEither[T, core.ErrContext](core.NewErrorWithCode(core.Invalid, err))
	}
	return core.RightEither[T, core.ErrContext](*result)
}

func Serialize(value any) core.Either[string, core.ErrContext] {
	bytes, err := json.Marshal(value)
	if err != nil {
		log.Error("Error cannot Marshal", err)
		return core.LeftEither[string, core.ErrContext](core.NewErrorWithCode(core.Invalid, err))
	}
	return core.RightEither[string, core.ErrContext](string(bytes))

}
