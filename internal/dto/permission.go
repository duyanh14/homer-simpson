package dto

import "time"

type Permission struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Alias       string    `json:"alias"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// DeletedAt   time.Time `json:"deleted_at"`
}
type AddPermissionReqDTO struct {
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ListPermissionReqDTO struct{}
type ListPermissionRespDTO struct{}

type UpdatePermissionReqDTO struct {
	PermissionID uint   `json:"permission_id"`
	Name         string `json:"name"`
	Alias        string `json:"alias"`
	Code         string `json:"code"`
	Description  string `json:"description"`
	CreatedBy    uint   `json:"created_by"`
}
type UpdatePermissionRespDTO struct{}

type DeletePermissionReqDTO struct {
	PermissionID uint `json:"permission_id"`
}
type DeletePermissionRespDTO struct{}
