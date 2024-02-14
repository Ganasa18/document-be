package domain

import (
	"time"

	"gorm.io/gorm"
)

type MenuMasterModel struct {
	Id         int            `json:"id"`
	Name       string         `json:"name"`
	Title      string         `json:"title"`
	Path       *string        `json:"path"`
	IconName   *string        `json:"icon_name"`
	IsSubMenu  bool           `json:"is_submenu"`
	ParentName *string        `json:"parent_name"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type MenuUpdate struct {
	Name       string  `json:"name"`
	Title      string  `json:"title"`
	Path       *string `json:"path"`
	IconName   *string `json:"icon_name"`
	IsSubMenu  bool    `json:"is_sub_menu"`
	ParentName *string `json:"parent_name"`
}

func (MenuMasterModel) TableName() string {
	return "ms_menu"
}
