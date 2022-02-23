package db

import (
	"github.com/go-pg/pg"
	"github.com/sepuka/vkbotserver/domain"
)

type UserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (db *UserRepository) GetByExternalId(auth domain.Oauth, id string) (*domain.User, error) {
	var (
		user *domain.User
		err  error
	)

	err = db.
		db.
		Model(user).
		Where(`"user"."o_auth" = ? AND "user"."external_id" = ?`, auth, id).
		Select()

	return user, err
}

func (db *UserRepository) Create(user *domain.User) error {
	var (
		err error
	)

	_, err = db.
		db.
		Model(user).
		Insert()

	return err
}
