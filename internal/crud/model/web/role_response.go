package web

import (
	"time"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
)

type RoleMasterResponse struct {
	Id        int        `json:"id"`
	RoleName  string     `json:"role_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type RoleMasterCreateRequest struct {
	RoleName string `validate:"required,min:1" json:"role_name"`
}

func ToRoleMasterResponse(role domain.RoleMasterModel) RoleMasterResponse {
	return RoleMasterResponse{
		Id:        role.Id,
		RoleName:  role.RoleName,
		CreatedAt: role.CreatedAt,
		DeletedAt: role.DeletedAt,
	}
}

func ToRoleMasterResponses(roles []domain.RoleMasterModel, totalRow int64, err error) ([]RoleMasterResponse, int64, error) {
	var roleResponse []RoleMasterResponse
	for _, role := range roles {
		roleResponse = append(roleResponse, ToRoleMasterResponse(role))
	}
	return roleResponse, totalRow, err
}
