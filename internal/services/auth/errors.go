package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"pc_club_server/internal/storage/mssql"
)

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

const (
	ErrTokenMalformedCode        = "TokenMalformed"
	ErrTokenSignatureInvalidCode = "TokenSignatureInvalid"
	ErrTokenExpiredCode          = "TokenExpired"
	ErrTokenInBlackListCode      = "TokenInBlackList"
	ErrUserNotFoundCode          = "UserNotFound"
	ErrInvalidRefreshVersionCode = "InvalidRefreshVersion"
)

var (
	ErrTokenMalformed = &Error{
		Code:    ErrTokenMalformedCode,
		Message: "token is malformed",
	}
	ErrTokenSignatureInvalid = &Error{
		Code:    ErrTokenSignatureInvalidCode,
		Message: "token signature is invalid",
	}
	ErrTokenExpired = &Error{
		Code:    ErrTokenExpiredCode,
		Message: "token is expired",
	}
	ErrTokenInBlackList = &Error{
		Code:    ErrTokenInBlackListCode,
		Message: "token is in blacklist",
	}
	ErrUserNotFound = &Error{
		Code:    ErrUserNotFoundCode,
		Message: "user not found",
	}
	ErrInvalidRefreshVersion = &Error{
		Code:    ErrInvalidRefreshVersionCode,
		Message: "invalid refresh version",
	}
)

func TokenError(err error) error {
	switch {
	case errors.Is(err, jwt.ErrTokenMalformed):
		err = ErrTokenMalformed
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		err = ErrTokenSignatureInvalid
	case errors.Is(err, jwt.ErrTokenExpired), errors.Is(err, jwt.ErrTokenNotValidYet):
		err = ErrTokenExpired
	}

	return err
}

func HandleMssqlError(err error) error {
	var mssqlError *mssql.Error
	if !errors.As(err, &mssqlError) {
		return fmt.Errorf("unknown error: %w", err)
	}
	switch mssqlError.Code {
	case mssql.ErrNotFoundCode:
		err = ErrUserNotFound
	default:
		err = fmt.Errorf("unknown mssql error: %w", mssqlError)
	}
	return err
}
