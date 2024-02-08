package domain

import "time"

type UserModel struct {
	Id           int
	UserUniqueId string
	Username     string
	Email        string
	Password     string
	IsActive     bool
	RoleId       *int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
