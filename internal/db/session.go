package db

import (
	"github.com/go-pg/pg"
	domain2 "github.com/sepuka/myza/domain"
	"github.com/sepuka/vkbotserver/domain"
	"time"
)

type (
	SessionsRepository struct {
		db *pg.DB
	}
)

func NewSessionsRepository(db *pg.DB) *SessionsRepository {
	return &SessionsRepository{
		db: db,
	}
}

func (s *SessionsRepository) Create(user *domain.User, token string) error {
	var (
		err     error
		session = &domain2.Session{
			UserId:   user.UserId,
			Token:    token,
			OAuth:    user.OAuth,
			Datetime: time.Now(),
		}
	)

	_, err = s.
		db.
		Model(session).
		Insert()

	return err
}
