package collection

import "n4a3/clean-architecture/app/core/generic"

type Collection[T any] []T

type Mapping[T, S any] struct {
	Source T
	Dest   S
	Collection[T]
}

func NewMapping[T, S any](c []T) Mapping[T, S] {
	return Mapping[T, S]{Collection: c}
}

func (m Mapping[T, S]) Map(transform func(T) S) Collection[S] {
	return Map(m.Collection, transform)
}

func Map[T any, S any](c Collection[T], transform func(T) S) Collection[S] {
	var result []S
	for i := range c {
		result = append(result, generic.Map(c[i], transform))
	}
	return result
}
