package user

import (
	def "github.com/Rahugg/Golang-clean-arch-template/internal/repository"
	"github.com/jmoiron/sqlx"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: nil,
	}
}
