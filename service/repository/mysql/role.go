package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"golang-course/service/models"
)

type IRoleRepo interface {
	GetRole(ctx context.Context) (listRole []*models.Role, err error)
	AddRole(ctx context.Context, role models.AddRole) error
	DeleteRole(ctx context.Context, id int64) error
	UpdateRole(ctx context.Context, role models.UpdateRole) error
}

type roleRepo struct {
	db *sql.DB
}

func NewRoleRepo(db *sql.DB) IRoleRepo {
	return &roleRepo{
		db: db,
	}
}

func (r *roleRepo) GetRole(ctx context.Context) (listRole []*models.Role, err error) {
	stmt := `select id, role_name, role_alias, role_description, is_active, created_by from role`
	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		role := &models.Role{}
		err := rows.Scan(&role.RoleID, &role.RoleName, &role.RoleAlias, &role.RoleDescription, &role.IsActive, &role.CreatedBy)
		if err != nil {
			return nil, err
		}
		listRole = append(listRole, role)
	}
	return listRole, nil
}

func (r *roleRepo) AddRole(ctx context.Context, role models.AddRole) error {
	stmt := `insert into role(role_name, role_alias, role_description, created_by)
				values (?, ?, ?, ?)`
	result, err := r.db.Exec(stmt, role.RoleName, role.RoleAlias, role.RoleDescription, role.CreatedBy)
	if err != nil {
		return err
	}
	if num, err := result.RowsAffected(); num == 0 || err != nil {
		return err
	}
	return nil
}

func (r *roleRepo) UpdateRole(ctx context.Context, req models.UpdateRole) error {
	stmt := `update role set role_name = ?, role_alias = ?, role_description = ? where id = ?`
	result, err := r.db.Exec(stmt, req.RoleName, req.RoleAlias, req.RoleDescription, req.RoleID)
	if err != nil {
		return err
	}
	numEffect, _ := result.RowsAffected()
	if numEffect == 0 {
		return fmt.Errorf("effected empty")
	}
	return nil
}

func (r *roleRepo) DeleteRole(ctx context.Context, roleID int64) error {
	stmt := `delete from role where id = ?`
	result, err := r.db.Exec(stmt, roleID)
	if err != nil {
		return err
	}
	numEffect, _ := result.RowsAffected()
	if numEffect == 0 {
		return fmt.Errorf("effected empty")
	}
	return nil
}
