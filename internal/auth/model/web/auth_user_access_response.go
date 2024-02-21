package web

type UpdateUserAccessRequest struct {
	UserRole     int    `json:"user_role" validate:"required"`
	UserUniqueId string `json:"user_unique_id" validate:"required"`
}
