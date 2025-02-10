package either

import "n4a3/clean-architecture/app/base"

func Bind[R, R2, L any](e base.Either[R, L], fun func(right R) base.Either[R2, L]) base.Either[R2, L] {
	if e.IsLeft() {
		return base.Either[R2, L]{Left: e.Left}
	}
	return fun(*e.Right)
}

func Map[R, R2, L any](e base.Either[R, L], fun func(left *L, right *R) base.Either[R2, L]) base.Either[R2, L] {
	if e.IsLeft() {
		return base.Either[R2, L]{Left: e.Left}
	}
	return fun(e.Left, e.Right)
}
