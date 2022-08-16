package api

import (
	"simpson/internal/dto"
	"simpson/internal/helper"
	"simpson/internal/helper/logger"
	"simpson/internal/usecase"
	"strconv"

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

func (h *permissionRouter) updatePermission() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.UpdatePermissionReqDTO{}
			log = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("update permission, error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		err = h.permissionUsecase.UpdateRole(ctx, req)
		if err != nil {
			log.Error("update permission, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(nil)
	})
}

func (h *permissionRouter) listPermission() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req  = dto.ListPermissionReqDTO{}
			resp = []dto.Permission{}
			log  = logger.GetLogger()
			err  error
		)
		// err = ctx.ShouldBindJSON(&req)
		// if err != nil {
		// 	log.Error("list permission, error while bind json %v", err)
		// 	ctx.BadRequest(err)
		// 	return
		// }
		resp, err = h.permissionUsecase.ListPermission(ctx, req)
		if err != nil {
			log.Error("list permission, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(resp)
	})
}

func (h *permissionRouter) deletePermission() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.DeletePermissionReqDTO{}
			log = logger.GetLogger()
		)
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Error("delete permission, error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		err = h.permissionUsecase.DeletePermission(ctx, req)
		if err != nil {
			log.Error("delete permission, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(nil)
	})
}

func (h *permissionRouter) detailPermission() gin.HandlerFunc {
	return helper.WithContext(func(ctx *helper.ContextGin) {
		var (
			req = dto.DetailPermissionReqDTO{}
			// resp = []dto.Role{}
			log = logger.GetLogger()
			err error
		)
		id := ctx.Query("permission_id")
		permsisonID, err := strconv.Atoi(id)
		if err != nil {
			log.Error("detail permisison, error while bind json %v", err)
			ctx.BadRequest(err)
			return
		}
		req.PermissionID = uint(permsisonID)
		resp, err := h.permissionUsecase.DetailPermission(ctx, req)
		if err != nil {
			log.Error("detail role, error %w", err)
			ctx.BadLogic(err)
			return
		}
		ctx.OKResponse(resp)
	})
}
