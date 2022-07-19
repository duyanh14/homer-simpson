package api

import (
	"simpson/internal/dto"
	"simpson/internal/helper"
	"simpson/internal/helper/logger"
	"simpson/internal/usecase"

	"github.com/gin-gonic/gin"
)

type roleRouter struct {
	roleUsecase usecase.RoleUsecase
}

func NewRoleHandler(
	roleUsecase usecase.RoleUsecase,
) roleRouter {
	return roleRouter{
		roleUsecase: roleUsecase,
	}
}

func (h *roleRouter) addRole() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.AddRoleReqDTO{}
			log = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("add role, error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		err = h.roleUsecase.AddRole(ctx, req)
		if err != nil {
			log.Error("add role, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(nil)
	})

}

func (h *roleRouter) deleteRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
