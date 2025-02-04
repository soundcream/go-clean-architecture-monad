package util

import "reflect"

func ToPtr[T any](value T) *T {
	return &value
}

func PtrToPtr[S, D any](value *S, fn func(S) D) *D {
	if value == nil {
		return nil
	}
	r := fn(*value)
	return &r
}

func GetFieldName(field interface{}) string {
	typ := reflect.TypeOf(field)
	if typ.Kind() == reflect.Ptr {
		return typ.Elem().Name()
	}
	return typ.Name()
}

func Map(k string, in ...interface{}) *map[string][]interface{} {
	return &map[string][]interface{}{
		k: in,
	}
}
