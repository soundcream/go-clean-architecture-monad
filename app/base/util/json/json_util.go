package json

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"n4a3/clean-architecture/app/base"
)

func Unmarshal[T any](b []byte) base.Either[T, base.ErrContext] {
	result := new(T)
	if b != nil {
		if err := json.Unmarshal(b, &result); err != nil {
			log.Error("Error cannot Unmarshalling", err)
			return base.LeftEither[T, base.ErrContext](base.NewErrorWithCode(base.Invalid, err))
		}
		return base.RightEither[T, base.ErrContext](*result)
	}
	return base.RightEither[T, base.ErrContext](*result)
}

func Parse[T any](value string) base.Either[T, base.ErrContext] {
	result := new(T)
	if err := json.Unmarshal([]byte(value), &result); err != nil {
		log.Error("Error cannot Unmarshalling", err)
		return base.LeftEither[T, base.ErrContext](base.NewErrorWithCode(base.Invalid, err))
	}
	return base.RightEither[T, base.ErrContext](*result)
}

func Serialize(value any) base.Either[string, base.ErrContext] {
	bytes, err := json.Marshal(value)
	if err != nil {
		log.Error("Error cannot Marshal", err)
		return base.LeftEither[string, base.ErrContext](base.NewErrorWithCode(base.Invalid, err))
	}
	return base.RightEither[string, base.ErrContext](string(bytes))

}
