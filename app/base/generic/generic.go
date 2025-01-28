package generic

type Mapping[T, S any] struct {
	Source T
	Dest   S
}

func NewMapping[T, S any](source T) Mapping[T, S] {
	return Mapping[T, S]{Source: source}
}

func (m Mapping[T, S]) Map(transform func(T) S) S {
	return Map(m.Source, transform)
}

func Map[T any, S any](c T, transform func(T) S) S {
	return transform(c)
}
