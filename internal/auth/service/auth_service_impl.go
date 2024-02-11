package service

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	appconfig "github.com/Ganasa18/document-be/config"
	"github.com/Ganasa18/document-be/internal/auth/model/domain"
	"github.com/Ganasa18/document-be/internal/auth/model/web"
	"github.com/Ganasa18/document-be/internal/auth/repository"
	"github.com/Ganasa18/document-be/pkg/email"
	"github.com/Ganasa18/document-be/pkg/exception"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/Ganasa18/document-be/pkg/loghelper"
	"github.com/Ganasa18/document-be/pkg/queue"
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

	roleDefault := 1

	// LOGIC
	OpenId := request.OpenId
	register := domain.UserModel{
		UserUniqueId: uniqueID,
		Email:        request.Email,
		Password:     nil,
		OpenId:       request.OpenId,
		RoleId:       &roleDefault,
		Username:     request.Username,
	}

	if register.Password == nil && OpenId != utils.OPEN_API_GOOGLE {
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

func (service *AuthServiceImpl) ForgotLinkPassword(ctx *gin.Context, request web.ForgotPasswordRequest) error {

	err := service.Validate.Struct(request)
	utils.PanicIfError(err)

	jobs := make(chan queue.Job, 10)
	var wg sync.WaitGroup

	go func() {
		for job := range jobs {
			fmt.Printf("Processing job: %+v\n", job)
			updateForgotPassword := domain.ForgotPasswordLink{
				HashId: job.Payload,
			}
			err = service.AuthRepository.ExpiredForgotPassword(ctx, updateForgotPassword)
			if err != nil {
				loghelper.Errorln(ctx, fmt.Sprintf("ExpiredForgotPassword | Error when updating data, err:%s", err.Error()))
			}
			time.Sleep(1 * time.Second)
			wg.Done()
		}
	}()

	// MAKE LINK RANDOM
	uniqueID, err := helper.GenerateRandomString(100)
	utils.PanicIfError(err)

	baseHost := os.Getenv("APP_URL_FRONTEND")
	stringLink := fmt.Sprintf("%s/forgot-password/%s/reset", baseHost, uniqueID)

	// Expired after 5 minutes
	expiredAt := time.Now().Add(5 * time.Minute)

	forgotPassword := domain.ForgotPasswordLink{
		LinkUrl: stringLink,
		HashId:  uniqueID,
	}

	templateData := struct {
		Name string
		URL  string
	}{
		Name: request.Email,
		URL:  stringLink,
	}

	// Enqueue delayed jobs
	err = service.AuthRepository.ForgotLinkPassword(ctx, forgotPassword, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	wg.Add(3)
	queue.EnqueueDelayedJob(jobs, queue.Job{ID: time.Now().Local().String(), Payload: uniqueID, ExecuteAt: expiredAt})

	// Create a new request email
	res := email.NewRequest([]string{request.Email}, "My Apps | Forgot Password", "Forgot Password")
	// Parse the template with template data
	basePath := appconfig.InitAppConfig().AppDir
	templatePath := basePath + "/pkg/email/template/forgot-password.html"

	utils.PanicIfError(res.ParseTemplate(templatePath, templateData))
	// Send the email if template parsing is successful
	if ok, err := res.SendEmail(); err == nil {
		fmt.Println(ok)
	}

	return err
}
