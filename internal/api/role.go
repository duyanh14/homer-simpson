package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type roleRouter struct {
}

func NewRoleHandler() roleRouter {
	return roleRouter{}
}

func (h *roleRouter) addRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	}
}

func (h *roleRouter) deleteRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	}
}
