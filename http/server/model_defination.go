package server

import (
	authController "github.com/Ganasa18/document-be/internal/auth/controller"
	authRepository "github.com/Ganasa18/document-be/internal/auth/repository"
	authService "github.com/Ganasa18/document-be/internal/auth/service"
	crudController "github.com/Ganasa18/document-be/internal/crud/controller"
	crudRepository "github.com/Ganasa18/document-be/internal/crud/repository"
	crudService "github.com/Ganasa18/document-be/internal/crud/service"
	loggingRepository "github.com/Ganasa18/document-be/internal/logging/repository"
	wsController "github.com/Ganasa18/document-be/internal/websocket/controller"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InitializeModel(db *gorm.DB, validate *validator.Validate) (authController.AuthController, crudController.RoleController, crudController.MenuController, crudController.UserAccessController, wsController.WebSocketController) {
	// Auth module

	authRepo := authRepository.NewAuthRepository(db)
	logRepo := loggingRepository.NewLoggingRepository(db)
	authSvc := authService.NewAuthService(authRepo, logRepo, validate)
	authCtrl := authController.NewAuthController(authSvc)

	// Role module
	roleRepo := crudRepository.NewRoleRepository(db)
	roleSvc := crudService.NewRoleService(roleRepo, validate)
	roleCtrl := crudController.NewRoleController(roleSvc)

	// Menu Module
	menuRepo := crudRepository.NewMenuRepository(db)
	menuSvc := crudService.NewMenuService(menuRepo, validate)
	menuCtrl := crudController.NewMenuControllere(menuSvc)

	// User Access Module
	userAccessRepo := crudRepository.NewUserAccessRepository(db)
	userAccessSvc := crudService.NewUserAccessService(userAccessRepo, validate)
	userAccessCtrl := crudController.NewUserAccessControllere(userAccessSvc)

	// WebSocket Module
	webSocketCtrl := wsController.NewWebSocketController()

	return authCtrl, roleCtrl, menuCtrl, userAccessCtrl, webSocketCtrl
}
