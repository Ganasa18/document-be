package domain

import (
	"time"
)

type RoleMasterModel struct {
	Id        int        `json:"id"`
	RoleName  string     `json:"role_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (RoleMasterModel) TableName() string {
	return "ms_roles"
}
