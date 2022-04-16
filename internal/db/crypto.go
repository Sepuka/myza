package db

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/sepuka/myza/domain"
	domain2 "github.com/sepuka/vkbotserver/domain"
	"go.uber.org/zap"
	"time"
)

type CryptoRepository struct {
	db     *pg.DB
	logger *zap.SugaredLogger
}

func NewCryptoRepository(db *pg.DB, log *zap.SugaredLogger) *CryptoRepository {
	return &CryptoRepository{
		db:     db,
		logger: log,
	}
}

func (r *CryptoRepository) Get(user *domain2.User, currency domain.CryptoCurrency) *domain.Crypto {
	var (
		crypto = &domain.Crypto{}
		err    error
	)

	err = r.
		db.
		Model(crypto).
		Where(`user_id = ? AND currency = ?`, user.UserId, currency).
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil
		}

		r.
			logger.
			With(
				zap.Int(`user`, user.UserId),
				zap.String(`currency`, string(currency)),
				zap.Error(err),
			).
			Error(`Error while searching an crypto address`)

		return nil
	}

	return crypto
}

func (r *CryptoRepository) AssignAddress(model *domain.Crypto, address domain.Address) error {
	var (
		err    error
		result orm.Result
	)

	result, err = r.
		db.
		Model(model).
		Insert()

	if err != nil {
		return err
	}

	r.
		logger.
		With(
			zap.Uint32(`user`, model.UserId),
			zap.String(`address`, address.String()),
			zap.String(`currency`, string(model.Currency)),
			zap.Int(`affected rows`, result.RowsAffected()),
		).
		Debug(`crypto address was assigned`)

	return nil
}

func (r *CryptoRepository) UpdateBalance(model *domain.Crypto) error {
	var (
		err error
	)

	model.UpdatedAt = time.Now()
	_, err = r.
		db.
		Model(model).
		Column(`balance`, `fiat`, `updated_at`).
		Where(`currency = ? AND address = ?`, model.Currency, model.Address).
		Update()

	return err
}

func (r *CryptoRepository) FindOutdated(date time.Time, limit int) ([]*domain.Crypto, error) {
	var (
		cryptos = make([]*domain.Crypto, limit)
		err     error
	)

	err = r.
		db.
		Model(&cryptos).
		ColumnExpr(`crypto.*`).
		Join(`JOIN users u ON u.user_id = crypto.user_id`).
		Where(`u.active = ? AND (crypto.updated_at < ? OR crypto.updated_at IS NULL)`, true, date).
		Limit(limit).
		Select()

	return cryptos, err
}
