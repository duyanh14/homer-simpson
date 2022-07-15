package api

import (
	"net/http"
	"simpson/internal/dto"
	"simpson/internal/helper"
	"simpson/internal/helper/logger"
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
	return helper.WithContext(func(ctx *helper.ContextGin) {

		var (
			req = dto.PartnerDTO{}
			log = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		// err = h.valida.Struct(req)
		// if err != nil {
		// 	log.Errorf("error while validator err %v", err)
		// 	ctx.BadRequest(err)
		// 	return
		// }
		// log.Info("start call api add partner")
		if err := h.partnerUsecase.AddPartner(ctx, req); err != nil {
			log.Error("error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(nil)
	})
}

func (h *partnerRouter) deletePartner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := logger.GetLogger()
		log.Debug("delete partner")
		log.Debug("delete partner 1 ")
		log.Debug("delete partner 2")
		log.Debug("delete partner 3")
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
