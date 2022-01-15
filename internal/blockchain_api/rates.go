package blockchain_api

import (
	"encoding/json"
	"fmt"
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/errors"
	http2 "github.com/sepuka/myza/internal/http"
	"go.uber.org/zap"
	"net/http"
)

const (
	keyTmpl   = `exchange_rate`
	urlTicker = `https://blockchain.info/ticker`
)

type (
	ExchangeConverter struct {
		logger *zap.SugaredLogger
		cache  domain.Cache
		gate   *http2.Gate
	}
)

func NewExchangeConverter(
	log *zap.SugaredLogger,
	cache domain.Cache,
	gate *http2.Gate,
) *ExchangeConverter {
	return &ExchangeConverter{
		logger: log,
		cache:  cache,
		gate:   gate,
	}
}

func (c *ExchangeConverter) Convert(wallet *domain.Wallet, fiat domain.FiatCurrency) (float64, error) {
	var (
		rate *domain.Rate
		err  error
	)
	if rate, err = c.getRate(fiat); err != nil {
		return 0, nil
	}

	return wallet.FinalBalance * rate.Last, nil
}

func (c *ExchangeConverter) getRate(currency domain.FiatCurrency) (*domain.Rate, error) {
	var (
		rate    *domain.Rate
		rates   domain.Rates
		err     error
		resp    *http.Response
		request *http.Request
		ok      bool
	)

	if rates = c.getCache(); rates == nil {
		if request, err = http.NewRequest(`GET`, urlTicker, nil); err != nil {
			c.
				logger.
				With(
					zap.Error(err),
					zap.String(`url`, urlTicker),
				).
				Errorf(`Build exchange rate API request error`)

			return nil, err
		}

		if resp, err = c.gate.Send(request); err != nil {
			c.
				logger.
				Error(`Could not get exchange rates`)
			return nil, err
		}

		rates = domain.Rates{}
		if err = json.NewDecoder(resp.Body).Decode(&rates); err != nil {
			c.
				logger.
				With(
					zap.Error(err),
				).
				Error(`Error while decoding Api response`)

			return nil, err
		}
	}

	if rate, ok = rates[currency]; ok {
		return rate, nil
	}

	return nil, errors.NewUnknownCurrencyError(`could not find fiat currency rate`, nil)
}

func (c *ExchangeConverter) getCache() domain.Rates {
	var (
		key   = fmt.Sprintf(keyTmpl)
		value string
		rate  domain.Rates
		err   error
	)

	value = c.cache.Get(c.cache.Context(), key).Val()
	if len(value) > 0 {
		if err = json.Unmarshal([]byte(value), &rate); err == nil {
			return rate
		}
	}

	return nil
}
