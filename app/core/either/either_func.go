package either

import "n4a3/clean-architecture/app/core"

func Bind[R, R2, L any](e core.Either[R, L], fun func(right R) core.Either[R2, L]) core.Either[R2, L] {
	if e.IsLeft() {
		return core.Either[R2, L]{Left: e.Left}
	}
	return fun(*e.Right)
}

func Map[R, R2, L any](e core.Either[R, L], fun func(left *L, right *R) core.Either[R2, L]) core.Either[R2, L] {
	if e.IsLeft() {
		return core.Either[R2, L]{Left: e.Left}
	}
	return fun(e.Left, e.Right)
}
