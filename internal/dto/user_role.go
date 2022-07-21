package dto

type (
	AddUserRoleReqDTO struct {
		RoleIds     []uint `json:"role_ids"`
		UserID      uint   `json:"user_id"`
		Description string `json:"description"`
	}
)
