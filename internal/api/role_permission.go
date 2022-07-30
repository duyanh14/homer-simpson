package api

import (
	"simpson/internal/dto"
	"simpson/internal/helper"
	"simpson/internal/helper/logger"
	"simpson/internal/usecase"

	"github.com/gin-gonic/gin"
)

type rolePermissionRouter struct {
	rolePermissionUsecase usecase.RolePermissionUsecase
}

func NewRolePermissionHandler(
	rolePermissionUsecase usecase.RolePermissionUsecase,
) rolePermissionRouter {
	return rolePermissionRouter{
		rolePermissionUsecase: rolePermissionUsecase,
	}
}

func (h *rolePermissionRouter) addRolePermission() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.AddRolePermissionReqDTO{}
			log = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("add role permission, error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		err = h.rolePermissionUsecase.AddRolePermission(ctx, req)
		if err != nil {
			log.Error("add role permission, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(nil)
	})

}
