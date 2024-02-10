package domain

import (
	"time"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
)

type UserModel struct {
	Id              int                    `json:"id" gorm:"primaryKey"`
	UserUniqueId    string                 `json:"user_unique_id"`
	OpenId          string                 `json:"open_id"`
	Username        *string                `json:"username"`
	Email           string                 `json:"email"`
	Password        *string                `json:"password"`
	IsActive        bool                   `json:"is_active"`
	RoleId          *int                   `json:"role_id"`
	RoleMasterModel domain.RoleMasterModel `gorm:"foreignKey:RoleId"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
	DeletedAt       *time.Time             `json:"deleted_at"`
}

func (UserModel) TableName() string {
	return "ms_users"
}
