package validation

import (
	"fmt"
	"simpson/internal/dto"

	"github.com/go-playground/validator/v10"
)

func AddPartnerValidator(in dto.PartnerDTO) error {
	vali := validator.New()
	fmt.Println(vali)
	err := vali.Struct(in)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
