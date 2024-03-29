package api

import "github.com/gin-gonic/gin"

func (h *healthRouter) HealthRouter(router *gin.RouterGroup) {
	router.GET("/liveness", h.liveness())
	router.GET("/readiness", h.readiness())
}

func (h *userRouter) UserRouter(router *gin.RouterGroup) {
	router.POST("/user/verify", h.verifyToken())
	router.POST("/user/register", h.register())
	router.POST("/user/login", h.login())
	router.GET("/user/permissions", h.listPermission())
	router.GET("/user/access", h.checkAccess())
	router.GET("/user", h.userInfo())

	//router.GET("/api/v1/state", h.locationMap()) // not
}

func (h *userRoleRouter) UserRoleRouter(router *gin.RouterGroup) {
	// add list role for user,
	router.POST("/user/role", h.addUserRole())
}

func (h *accessRouter) UserAccessRouter(router *gin.RouterGroup) {
	router.POST("/access", h.addAccess())
}

func (h *rolePermissionRouter) RolePermissionRouter(router *gin.RouterGroup) {
	router.POST("/role/permission", h.addRolePermission())
	router.GET("/role/permissions", h.listPermissionByRole())
	router.GET("/permission/roles", h.listRoleByPermission())
}

func (h *roleRouter) RoleRouter(router *gin.RouterGroup) {
	router.DELETE("/role", h.deleteRole())
	router.POST("/role", h.addRole())
	router.GET("/roles", h.listRole())
	router.PUT("/role", h.updateRole())
	router.GET("/role", h.detailRole())
}

func (h *permissionRouter) PermissionRouter(router *gin.RouterGroup) {
	router.POST("/permission", h.addPermission())
	router.DELETE("/permission", h.deletePermission())
	router.GET("/permissions", h.listPermission())
	router.PUT("/permission", h.updatePermission())
	router.GET("/permission", h.detailPermission())
}

func (h *partnerRouter) PartnerRouter(router *gin.RouterGroup) {
	router.GET("/partner/:id", h.getPartnerDetail())
	router.GET("/partners", h.getPartnerList())
	router.POST("/partner", h.addPartner())
	router.PUT("/partner", h.updatePartner())
	router.DELETE("/partner/:id", h.deletePartner())

}
