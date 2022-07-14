package dto

type (
	PartnerDTO struct {
		Name       string `json:"name" validate:"required,min=3"`
		Code       string `json:"code" validate:"required,min=3,max=10"`
		Descrition string `json:"description"`
		Alias      string `json:"alias"`
		Address    string `json:"address"`
	}
)
