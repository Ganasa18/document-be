package service

import (
	"fmt"
	"time"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
	"github.com/Ganasa18/document-be/internal/auth/model/web"
	"github.com/Ganasa18/document-be/internal/auth/repository"
	"github.com/Ganasa18/document-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	Validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		Validate:       validate,
	}
}

func toUserRegisterResponse(user domain.UserModel) web.UserRegisterRequest {
	return web.UserRegisterRequest{
		Email:    user.Email,
		Password: user.Password,
	}
}

func (service *AuthServiceImpl) LoginOrRegister(ctx *gin.Context, request web.UserRegisterRequest) web.UserRegisterRequest {
	fmt.Println(request, "REQUEST")
	err := service.Validate.Struct(request)
	utils.PanicIfError(err)

	// LOGIC
	register := domain.UserModel{
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	register = service.AuthRepository.LoginOrRegister(ctx, register)
	return toUserRegisterResponse(register)

}
