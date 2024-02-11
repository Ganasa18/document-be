package web

import "time"

type ForgotPasswordResponse struct {
	LinkUrl   string    `json:"link_url"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type ForgotPasswordRequest struct {
	Email string `validate:"required" json:"email"`
}
