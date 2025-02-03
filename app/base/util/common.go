package util

func ToPtr[T any](value T) *T {
	return &value
}

func Map(k string, in ...interface{}) map[string][]interface{} {
	return map[string][]interface{}{
		k: in,
	}
}
