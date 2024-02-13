package domain

import (
	"time"

	"gorm.io/gorm"
)

type UserAccessMenuModel struct {
	Id              int             `json:"id"`
	RoleId          int             `json:"role_id"`
	MenuId          int             `json:"menu_id"`
	RoleMasterModel RoleMasterModel `gorm:"foreignKey:RoleId"`
	MenuMasterModel MenuMasterModel `gorm:"foreignKey:MenuId"`
	Create          bool            `json:"create"`
	Read            bool            `json:"read"`
	Update          bool            `json:"update"`
	Delete          bool            `json:"delete"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       gorm.DeletedAt  `json:"deleted_at"`
}

func (UserAccessMenuModel) TableName() string {
	return "user_access_menu"
}
