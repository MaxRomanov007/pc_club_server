package reflect

import "reflect"

func IsPrimary(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Struct,
		reflect.Map, reflect.Slice, reflect.Array, reflect.Func,
		reflect.Chan, reflect.UnsafePointer, reflect.Invalid:
		return false
	default:
		return true
	}
}
