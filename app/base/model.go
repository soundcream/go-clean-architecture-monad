package base

type KeyValue[T, S any] struct {
	Key T
	Val S
}

func NewKeyValue[T, S any](key T, value S) KeyValue[T, S] {
	return KeyValue[T, S]{Key: key, Val: value}
}
