package domain

import (
	"time"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"gorm.io/gorm"
)

type UserModel struct {
	Id              int                    `json:"id" gorm:"primaryKey"`
	UserUniqueId    string                 `json:"user_unique_id"`
	OpenId          string                 `json:"open_id" gorm:"type:varchar(100)"`
	Username        *string                `json:"username" gorm:"type:varchar(255)"`
	Email           string                 `json:"email" gorm:"type:varchar(100);unique"`
	Password        *string                `json:"password"`
	IsActive        bool                   `json:"is_active"`
	RoleId          *int                   `json:"role_id"`
	RoleMasterModel domain.RoleMasterModel `gorm:"foreignKey:RoleId"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
	DeletedAt       gorm.DeletedAt         `json:"deleted_at"`
}

func (UserModel) TableName() string {
	return "ms_users"
}
