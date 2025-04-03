package redisEx

import (
	"github.com/redis/go-redis/v9"
	"strings"
)

type Error struct {
	Code    string
	Message string
	Field   string
	Err     error
}

func (e *Error) Error() string {
	if e.Message == "" {
		return e.Field + ": " + e.Code
	}
	return e.Field + ": " + e.Code + ": " + e.Message
}

type Errors struct {
	Errors []*Error
}

func (e *Errors) Error() string {
	errs := make([]string, len(e.Errors))
	for i, err := range e.Errors {
		errs[i] = err.Error()
	}
	return strings.Join(errs, "; ")
}

const (
	ErrUnsupportedTypeCode  = "UnsupportedType"
	ErrTranslationErrorCode = "TranslationError"
	ErrRedisErrorCode       = "RedisError"
	ErrNotPointerCode       = "ValueNotPointer"
	ErrFieldNotExistsCode   = "FieldNotExists"
	ErrNilResultCode        = "NilResult"
	ErrUnexpectedRedisType  = "UnexpectedRedisType"
	ErrObjectMismatchCode   = "ObjectMismatch"
)

var (
	ErrNil = &Error{
		Code:    ErrNilResultCode,
		Message: "nil result",
		Err:     redis.Nil,
	}
	ErrNotPointer = &Error{
		Code:    ErrNotPointerCode,
		Message: "value isn't a pointer",
	}
)

func WithTrace(err *Error, trace string) *Error {
	if err.Field == "" {
		err.Field = trace
	} else {
		err.Field = trace + "." + err.Field
	}
	return err
}

func WithTraceAll(errs []*Error, trace string) []*Error {
	for i := 0; i < len(errs); i++ {
		errs[i] = WithTrace(errs[i], trace)
	}
	return errs
}
