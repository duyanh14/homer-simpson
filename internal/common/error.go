package common

import "errors"

var (
	ErrCommon       = errors.New("common error")
	ErrDatabase     = errors.New("database error")
	ErrTokenExpired = errors.New("token expired")
	ErrTokenInvalid = errors.New("token invalid")
	ErrUserNotFound = errors.New("user not found")
	ErrRoleNotFound = errors.New("role not found")
)

var ErrorCodeValue = map[error]string{
	ErrCommon:       "-1",
	ErrDatabase:     "-2",
	ErrTokenExpired: "-3",
	ErrTokenInvalid: "-4",
	ErrUserNotFound: "-5",
	ErrRoleNotFound: "-6",
}

func ErrorMessage(err error) string {
	return err.Error()
}

func ErrorCode(err error) string {
	code, ok := ErrorCodeValue[err]
	if !ok {
		return "400"
	}
	return code
}
