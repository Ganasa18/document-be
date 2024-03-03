package domain

import (
	"time"

	"gorm.io/gorm"
)

type LoginLogModel struct {
	Id        int            `json:"id" gorm:"primaryKey"`
	Agent     string         `json:"agent"`
	Email     string         `json:"email"`
	Action    string         `json:"action"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (LoginLogModel) TableName() string {
	return "tb_log_login"
}
