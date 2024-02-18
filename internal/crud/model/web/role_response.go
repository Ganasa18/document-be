package web

import (
	"time"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"gorm.io/gorm"
)

type RoleMasterResponse struct {
	Id        int            `json:"id"`
	RoleName  string         `json:"role_name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type RoleMasterRequest struct {
	RoleName string `validate:"required,min=1" json:"role_name"`
}

type RoleMasterResponseJoin struct {
	Id       int    `json:"id"`
	RoleName string `json:"role_name"`
}

func ToRoleMasterResponseWithError(role domain.RoleMasterModel, err error) (RoleMasterResponse, error) {
	var roleResponse = RoleMasterResponse{
		Id:        role.Id,
		RoleName:  role.RoleName,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: role.DeletedAt,
	}

	return roleResponse, err
}

func ToRoleMasterResponse(role domain.RoleMasterModel) RoleMasterResponse {
	var roleResponse = RoleMasterResponse{
		Id:        role.Id,
		RoleName:  role.RoleName,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: role.DeletedAt,
	}

	return roleResponse
}

func ToRoleMasterResponses(roles []domain.RoleMasterModel, totalRow int64, err error) ([]RoleMasterResponse, int64, error) {
	var roleResponse []RoleMasterResponse
	for _, role := range roles {
		roleResponse = append(roleResponse, ToRoleMasterResponse(role))
	}
	return roleResponse, totalRow, err
}
