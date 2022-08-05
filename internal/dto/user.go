package dto

type (
	UserDTO struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
	}

	UserLoginReqDTO struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	UserLoginRespDTO struct {
		Jwt string `json:"jwt"`
	}

	UserVerifyDTO struct {
		Jwt string `json:"jwt"`
	}
)

type (
	CheckAccessRespDTO struct {
		IsAccess bool `json:"is_access"`
	}
	CheckAccessReqDTO struct {
		PermissionCode string `json:"code"`
	}
)
