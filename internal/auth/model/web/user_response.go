package web

import (
	"time"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
	role "github.com/Ganasa18/document-be/internal/crud/model/web"
	"gorm.io/gorm"
)

type UserRegisterRequest struct {
	Email        string  `validate:"required,email" json:"email"`
	Password     string  `json:"password"`
	OpenId       string  `validate:"required" json:"open_id"`
	Username     *string `json:"username"`
	ProfileImage *string `json:"profile_image"`
}

type ProfileUserResponse struct {
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	ProfileImage *string `json:"profile_image"`
}

type UserRegisterResponse struct {
	Token        *string                     `json:"token"`
	Id           int                         `json:"id"`
	OpenId       string                      `json:"open_id"`
	UserUniqueId string                      `json:"user_unique_id"`
	Username     *string                     `json:"username"`
	Email        string                      `json:"email"`
	IsActive     bool                        `json:"is_active"`
	Role         role.RoleMasterResponseJoin `json:"role"`
	Profile      ProfileUserResponse         `json:"profile"`
	CreatedAt    time.Time                   `json:"created_at"`
	UpdatedAt    time.Time                   `json:"updated_at"`
	DeletedAt    *gorm.DeletedAt             `json:"deleted_at"`
}

func ToUserRegisterResponse(user domain.UserModel, profile domain.ProfileUser, errorData error) (UserRegisterResponse, error) {

	userRole := role.RoleMasterResponseJoin{
		Id:       user.RoleMasterModel.Id,
		RoleName: user.RoleMasterModel.RoleName,
	}

	userProfile := ProfileUserResponse{
		FirstName:    profile.FirstName,
		LastName:     profile.LastName,
		ProfileImage: profile.ProfileImage,
	}

	loginResponse := UserRegisterResponse{
		Id:           user.Id,
		OpenId:       user.OpenId,
		UserUniqueId: user.UserUniqueId,
		Email:        user.Email,
		Username:     user.Username,
		Role:         userRole,
		Profile:      userProfile,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return loginResponse, errorData
}
