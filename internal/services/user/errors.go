package user

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	gorm "pc_club_server/internal/storage/mssql"
)

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

const (
	ErrInvalidCredentialsCode = "InvalidCredentials"
	ErrUserAlreadyExistsCode  = "UserAlreadyExists"
	ErrAccessDeniedCode       = "AccessDenied"
	ErrUserNotFoundCode       = "UserNotFound"
)

var (
	ErrInvalidCredentials = &Error{
		Code:    ErrInvalidCredentialsCode,
		Message: "credentials are not valid",
	}
	ErrAlreadyExists = &Error{
		Code:    ErrUserAlreadyExistsCode,
		Message: "user already exists",
	}
	ErrAccessDenied = &Error{
		Code:    ErrAccessDeniedCode,
		Message: "access denied",
	}
	ErrNotFound = &Error{
		Code:    ErrUserNotFoundCode,
		Message: "user not found",
	}
)

func HandleStorageError(err error) error {
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return ErrInvalidCredentials
	}
	var ssmsErr *gorm.Error
	if !errors.As(err, &ssmsErr) {
		return fmt.Errorf("unknown error: %w", err)
	}
	switch ssmsErr.Code {
	case gorm.ErrNotFoundCode:
		err = ErrNotFound
	case gorm.ErrAlreadyExistsCode:
		err = ErrAlreadyExists
	default:
		err = fmt.Errorf("unknown ssms error: %w", ssmsErr)
	}

	return err
}
