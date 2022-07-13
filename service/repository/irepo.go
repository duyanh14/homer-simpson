package repository

import (
	"database/sql"

	"golang-course/service/repository/mysql"
)

type IRepo interface {
	NewRoleRepo() mysql.IRoleRepo
}

type repo struct {
	db *sql.DB
}

func (r repo) NewRoleRepo() mysql.IRoleRepo {
	return mysql.NewRoleRepo(r.db)
}

func NewRepo(
	db *sql.DB,
) IRepo {
	return &repo{
		db: db,
	}
}
