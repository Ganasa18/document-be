package domain

import (
	"time"

	"gorm.io/gorm"
)

type ForgotPasswordLink struct {
	gorm.Model
	Id        int       `json:"id" gorm:"primaryKey"`
	LinkUrl   string    `json:"link_url"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}
