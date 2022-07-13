package http

import "github.com/gin-gonic/gin"

func (h *RoleHandler) RoleAPIRouter(router *gin.RouterGroup) {
	router.POST("/roles", h.addRole())
	router.GET("/roles", h.getListRole())
	router.DELETE("/roles/:id", h.deleteRole())
	router.PUT("/roles", h.updateRole())
}
