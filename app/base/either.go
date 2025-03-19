package base

type Either[R, L any] struct {
	Right *R
	Left  *L
}

type Unit struct {
}

func NewUnit() *Unit {
	return &Unit{}
}

func (e Either[R, L]) IsRight() bool {
	return e.Right != nil && !e.IsLeft()
}

func (e Either[R, L]) IsLeft() bool {
	return e.Left != nil
}

// Then chains another if Right. And change Right (side effect)
func (e Either[R, L]) Then(fun func(R) Either[R, L]) Either[R, L] {
	if e.IsLeft() {
		return LeftEither[R, L](*e.Left)
	}
	return fun(*e.Right)
}

// ThenPtr chains another if Right. And change Right (side effect)
func (e Either[R, L]) ThenPtr(fun func(*R) Either[R, L]) Either[R, L] {
	if e.IsLeft() {
		return LeftEither[R, L](*e.Left)
	}
	return fun(e.Right)
}

// Next chains another if Right. And not change Right
func (e Either[R, L]) Next(fun func(R) Either[R, L]) Either[R, L] {
	if e.IsLeft() || e.Right == nil {
		return e
	}
	fun(*e.Right)
	return e
}

// NextPtr chains another if Right. And not change Right
func (e Either[R, L]) NextPtr(fun func(*R) Either[R, L]) Either[R, L] {
	if e.IsLeft() || e.Right == nil {
		return e
	}
	fun(e.Right)
	return e
}

// BindErrContext Bind Left to base.ErrContext
func (e Either[R, L]) BindErrContext(fun func(L) Either[R, ErrContext]) Either[R, ErrContext] {
	if e.IsLeft() {
		return fun(*e.Left)
	}
	return RightEither[R, ErrContext](*e.Right)
}

func (c Either[R, L]) DoNext(fun func(*R, *L) Either[R, ErrContext]) Either[R, ErrContext] {
	e := fun(c.Right, c.Left)
	return e
}

func NewEither[R any, L any](right *R, left *L) Either[R, L] {
	return Either[R, L]{Right: right, Left: left}
}

func NewRightEither[R, L any](right *R) Either[R, L] {
	return Either[R, L]{Right: right}
}

func LeftEither[R, L any](left L) Either[R, L] {
	return Either[R, L]{Left: &left}
}

func RightEither[R, L any](right R) Either[R, L] {
	return Either[R, L]{Right: &right}
}

func Validate[T any](input *T, err *ErrContext, fn func(*T) bool, error *ErrContext) Either[T, ErrContext] {
	if !fn(input) {
		return NewEither[T, ErrContext](input, error.AppendExt(err))
	}
	return NewEither[T, ErrContext](input, err)
}

func CheckNull[T any](model *T) Either[T, ErrContext] {
	if model == nil {
		return LeftEither[T, ErrContext](NewErrorCode(BadRequest))
	}
	return RightEither[T, ErrContext](*model)
}
