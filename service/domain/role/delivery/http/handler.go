package http

import (
	"golang-course/service/domain/role/usecase"
	"golang-course/service/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleUsecase usecase.IRoleUsecase
}

func NewRoleHandler(
	roleUsecase usecase.IRoleUsecase,
) *RoleHandler {
	return &RoleHandler{
		roleUsecase: roleUsecase,
	}
}

func (h *RoleHandler) addRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := dto.RoleRequestDTO{}
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		err = h.roleUsecase.AddRole(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK,
			dto.Response{
				Status:        "OK",
				ReasonCode:    200,
				ReasonMessage: "",
			})
	}
}

func (h *RoleHandler) getListRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		list, err := h.roleUsecase.GetListRole(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, list)
	}
}

func (h *RoleHandler) deleteRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := dto.RoleDeleteRequestDTO{}
		req.ID = ctx.Param("id")
		if req.ID == "" {
			ctx.JSON(http.StatusBadRequest, dto.ResponseError{
				Status:        "FAIL",
				ReasonMessage: "ID empty",
				ReasonCode:    400,
			})
			return
		}
		err := h.roleUsecase.DeleteRole(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK,
			dto.Response{
				Status:        "OK",
				ReasonCode:    200,
				ReasonMessage: "",
			})
	}
}

func (h *RoleHandler) updateRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// h.roleUsecase.UpdateRole(ctx)
		// ctx.JSON(http.StatusOK, "updateRole")
		req := dto.RoleUpdateRequestDTO{}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err := h.roleUsecase.UpdateRole(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK,
			dto.Response{
				Status:        "OK",
				ReasonCode:    200,
				ReasonMessage: "",
			})

	}
}
