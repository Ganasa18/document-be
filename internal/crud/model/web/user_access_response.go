package web

import (
	"time"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"gorm.io/gorm"
)

type UserAccessResponse struct {
	Id        int            `json:"id"`
	RoleId    int            `json:"role_id"`
	MenuId    int            `json:"menu_id"`
	Create    bool           `json:"create"`
	Read      bool           `json:"read"`
	Update    bool           `json:"update"`
	Delete    bool           `json:"delete"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type UserAccessRequest struct {
	RoleId int  `json:"role_id" validate:"required"`
	MenuId int  `json:"menu_id" validate:"required"`
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

func ToUserAccessMasterResponseWithError(userAccess domain.UserAccessMenuModel, err error) (UserAccessResponse, error) {
	var userAccessResponse = UserAccessResponse{
		Id:        userAccess.Id,
		RoleId:    userAccess.RoleId,
		MenuId:    userAccess.MenuId,
		Create:    userAccess.Create,
		Read:      userAccess.Read,
		Update:    userAccess.Update,
		Delete:    userAccess.Delete,
		CreatedAt: userAccess.CreatedAt,
		UpdatedAt: userAccess.UpdatedAt,
		DeletedAt: userAccess.DeletedAt,
	}

	return userAccessResponse, err
}
func ToUserAccessMasterResponse(userAccess domain.UserAccessMenuModel) UserAccessResponse {
	var userAccessResponse = UserAccessResponse{
		Id:        userAccess.Id,
		RoleId:    userAccess.RoleId,
		MenuId:    userAccess.MenuId,
		Create:    userAccess.Create,
		Read:      userAccess.Read,
		Update:    userAccess.Update,
		Delete:    userAccess.Delete,
		CreatedAt: userAccess.CreatedAt,
		UpdatedAt: userAccess.UpdatedAt,
		DeletedAt: userAccess.DeletedAt,
	}

	return userAccessResponse
}
