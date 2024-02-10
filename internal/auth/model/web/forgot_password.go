package web

import "time"

type ForgotPasswordResponse struct {
	LinkUrl   string    `json:"link_url"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}
