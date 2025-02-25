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

// MapValue Auto mapper source to dest
func MapValue[T any, S any](source *T, dest *S) {
	MapValueOf(reflect.ValueOf(source), reflect.ValueOf(dest))
}

// MapValueOf Auto mapper source to dest
func MapValueOf(source reflect.Value, dest reflect.Value) {
	if source.Kind() == reflect.Ptr {
		source = source.Elem()
	}
	if dest.Kind() == reflect.Ptr {
		dest = dest.Elem()
	}
	for i := 0; i < dest.NumField(); i++ {
		fieldType := dest.Type().Field(i)
		field := dest.Field(i)
		if fieldValue := source.FieldByName(fieldType.Name); true {
			val := getValue(fieldValue)
			fieldSet := getField(field, fieldValue)
			if fieldSet.Kind() == reflect.Struct {
				if fieldSet.Type().Name() == val.Type().Name() && fieldSet.CanSet() {
					fieldSet.Set(reflect.ValueOf(val.Interface()))
				}
			} else if fieldSet.CanSet() {
				fieldSet.Set(reflect.ValueOf(val.Interface()))
			}
		}
	}
}

func getValue(fieldValue reflect.Value) reflect.Value {
	val := fieldValue
	if fieldValue.Kind() == reflect.Ptr {
		val = fieldValue.Elem()
	}
	return val
}

func getField(field reflect.Value, fieldValue reflect.Value) reflect.Value {
	fieldSet := field
	if field.CanSet() && field.Kind() == reflect.Ptr {
		if field.IsNil() && !fieldValue.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		fieldSet = field.Elem()
	}
	return fieldSet
}
