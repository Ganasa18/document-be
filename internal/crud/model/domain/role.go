package domain

import (
	"time"

	"gorm.io/gorm"
)

type RoleMasterModel struct {
	Id        int            `json:"id"`
	RoleName  string         `json:"role_name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (RoleMasterModel) TableName() string {
	return "ms_roles"
}
