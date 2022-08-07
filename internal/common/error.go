package common

import "errors"

// postgres code
var (
	DuplicateKeyValue string = "23505"
)

var (
	ErrCommon         = errors.New("common error")
	ErrDatabase       = errors.New("database error")
	ErrRecordNotFound = errors.New("record not found")
	// user
	ErrTokenExpired        = errors.New("token expired")
	ErrTokenInvalid        = errors.New("token invalid")
	ErrUserNotFound        = errors.New("user not found")
	ErrUserIDNotFoundInJwt = errors.New("get userid in jwt failed")
	ErrCodeAccessRequire   = errors.New("permisison code access require")

	// role
	ErrRoleNotFound     = errors.New("role not found")
	ErrRoleCodeRequire  = errors.New("role code is required")
	ErrRoleIdRequire    = errors.New("role id is required")
	ErrRoleCodeIsExists = errors.New("role code is exists")

	// permisson
	ErrPermisisonNotFound     = errors.New("permisison not found")
	ErrPermissionCodeRequire  = errors.New("permisison code is required")
	ErrPermissionIdRequire    = errors.New("permisison id is required")
	ErrPermissionCodeIsExists = errors.New("permisison code is exists")

	// accept
	ErrAccessTypeRequire = errors.New("accept type is required")
)

var ErrorCodeValue = map[error]string{
	ErrCommon:         "-1",
	ErrDatabase:       "-2",
	ErrRecordNotFound: "-3",

	ErrTokenExpired:        "-100",
	ErrTokenInvalid:        "-101",
	ErrUserNotFound:        "-102",
	ErrUserIDNotFoundInJwt: "-103",
	ErrCodeAccessRequire:   "-104",

	ErrRoleNotFound:     "-200",
	ErrRoleCodeRequire:  "-201",
	ErrRoleIdRequire:    "-202",
	ErrRoleCodeIsExists: "-203",

	ErrPermisisonNotFound:    "-300",
	ErrPermissionCodeRequire: "-301",
	ErrPermissionIdRequire:   "-302",

	ErrAccessTypeRequire: "-400",
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
