package api

import (
	"net/http"
	"simpson/internal/common"
	"simpson/internal/dto"
	"simpson/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type partnerRouter struct {
	partnerUsecase usecase.PartnerUsecase
	valida         *validator.Validate
}

func NewPartnerHandler(
	partnerUsecase usecase.PartnerUsecase,
	valida *validator.Validate,
) partnerRouter {
	return partnerRouter{
		partnerUsecase: partnerUsecase,
		valida:         valida,
	}
}

func (h *partnerRouter) addPartner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req = dto.PartnerDTO{}
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		err = h.valida.Struct(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ResponseError(err))
			return
		}
		if err := h.partnerUsecase.AddPartner(ctx, req); err != nil {
			ctx.JSON(http.StatusOK, common.ResponseError(err))
			return
		}
		ctx.JSON(http.StatusOK, common.ResponseOK())
	}
}

func (h *partnerRouter) deletePartner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, abc{
			Name: "ducnp",
		})

	}
}

func (h *partnerRouter) getPartnerDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, abc{
			Name: "ducnp",
		})

	}
}

func (h *partnerRouter) getPartnerList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, abc{
			Name: "ducnp",
		})

	}
}

func (h *partnerRouter) updatePartner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, abc{
			Name: "ducnp",
		})
	}
}
