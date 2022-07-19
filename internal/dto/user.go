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
)
