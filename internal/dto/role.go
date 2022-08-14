package dto

import "time"

type Role struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Alias       string    `json:"alias"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
type AddRoleReqDTO struct {
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ListRoleReqDTO struct {
	IsActive bool
}
type ListRoleRespDTO struct{}

type UpdateRoleReqDTO struct {
	RoleID      uint   `json:"role_id"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Code        string `json:"code"`
	Description string `json:"description"`
	CreatedBy   uint   `json:"created_by"`
}
type UpdateRoleRespDTO struct{}

type DeleteRoleReqDTO struct {
	RoleID uint `json:"role_id"`
}
type DeleteRoleRespDTO struct{}
