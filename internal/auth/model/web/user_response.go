package web

import (
	"time"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
)

type UserRegisterRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `json:"password"`
	OpenId   string `validate:"required" json:"open_id"`
}

type UserRegisterResponse struct {
	Id           int        `json:"id"`
	OpenId       string     `json:"open_id"`
	UserUniqueId string     `json:"user_unique_id"`
	Username     *string    `json:"username"`
	Email        string     `json:"email"`
	IsActive     bool       `json:"is_active"`
	RoleId       *int       `json:"role_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

func ToUserRegisterResponse(user domain.UserModel, errorData error) (UserRegisterResponse, error) {
	loginResponse := UserRegisterResponse{
		Id:           user.Id,
		UserUniqueId: user.UserUniqueId,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
	}

	return loginResponse, errorData
}
