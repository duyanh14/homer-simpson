package models

import "database/sql"

type AddRole struct {
	RoleName        string
	RoleAlias       string
	RoleDescription string
	CreatedBy       string
}

type UpdateRole struct {
	RoleID          int64
	RoleName        string
	RoleAlias       string
	RoleDescription string
	CreatedBy       string
}

type Role struct {
	RoleID          uint32
	RoleName        string
	RoleAlias       sql.NullString
	RoleDescription sql.NullString
	CreatedBy       sql.NullString
	UpdatedBy       sql.NullString
	CreatedDate     string
	IsActive        sql.NullBool
	UpdatedDate     sql.NullTime
}
