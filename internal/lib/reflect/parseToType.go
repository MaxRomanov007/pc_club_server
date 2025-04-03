package reflect

import (
	"fmt"
	"reflect"
	"strconv"
)

func ParseToType(s string, t reflect.Type) (any, error) {
	var res any
	var err error
	switch t.Kind() {
	case reflect.String:
		res = s
	case reflect.Int:
		res, err = strconv.Atoi(s)
	case reflect.Int8:
		var val int64
		val, err = strconv.ParseInt(s, 10, 8)
		res = int8(val)
	case reflect.Int16:
		var val int64
		val, err = strconv.ParseInt(s, 10, 16)
		res = int16(val)
	case reflect.Int32:
		var val int64
		val, err = strconv.ParseInt(s, 10, 32)
		res = int32(val)
	case reflect.Int64:
		res, err = strconv.ParseInt(s, 10, 64)
		res = res.(int64)
	case reflect.Uint:
		var val uint64
		val, err = strconv.ParseUint(s, 10, 0)
		res = uint(val)
	case reflect.Uint8:
		var val uint64
		val, err = strconv.ParseUint(s, 10, 8)
		res = uint8(val)
	case reflect.Uint16:
		var val uint64
		val, err = strconv.ParseUint(s, 10, 16)
		res = uint16(val)
	case reflect.Uint32:
		var val uint64
		val, err = strconv.ParseUint(s, 10, 32)
		res = uint32(val)
	case reflect.Uint64:
		res, err = strconv.ParseUint(s, 10, 64)
	case reflect.Float32:
		var val float64
		val, err = strconv.ParseFloat(s, 32)
		res = float32(val)
	case reflect.Float64:
		res, err = strconv.ParseFloat(s, 64)
	case reflect.Bool:
		res, err = strconv.ParseBool(s)
	default:
		err = fmt.Errorf("%w: unsupported type %s", ErrUnsupportedType, t.Name())
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}
