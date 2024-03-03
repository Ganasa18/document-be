package web

import "time"

type LoginLogResponse struct {
	Id        int       `json:"id"`
	Agent     string    `json:"agent"`
	Email     string    `json:"email"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
}
