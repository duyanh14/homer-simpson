package api

import "github.com/gin-gonic/gin"

func (h *userRouter) UserRouter(router *gin.RouterGroup) {
	router.GET("/user/register", h.register())
	router.GET("/user/login", h.login())
}

func (h *roleRouter) RoleRouter(router *gin.RouterGroup) {
	router.DELETE("/role", h.deleteRole())
	router.POST("/role", h.addRole())
}

func (h *permissionRouter) PermissionRouter(router *gin.RouterGroup) {
	router.POST("/permission", h.addPermission())
	router.DELETE("/permission", h.deletePermission())
}
