package web

import (
	"time"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
	role "github.com/Ganasa18/document-be/internal/crud/model/web"
)

type UserRegisterRequest struct {
	Email    string  `validate:"required,email" json:"email"`
	Password string  `json:"password"`
	OpenId   string  `validate:"required" json:"open_id"`
	Username *string `json:"username"`
}

type UserRegisterResponse struct {
	Id           int                         `json:"id"`
	OpenId       string                      `json:"open_id"`
	UserUniqueId string                      `json:"user_unique_id"`
	Username     *string                     `json:"username"`
	Email        string                      `json:"email"`
	IsActive     bool                        `json:"is_active"`
	Role         role.RoleMasterResponseJoin `json:"role"`
	Token        *string                     `json:"token"`
	CreatedAt    time.Time                   `json:"created_at"`
	UpdatedAt    time.Time                   `json:"updated_at"`
	DeletedAt    *time.Time                  `json:"deleted_at"`
}

func ToUserRegisterResponse(user domain.UserModel, errorData error) (UserRegisterResponse, error) {

	userRole := role.RoleMasterResponseJoin{
		Id:       user.RoleMasterModel.Id,
		RoleName: user.RoleMasterModel.RoleName,
	}

	loginResponse := UserRegisterResponse{
		Id:           user.Id,
		OpenId:       user.OpenId,
		UserUniqueId: user.UserUniqueId,
		Email:        user.Email,
		Username:     user.Username,
		Role:         userRole,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return loginResponse, errorData
}
