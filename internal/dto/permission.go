package dto

type (
	AddPermissionReqDTO struct {
		Name        string `json:"name"`
		Alias       string `json:"alias"`
		Code        string `json:"code"`
		Description string `json:"description"`
	}
)
