package api

import (
	"simpson/internal/dto"
	"simpson/internal/helper"
	"simpson/internal/helper/logger"
	"simpson/internal/usecase"

	"github.com/gin-gonic/gin"
)

type permissionRouter struct {
	permissionUsecase usecase.PermissionUsecase
}

func NewPermissionHandler(
	permissionUsecase usecase.PermissionUsecase,
) permissionRouter {
	return permissionRouter{
		permissionUsecase: permissionUsecase,
	}
}

func (h *permissionRouter) addPermission() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.AddPermissionReqDTO{}
			log = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("add permission, error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		err = h.permissionUsecase.AddPermission(ctx, req)
		if err != nil {
			log.Error(" add permission, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(nil)
	})

}

func (h *permissionRouter) deletePermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
