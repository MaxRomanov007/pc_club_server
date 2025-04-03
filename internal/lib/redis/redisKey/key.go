package redisKey

import (
	"fmt"
	"reflect"
	"strings"
)

type Settable interface {
	RedisKey() string
}

type TableSettable interface {
	TableName() string
}

func Key(value any, tags ...any) string {
	return KeyFor(reflect.TypeOf(value), tags...)
}

func KeyFor(t reflect.Type, tags ...any) string {
	for t.Kind() == reflect.Array ||
		t.Kind() == reflect.Slice ||
		t.Kind() == reflect.Map ||
		t.Kind() == reflect.Pointer {

		t = t.Elem()
	}

	t = reflect.PointerTo(t)
	reflect.Zero(t)
	var key string
	switch {
	case t.Implements(reflect.TypeFor[Settable]()):
		key = reflect.Zero(t).MethodByName("RedisKey").Call(nil)[0].String()
	case t.Implements(reflect.TypeFor[TableSettable]()):
		key = reflect.Zero(t).MethodByName("TableName").Call(nil)[0].String()
	default:
		key = t.Elem().Name()
	}
	if len(tags) > 0 {
		tagsStr := fmt.Sprintf("%v", tags)
		tagsStr = tagsStr[1 : len(tagsStr)-1]
		tagsStr = strings.ReplaceAll(tagsStr, " ", "-")
		key = key + ":" + tagsStr
	}
	return key
}
