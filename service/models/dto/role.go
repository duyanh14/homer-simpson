package dto

type RoleDeleteRequestDTO struct {
	ID string `json:"id"`
}

type RoleRequestDTO struct {
	RoleName        string `json:"role_name"`
	RoleAlias       string `json:"role_alias"`
	RoleDescription string `json:"role_description"`
	CreatedBy       string `json:"created_by"`
}

type RoleResponseDTO struct {
	RoleID          uint32
	RoleName        string
	RoleAlias       string
	RoleDescription string
	CreatedBy       string
	CreatedDate     string
	IsActive        bool
}

type RoleUpdateRequestDTO struct {
	ID              int64  `json:"id"`
	RoleName        string `json:"role_name"`
	RoleAlias       string `json:"role_alias"`
	RoleDescription string `json:"role_description"`
	CreatedBy       string `json:"created_by"`
}
