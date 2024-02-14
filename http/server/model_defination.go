package server

import (
	authController "github.com/Ganasa18/document-be/internal/auth/controller"
	authRepository "github.com/Ganasa18/document-be/internal/auth/repository"
	authService "github.com/Ganasa18/document-be/internal/auth/service"
	crudController "github.com/Ganasa18/document-be/internal/crud/controller"
	crudRepository "github.com/Ganasa18/document-be/internal/crud/repository"
	crudService "github.com/Ganasa18/document-be/internal/crud/service"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InitializeModel(db *gorm.DB, validate *validator.Validate) (authController.AuthController, crudController.RoleController, crudController.MenuController) {
	// Auth module
	authRepo := authRepository.NewAuthRepository(db)
	authSvc := authService.NewAuthService(authRepo, validate)
	authCtrl := authController.NewAuthController(authSvc)

	// Role module
	roleRepo := crudRepository.NewRoleRepository(db)
	roleSvc := crudService.NewRoleService(roleRepo, validate)
	roleCtrl := crudController.NewRoleController(roleSvc)

	// Menu Module
	menuRepo := crudRepository.NewMenuRepository(db)
	menuSvc := crudService.NewMenuService(menuRepo, validate)
	menuCtrl := crudController.NewMenuControllere(menuSvc)

	return authCtrl, roleCtrl, menuCtrl
}
