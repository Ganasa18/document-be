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
type UserAccessResponseJoinRoleAndMenu struct {
	Id        int       `json:"id"`
	RoleId    int       `json:"role_id"`
	MenuId    int       `json:"menu_id"`
	MenuName  string    `json:"menu_name"`
	MenuTitle string    `json:"menu_title"`
	RoleName  string    `json:"role_name"`
	Create    bool      `json:"create"`
	Read      bool      `json:"read"`
	Update    bool      `json:"update"`
	Delete    bool      `json:"delete"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func ToUserAccessMasterResponseJoinRoleAndMenu(userAccess domain.UserAccessMenuModel) UserAccessResponseJoinRoleAndMenu {
	var userAccessResponse = UserAccessResponseJoinRoleAndMenu{
		Id:        userAccess.Id,
		RoleId:    userAccess.RoleId,
		RoleName:  userAccess.RoleMasterModel.RoleName,
		MenuId:    userAccess.MenuId,
		MenuName:  userAccess.MenuMasterModel.Name,
		MenuTitle: userAccess.MenuMasterModel.Title,
		Create:    userAccess.Create,
		Read:      userAccess.Read,
		Update:    userAccess.Update,
		Delete:    userAccess.Delete,
		CreatedAt: userAccess.CreatedAt,
		UpdatedAt: userAccess.UpdatedAt,
	}

	return userAccessResponse
}

func ToUserAccessMasterResponses(userAccess []domain.UserAccessMenuModel, totalRow int64, err error) ([]UserAccessResponseJoinRoleAndMenu, int64, error) {
	var userAccessResponse []UserAccessResponseJoinRoleAndMenu
	for _, userAccs := range userAccess {
		userAccessResponse = append(userAccessResponse, ToUserAccessMasterResponseJoinRoleAndMenu(userAccs))
	}

	return userAccessResponse, totalRow, err
}
