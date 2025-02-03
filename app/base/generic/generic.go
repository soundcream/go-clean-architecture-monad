package generic

import "reflect"

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

func GetTagByName[T any](tag string) string {
	var result = ""
	var allFields = GetAllFields[T]()
	for i := range allFields {
		var field = allFields[i]
		var tagValue = field.Tag.Get(tag)
		if tagValue != "" {
			return tagValue
		}
	}
	return result
}

func GetAllFields[T any]() []reflect.StructField {
	var result []reflect.StructField
	var model T
	typ := reflect.TypeOf(model)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		result = append(result, typ.Field(i))
	}
	return result
}
