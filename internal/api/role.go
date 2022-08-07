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

func (h *roleRouter) updateRole() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.UpdateRoleReqDTO{}
			log = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("update role, error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		err = h.roleUsecase.UpdateRole(ctx, req)
		if err != nil {
			log.Error("update role, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(nil)
	})
}

func (h *roleRouter) listRole() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req  = dto.ListRoleReqDTO{}
			resp = []dto.Role{}
			log  = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("list role, error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		resp, err = h.roleUsecase.ListRole(ctx, req)
		if err != nil {
			log.Error("list role, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(resp)
	})
}

func (h *roleRouter) deleteRole() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.DeleteRoleReqDTO{}
			log = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("delete role, error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		err = h.roleUsecase.DeleteRole(ctx, req)
		if err != nil {
			log.Error("delete role, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(nil)
	})
}
