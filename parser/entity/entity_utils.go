package entity

import (
	"reflect"
	"time"
)

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return true
		}

		return false
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}

		return z
	case reflect.Struct:
		if v.Type() == reflect.TypeOf(t) {
			if v.Interface().(time.Time).IsZero() {
				return true
			}
			return false
		}

		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}

		return z
	}

	// compare other types directly
	z := reflect.Zero(v.Type())

	return v.Interface() == z.Interface()
}
