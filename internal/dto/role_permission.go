package dto

type (
	AddRolePermissionReqDTO struct {
		RoleId       uint   `json:"role_id"`
		PermissionID uint   `json:"permission_id"`
		Description  string `json:"description"`
	}
)

type GetListPermissionOfRole struct {
	RoleId string `json:"role_id"`
}
