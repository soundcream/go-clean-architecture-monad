package base

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
)

func JsonUnmarshal[T any](b []byte) Either[T, ErrContext] {
	result := new(T)
	if b != nil {
		if err := json.Unmarshal(b, &result); err != nil {
			log.Error("Error cannot Unmarshalling response")
			return LeftEither[T, ErrContext](NewErrorWithCode(Invalid, err))
		}
		return RightEither[T, ErrContext](*result)
	}
	return RightEither[T, ErrContext](*result)
}
