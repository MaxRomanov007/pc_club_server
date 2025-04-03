package redisEx

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"pc_club_server/internal/lib/redis/redisKey"
	reflect2 "pc_club_server/internal/lib/reflect"
	"reflect"
	"time"
)

func (c *Client) SetWithTTL(ctx context.Context, ttl time.Duration, value any, tags ...any) error {
	st := settable{
		Client: c,
		ctx:    ctx,
		key:    redisKey.Key(value, tags...),
		id:     1,
	}

	_ = c.Del(ctx, value, tags...)

	_, err := c.cl.Pipelined(ctx, func(cl redis.Pipeliner) error {
		st.cl = cl
		errs := st.set(reflect.ValueOf(value))
		if len(errs) > 0 {
			return &Errors{
				Errors: errs,
			}
		}
		for i := uint(1); i <= st.id; i++ {
			cl.Expire(ctx, fmt.Sprintf("%s.%d", st.key, i), ttl)
		}

		return nil
	})
	return err
}

type settable struct {
	*Client
	cl  redis.Pipeliner
	tx  *redis.Tx
	ctx context.Context
	key string
	id  uint
}

func (s *settable) set(val reflect.Value) (errs []*Error) {
	key := s.Key()
	for val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
		val = val.Elem()
	}
	switch val.Kind() {
	case reflect.Struct:
		if val.NumField() == 0 {
			break
		}
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			for field.Kind() == reflect.Ptr || field.Kind() == reflect.Interface {
				field = val.Elem()
			}

			if s.IsPrimal(field.Type()) {
				value, isSet, err := s.Translate(field)
				if err != nil {
					errs = append(errs, &Error{
						Code:    ErrTranslationErrorCode,
						Message: "failed to translate primal struct field",
						Field:   val.Type().Field(i).Name,
						Err:     err,
					})
					continue
				}
				if !isSet || reflect.ValueOf(value).IsZero() {
					continue
				}
				s.cl.HSet(s.ctx, key, val.Type().Field(i).Name, value)
				continue
			}

			if field.IsZero() {
				continue
			}
			s.id++
			s.cl.HSet(s.ctx, key, val.Type().Field(i).Name, s.Key())
			if recErrs := s.set(field); len(recErrs) > 0 {
				errs = append(errs, WithTraceAll(recErrs, val.Type().Field(i).Name)...)
				continue
			}
		}
	case reflect.Array, reflect.Slice:
		if val.Len() == 0 {
			break
		}
		for i := 0; i < val.Len(); i++ {
			el := val.Index(i)
			for el.Kind() == reflect.Ptr || el.Kind() == reflect.Interface {
				el = val.Elem()
			}

			if s.IsPrimal(el.Type()) {
				value, isSet, err := s.Translate(el)
				if err != nil {
					errs = append(errs, &Error{
						Code:    ErrTranslationErrorCode,
						Message: "failed to translate primal struct field",
						Field:   val.Type().Field(i).Name,
						Err:     err,
					})
					continue
				}
				if !isSet || reflect.ValueOf(value).IsZero() {
					continue
				}
				s.cl.RPush(s.ctx, key, value)
				continue
			}

			s.id++
			s.cl.RPush(s.ctx, key, s.Key())
			if recErrs := s.set(el); len(recErrs) > 0 {
				errs = append(errs, WithTraceAll(recErrs, fmt.Sprintf("[%d]", i))...)
				continue
			}
		}
	case reflect.Map:
		if val.Len() == 0 {
			break
		}
		for _, k := range val.MapKeys() {
			el := val.MapIndex(k)
			for el.Kind() == reflect.Ptr || el.Kind() == reflect.Interface {
				el = val.Elem()
			}

			mapKey, _, err := s.Translate(k)
			if err != nil {
				errs = append(errs, &Error{
					Code:    ErrTranslationErrorCode,
					Message: "failed to translate map key element",
					Field:   "[" + k.String() + "]",
					Err:     err,
				})
			}

			if s.IsPrimal(el.Type()) {
				value, isSet, err := s.Translate(el)
				if err != nil {
					errs = append(errs, &Error{
						Code:    ErrTranslationErrorCode,
						Message: "failed to translate primal struct field",
						Field:   "[" + k.String() + "]",
						Err:     err,
					})
					continue
				}
				if !isSet || reflect.ValueOf(value).IsZero() {
					continue
				}
				s.cl.HSet(s.ctx, key, mapKey, value)
				continue
			}

			s.id++
			s.cl.HSet(s.ctx, fmt.Sprintf("%v", mapKey), k.Interface(), s.Key())
			if recErrs := s.set(el); len(recErrs) > 0 {
				errs = append(errs, WithTraceAll(recErrs, "["+k.String()+"]")...)
				continue
			}
		}
	default:
		errs = append(errs, &Error{
			Code:    ErrUnsupportedTypeCode,
			Message: fmt.Sprintf("unsupported type: %s (kind: %s)", val.Type().String(), val.Kind().String()),
		})
	}
	return
}

func (s *settable) Key() string {
	return fmt.Sprintf("%s.%d", s.key, s.id)
}

func (s *settable) IsPrimal(t reflect.Type) bool {
	if reflect2.IsPrimary(t) {
		return true
	}
	if _, ok := s.TranslateSetMap[t.String()]; ok {
		return true
	}
	return false
}

func (s *settable) Translate(value reflect.Value) (any, bool, error) {
	if fn, ok := s.TranslateSetMap[value.Type().String()]; ok {
		return fn(value)
	}
	if !reflect2.IsPrimary(value.Type()) {
		return nil, false, fmt.Errorf("type to translate is not primary")
	}
	return value.Interface(), true, nil
}
