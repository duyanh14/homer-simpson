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
		Jwt string `json:"jwt_str"`
	}

	UserVerifyDTO struct {
		Jwt string `json:"jwt"`
	}
)

type (
	CheckAccessRespDTO struct {
		Message     string   `json:"message"`
		IsAccess    bool     `json:"is_access"`
		Permissions []string `json:"permissions"`
	}
	CheckAccessReqDTO struct {
		PermissionCode string `json:"code"`
	}

	UserListPermission struct {
		Permissions []string `json:"permissions"`
	}
)

type (
	UserInfoRespDTO struct {
		Phone    string `json:"phone"`
		Email    string `json:"email"`
		Username string `json:"user_name"`
	}

	UserInfoReqDTO struct {
	}
)
