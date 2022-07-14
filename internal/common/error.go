package common

import "errors"

var (
	ErrCommon   = errors.New("common error")
	ErrDatabase = errors.New("database error")
)

var ErrorCodeValue = map[error]string{
	ErrCommon:   "-1",
	ErrDatabase: "-2",
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
