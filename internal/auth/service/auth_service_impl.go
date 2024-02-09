package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
	"github.com/Ganasa18/document-be/internal/auth/model/web"
	"github.com/Ganasa18/document-be/internal/auth/repository"
	"github.com/Ganasa18/document-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
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

func (service *AuthServiceImpl) LoginOrRegister(ctx *gin.Context, request web.UserRegisterRequest) (web.UserRegisterResponse, error) {
	err := service.Validate.Struct(request)
	utils.PanicIfError(err)
	var passwordData string

	// GENERATE UUID
	uniqueID := uuid.New().String()

	// LOGIC
	OpenId := request.OpenId
	register := domain.UserModel{
		UserUniqueId: uniqueID,
		Email:        request.Email,
		Password:     nil,
		OpenId:       request.OpenId,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if register.Password == nil && OpenId != utils.OPEN_API_GOOGLE {
		fmt.Println("Error creating user:", request)
		if request.Password != "" {
			passwordData = request.Password
			register.Password = &passwordData
		} else {
			return web.UserRegisterResponse{}, errors.New("password must be provided")
		}
	}

	data, err := service.AuthRepository.LoginOrRegister(ctx, register, OpenId)

	return web.ToUserRegisterResponse(data, err)

}
