package stringutil

import (
	"fmt"
	"n4a3/clean-architecture/app/base"
	"strings"
)

func IsNullOrEmpty(str string) bool {
	return len(str) == 0 || strings.Trim(str, " ") == ""
}

func ToIntAndError(str string) (int, error) {
	var i int
	if _, err := fmt.Sscan(str, &i); err != nil {
		return i, err
	}
	return i, nil
}

func ToInt(str string) *int {
	if i, err := ToIntAndError(str); err != nil {
		return &i
	}
	return nil
}

func ToIntWithDefault(str string, def int) int {
	var i = ToInt(str)
	if i != nil {
		return *i
	}
	return def
}

func ToIntEither(str string) base.Either[int, base.ErrContext] {
	i, err := ToIntAndError(str)
	if err != nil {
		return base.LeftEither[int, base.ErrContext](base.NewErrorWithCode(base.Conflict, err))
	}
	return base.NewRightEither[int, base.ErrContext](&i)
}
