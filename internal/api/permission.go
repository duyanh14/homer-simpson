package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type permissionRouter struct {
}

func NewPermissionHandler() permissionRouter {
	return permissionRouter{}
}

func (h *permissionRouter) addPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	}
}

func (h *permissionRouter) deletePermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
