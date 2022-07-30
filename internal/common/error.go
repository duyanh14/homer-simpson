package common

import "errors"

var (
	ErrCommon   = errors.New("common error")
	ErrDatabase = errors.New("database error")
	// user
	ErrTokenExpired        = errors.New("token expired")
	ErrTokenInvalid        = errors.New("token invalid")
	ErrUserNotFound        = errors.New("user not found")
	ErrUserIDNotFoundInJwt = errors.New("get userid in jwt failed")
	// role
	ErrRoleNotFound    = errors.New("role not found")
	ErrRoleCodeRequire = errors.New("role code is required")

	// permisson
	ErrPermisisonNotFound    = errors.New("permisison not found")
	ErrPermissionCodeRequire = errors.New("permisison code is required")
)

var ErrorCodeValue = map[error]string{
	ErrCommon:              "-1",
	ErrDatabase:            "-2",
	ErrTokenExpired:        "-100",
	ErrTokenInvalid:        "-101",
	ErrUserNotFound:        "-102",
	ErrUserIDNotFoundInJwt: "-103",

	ErrRoleNotFound: "-200",

	ErrPermisisonNotFound:    "-300",
	ErrPermissionCodeRequire: "-301",
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
