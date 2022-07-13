package common

import "errors"

var (
	CommonError   = errors.New("common error")
	DatabaseError = errors.New("database error")
)

var ErrorCodeValue = map[error]string{
	CommonError:   "-1",
	DatabaseError: "-2",
}

func ErrorMessage(err error) string {
	return err.Error()
}

func ErrorCode(err error) string {
	code, ok := ErrorCodeValue[err]
	if !ok {
		return ""
	}
	return code
}
