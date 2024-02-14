package server

import "github.com/Ganasa18/document-be/pkg/middleware"

func (h *HttpServe) setupRouter() {
	v1 := h.router.Group("/api/v1")
	// AUTH
	v1.POST("/login-or-register", h.authController.LoginOrRegister)
	v1.POST("/forgot-password", h.authController.ForgotLinkPassword)
	v1.POST("/reset-password", h.authController.ResetPasswordUser)

	// WITH AUTHORZATION
	v1.Use(middleware.CustomAuthMiddleware())
	// ROLE MASTER
	v1.GET("/crud/role", h.roleController.GetRoles)
	v1.GET("/crud/role/:roleId", h.roleController.GetRoleById)
	v1.POST("/crud/role", h.roleController.CreateRole)
	v1.PATCH("/crud/role/:roleId", h.roleController.UpdateRole)
	v1.DELETE("/crud/role/:roleId", h.roleController.DeleteRole)
	// MENU MASTER
	v1.GET("/crud/menu", h.menuController.GetAllMenu)
	v1.GET("/crud/menu/:menuId", h.menuController.GetMenuById)
	v1.POST("/crud/menu", h.menuController.CreateMenu)
	v1.PATCH("/crud/menu/:menuId", h.menuController.CreateMenu)
	v1.DELETE("/crud/menu/:menuId", h.menuController.DeleteMenu)

}
