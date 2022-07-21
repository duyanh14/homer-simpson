package api

import (
	"simpson/internal/dto"
	"simpson/internal/helper"
	"simpson/internal/helper/logger"
	"simpson/internal/usecase"

	"github.com/gin-gonic/gin"
)

type userRoleRouter struct {
	userRoleUsecase usecase.UserRoleUsecase
}

func NewUserRoleHandler(
	userRoleUsecase usecase.UserRoleUsecase,
) userRoleRouter {
	return userRoleRouter{
		userRoleUsecase: userRoleUsecase,
	}
}

func (h *userRoleRouter) addUserRole() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.AddUserRoleReqDTO{}
			log = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("add user role, error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		err = h.userRoleUsecase.AddUserRole(ctx, req)
		if err != nil {
			log.Error("add user role, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(nil)
	})

}
