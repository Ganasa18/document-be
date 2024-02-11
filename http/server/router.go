package server

import "github.com/Ganasa18/document-be/pkg/middleware"

func (h *HttpServe) setupRouter() {
	v1 := h.router.Group("/api/v1")
	// AUTH
	v1.POST("/login-or-register", h.authController.LoginOrRegister)
	v1.POST("/forgot-password", h.authController.ForgotLinkPassword)

	// WITH AUTHORZATION
	v1.Use(middleware.CustomAuthMiddleware())
	// CRUD MASTER
	v1.GET("/crud/role", h.roleController.GetRoles)
	v1.GET("/crud/role/:roleId", h.roleController.GetRoleById)
	v1.POST("/crud/role", h.roleController.CreateRole)
	v1.PATCH("/crud/role/:roleId", h.roleController.UpdateRole)
	v1.DELETE("/crud/role/:roleId", h.roleController.DeleteRole)

}
