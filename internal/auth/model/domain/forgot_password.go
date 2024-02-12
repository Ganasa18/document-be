package domain

import (
	"time"
)

type ForgotPasswordLink struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	LinkUrl   string    `json:"link_url"`
	HashId    string    `json:"hash_id"`
	IsActive  *bool     `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
}
