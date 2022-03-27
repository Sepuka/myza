package db

import (
	"github.com/go-pg/pg"
	"github.com/sepuka/vkbotserver/domain"
	"github.com/sepuka/vkbotserver/errors"
)

type UserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (db *UserRepository) GetByExternalId(auth domain.Oauth, id string) (*domain.User, error) {
	var (
		user = &domain.User{}
		err  error
	)

	err = db.
		db.
		Model(user).
		Where(`"user"."o_auth" = ? AND "user"."external_id" = ?`, auth, id).
		Select()

	if err == pg.ErrNoRows {
		return nil, errors.NoUserFound
	}

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

func (db *UserRepository) Update(user *domain.User) error {
	var (
		err error
	)

	_, err = db.
		db.
		Model(user).
		Column(`last_name`, `first_name`, `token`, `updated_at`).
		WherePK().
		Update()

	return err
}
