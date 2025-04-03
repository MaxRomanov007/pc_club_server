package redisEx

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"pc_club_server/internal/lib/redis/redisKey"
	reflect2 "pc_club_server/internal/lib/reflect"
	"reflect"
)

func (c *Client) Get(ctx context.Context, value any, tags ...any) error {
	rec := &received{
		Client: c,
		ctx:    ctx,
	}

	keys, err := c.Keys(ctx, redisKey.Key(value, tags...)+".*")
	if err != nil {
		return err
	}
	mapCmd := make([]*redis.MapStringStringCmd, len(keys))
	sliceCmd := make([]*redis.StringSliceCmd, len(keys))
	_, err = c.cl.Pipelined(ctx, func(cl redis.Pipeliner) error {
		for i, key := range keys {
			mapCmd[i] = cl.HGetAll(ctx, key)
			sliceCmd[i] = cl.LRange(ctx, key, 0, -1)
		}
		return nil
	})

	res := make(map[string]any)
	for i, key := range keys {
		var val any
		val, err = redisValue[map[string]string](mapCmd[i], key)
		if err != nil {
			return err
		} else if len(val.(map[string]string)) > 0 {
			res[key] = val
			continue
		}
		val, err = redisValue[[]string](sliceCmd[i], key)
		if err != nil {
			return err
		} else if len(val.([]string)) > 0 {
			res[key] = val
			continue
		}
		return &Errors{
			Errors: []*Error{{
				Code:    ErrUnexpectedRedisType,
				Message: "unexpected redis type on key: " + key,
				Err:     err,
			}},
		}
	}
	rec.values = res

	key := redisKey.Key(value, tags...) + ".1"
	if errs := rec.get(reflect.ValueOf(value), key); len(errs) > 0 {
		if len(errs) == 1 && errs[0].Code == ErrNilResultCode {
			return ErrNil
		}
		return &Errors{
			Errors: errs,
		}
	}
	return nil
}

type received struct {
	*Client
	ctx    context.Context
	values map[string]any
}

func (r *received) get(dst reflect.Value, value string) []*Error {
	if dst.Kind() != reflect.Ptr {
		return []*Error{ErrNotPointer}
	}
	dst = dst.Elem()

	if r.IsPrimal(dst.Type()) {
		val, err := r.Parse(value, dst.Type())
		if err != nil {
			return []*Error{{
				Code:    ErrTranslationErrorCode,
				Message: "failed to translate primal field",
				Err:     err,
			}}
		}
		dst.Set(reflect.ValueOf(val))
		return nil
	}

	var errs []*Error
	switch dst.Kind() {
	case reflect.Struct:
		res, ok := r.values[value].(map[string]string)
		if !ok {
			errs = append(errs, &Error{
				Code:    ErrObjectMismatchCode,
				Message: "expected: map[string]string got: " + reflect.TypeOf(res).String() + " (key: \"" + value + "\")",
			})
			break
		}
		for k, v := range res {
			field := dst.FieldByName(k)
			if !field.IsValid() {
				errs = append(errs, &Error{
					Code:    ErrFieldNotExistsCode,
					Message: "field \"" + k + "\" doesnt exists in struct",
				})
				continue
			}
			if recErrs := r.get(dst.FieldByName(k).Addr(), v); recErrs != nil {
				errs = append(errs, WithTraceAll(recErrs, k)...)
			}
		}
	case reflect.Slice:
		res, ok := r.values[value].([]string)
		if !ok {
			errs = append(errs, &Error{
				Code:    ErrObjectMismatchCode,
				Message: "expected: []string got: " + reflect.TypeOf(res).String() + " (key: \"" + value + "\")",
			})
			break
		}
		slice := reflect.MakeSlice(dst.Type(), len(res), len(res))
		for i, v := range res {
			el := slice.Index(i)
			if resErrs := r.get(el.Addr(), v); resErrs != nil {
				errs = append(errs, WithTraceAll(resErrs, fmt.Sprintf("[%d]", i))...)
			}
		}
		dst.Set(slice)
	case reflect.Array:
		res, ok := r.values[value].([]string)
		if !ok {
			errs = append(errs, &Error{
				Code:    ErrObjectMismatchCode,
				Message: "expected: []string got: " + reflect.TypeOf(res).String() + " (key: \"" + value + "\")",
			})
			break
		}
		arr := reflect.New(reflect.ArrayOf(dst.Type().Len(), dst.Type().Elem())).Elem()
		for i, v := range res {
			el := arr.Index(i)
			if resErrs := r.get(el.Addr(), v); resErrs != nil {
				errs = append(errs, WithTraceAll(resErrs, fmt.Sprintf("[%d]", i))...)
			}
		}
		dst.Set(arr)
	case reflect.Map:
		res, ok := r.values[value].(map[string]string)
		if !ok {
			errs = append(errs, &Error{
				Code:    ErrObjectMismatchCode,
				Message: "expected: map[string]string got: " + reflect.TypeOf(res).String() + " (key: \"" + value + "\")",
			})
			break
		}
		newMap := reflect.MakeMap(dst.Type())
		for k, v := range res {
			var mapKey any
			mapKey, err := r.Parse(k, dst.Type().Key())
			if err != nil {
				errs = append(errs, &Error{
					Code:    ErrTranslationErrorCode,
					Message: "failed to parse map key (type: " + dst.Type().Key().String() + ")",
					Err:     err,
					Field:   "[" + k + "]",
				})
				continue
			}
			mapValue := reflect.New(dst.Type().Elem()).Elem()
			if resErrs := r.get(mapValue.Addr(), v); resErrs != nil {
				errs = append(errs, WithTraceAll(resErrs, "["+k+"]")...)
				continue
			}
			newMap.SetMapIndex(reflect.ValueOf(mapKey), mapValue)
		}
		dst.Set(newMap)
	default:
		errs = append(errs, &Error{
			Code:    ErrUnsupportedTypeCode,
			Message: fmt.Sprintf("unsupported type: %s (kind: %s)", dst.Type().String(), dst.Kind().String()),
		})
	}
	return errs
}

func (r *received) Parse(str string, t reflect.Type) (any, error) {
	if fn, ok := r.TranslateGetMap[t.String()]; ok {
		return fn(str)
	}
	return reflect2.ParseToType(str, t)
}

func (r *received) IsPrimal(t reflect.Type) bool {
	if reflect2.IsPrimary(t) {
		return true
	}
	if _, ok := r.TranslateGetMap[t.String()]; ok {
		return true
	}
	return false
}

type redisCmd[T map[string]string | []string] interface {
	Val() T
	Err() error
}

func redisValue[T map[string]string | []string](cmd redisCmd[T], key string) (T, error) {
	err := cmd.Err()
	if err == nil {
		val := cmd.Val()
		if len(val) == 0 {
			return nil, &Errors{
				Errors: []*Error{{
					Code:    ErrRedisErrorCode,
					Message: "empty value",
					Err:     redis.Nil,
				}},
			}
		}
		return val, nil
	} else if !redis.HasErrorPrefix(err, "WRONGTYPE") {
		return nil, &Errors{
			Errors: []*Error{{
				Code:    ErrRedisErrorCode,
				Message: "failed to get key: " + key,
				Err:     err,
			}},
		}
	}
	return nil, nil
}
