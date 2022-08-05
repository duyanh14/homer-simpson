package api

import (
	"simpson/internal/dto"
	"simpson/internal/helper"
	"simpson/internal/helper/logger"
	"simpson/internal/usecase"

	"github.com/gin-gonic/gin"
)

type accessRouter struct {
	accessUsecase usecase.AccessUsecase
}

func NewAccessHandler(
	accessUsecase usecase.AccessUsecase,
) accessRouter {
	return accessRouter{
		accessUsecase: accessUsecase,
	}
}

func (h *accessRouter) addAccess() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.AddAccessReqDTO{}
			log = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("add access, error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		err = h.accessUsecase.AddAccess(ctx, req)
		if err != nil {
			log.Error("add access, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(nil)
	})

}

func (h *accessRouter) deleteAccess() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
