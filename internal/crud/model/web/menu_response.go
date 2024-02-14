package web

import (
	"time"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
)

type MenuMasterUserResponse struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Title      string  `json:"title"`
	Path       *string `json:"path"`
	IconName   *string `json:"icon_name"`
	IsSubMenu  bool    `json:"is_submenu"`
	ParentName *string `json:"parent_name"`
	Create     bool    `json:"create"`
	Read       bool    `json:"read"`
	Update     bool    `json:"update"`
	Delete     bool    `json:"delete"`
}

type MenuMasterResponse struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Title      string     `json:"title"`
	Path       *string    `json:"path"`
	IconName   *string    `json:"icon_name"`
	IsSubMenu  bool       `json:"is_submenu"`
	ParentName *string    `json:"parent_name"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type MenuMasterRequest struct {
	Name       string  `json:"name" validate:"required,lowercase"`
	Title      string  `json:"title" validate:"required"`
	Path       *string `json:"path" validate:"required"`
	IconName   *string `json:"icon_name" validate:"required"`
	IsSubMenu  bool    `json:"is_submenu"`
	ParentName *string `json:"parent_name"`
}

type MenuMasterRequestEdit struct {
	Name       string  `json:"name" validate:"required,lowercase"`
	Title      string  `json:"title" validate:"required"`
	Path       *string `json:"path"`
	IconName   *string `json:"icon_name"`
	IsSubMenu  bool    `json:"is_submenu"`
	ParentName *string `json:"parent_name"`
}

func ToMenuMasterResponseWithError(menu domain.MenuMasterModel, err error) (MenuMasterResponse, error) {
	var menuResponse = MenuMasterResponse{
		Id:         menu.Id,
		Name:       menu.Name,
		Title:      menu.Title,
		Path:       menu.Path,
		IconName:   menu.IconName,
		IsSubMenu:  menu.IsSubMenu,
		ParentName: menu.ParentName,
		CreatedAt:  menu.CreatedAt,
		UpdatedAt:  menu.UpdatedAt,
	}

	return menuResponse, err
}

func ToMenuMasterResponse(menu domain.MenuMasterModel) MenuMasterResponse {
	var menuResponse = MenuMasterResponse{
		Id:         menu.Id,
		Name:       menu.Name,
		Title:      menu.Title,
		Path:       menu.Path,
		IconName:   menu.IconName,
		IsSubMenu:  menu.IsSubMenu,
		ParentName: menu.ParentName,
		CreatedAt:  menu.CreatedAt,
		UpdatedAt:  menu.UpdatedAt,
	}

	return menuResponse
}

func ToMenuMasterResponses(menus []domain.MenuMasterModel, totalRow int64, err error) ([]MenuMasterResponse, int64, error) {
	var menuResponse []MenuMasterResponse
	for _, menu := range menus {
		menuResponse = append(menuResponse, ToMenuMasterResponse(menu))
	}
	return menuResponse, totalRow, err
}
