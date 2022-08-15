package api

import (
	"simpson/internal/dto"
	"simpson/internal/helper"
	"simpson/internal/helper/logger"
	"simpson/internal/usecase"
	"strconv"

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
			err  error
		)
		// TODO
		req.IsActive = false
		if active := ctx.Query("active"); active != "" {
			req.IsActive = true
		}
		// err := ctx.ShouldBindJSON(&req)
		// if err != nil {
		// 	log.Error("list role, error while bind json %v", err)
		// 	ctx.BadRequest(err)
		// 	return
		// }
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

func (h *roleRouter) detailRole() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.DetailRoleReqDTO{}
			// resp = []dto.Role{}
			log = logger.GetLogger()
			err error
		)
		id := ctx.Query("role_id")
		roleID, err := strconv.Atoi(id)
		if err != nil {
			log.Error("delete role, error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		req.RoleID = uint(roleID)
		resp, err := h.roleUsecase.DetailRole(ctx, req)
		if err != nil {
			log.Error("detail role, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(resp)
	})
}
