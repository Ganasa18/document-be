package domain

import "time"

type UserModel struct {
	Id           int        `json:"id"`
	UserUniqueId string     `json:"user_unique_id"`
	OpenId       string     `json:"open_id"`
	Username     *string    `json:"username"`
	Email        string     `json:"email"`
	Password     *string    `json:"password"`
	IsActive     bool       `json:"is_active"`
	RoleId       *int       `json:"role_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

func (UserModel) TableName() string {
	return "ms_users"
}
