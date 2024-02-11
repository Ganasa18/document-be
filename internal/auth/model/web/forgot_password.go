package web

import "time"

type ForgotPasswordResponse struct {
	LinkUrl   string    `json:"link_url"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type ForgotPasswordResponseSuccess struct {
	Status string `json:"status"`
	HashId string `json:"hash_id"`
}

type ForgotPasswordRequest struct {
	Email string `validate:"required" json:"email"`
}

type ResetPasswordRequest struct {
	Email       string `validate:"required" json:"email"`
	NewPassword string `validate:"required" json:"new_password"`
	HashId      string `validate:"required" json:"hash_id"`
}
