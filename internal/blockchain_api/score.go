package blockchain_api

import (
	"encoding/json"
	"fmt"
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/errors"
	"github.com/sepuka/vkbotserver/api"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

const (
	keyTmpl     = `balance_%s_%s`
	addrPattern = `https://blockchain.info/rawaddr/%s`
	ttl         = 60 * time.Second
)

type (
	Score struct {
		logger *zap.SugaredLogger
		client api.HTTPClient
		cache  domain.Cache
	}
)

func NewScoreFetcher(
	logger *zap.SugaredLogger,
	client api.HTTPClient,
	cache domain.Cache,
) *Score {
	return &Score{
		logger: logger,
		client: client,
		cache:  cache,
	}
}

func (bc *Score) GetBalance(crypto *domain.Crypto) (float64, error) {
	var (
		err          error
		resp         *http.Response
		request      *http.Request
		dumpResponse []byte
		wallet       *domain.Wallet
		url          = fmt.Sprintf(addrPattern, crypto.Address)
	)

	if wallet = bc.getCache(*crypto); wallet == nil {
		if request, err = http.NewRequest(`GET`, url, nil); err != nil {
			bc.
				logger.
				With(
					zap.Error(err),
					zap.String(`url`, url),
				).
				Errorf(`Build balance API request error`)

			return 0, err
		}

		if resp, err = bc.client.Do(request); err != nil || resp.StatusCode != 200 {
			body, _ := ioutil.ReadAll(resp.Body)
			bc.
				logger.
				With(
					zap.Error(err),
					zap.String(`url`, url),
					zap.Int(`code`, resp.StatusCode),
					zap.String(`status`, resp.Status),
					zap.ByteString(`body`, body),
				).
				Error(`Send balance API request error`)

			if err != nil {
				return 0, err
			} else {
				return 0, errors.NewBlockchainBalanceError(string(body))
			}
		}

		if dumpResponse, err = httputil.DumpResponse(resp, true); err != nil {
			bc.
				logger.
				With(
					zap.Error(err),
					zap.Int64(`size`, resp.ContentLength),
					zap.Int(`code`, resp.StatusCode),
				).
				Errorf(`Dump API response error`)

			return 0, err
		}

		bc.
			logger.
			With(
				zap.String(`address`, crypto.Address),
				zap.String(`currency`, string(crypto.Currency)),
				zap.ByteString(`response`, dumpResponse),
			).
			Info(`Balance API response`)

		wallet = domain.NewWallet()
		if err = json.NewDecoder(resp.Body).Decode(wallet); err != nil {
			bc.
				logger.
				With(
					zap.Error(err),
					zap.ByteString(`response`, dumpResponse),
				).
				Error(`error while decoding Api response`)

			return 0, err
		}
		bc.setCache(*crypto, wallet)
	}

	return wallet.BalanceToBTC(), err
}

func (bc *Score) getCache(crypto domain.Crypto) *domain.Wallet {
	var (
		key      = fmt.Sprintf(keyTmpl, crypto.Currency, crypto.Address)
		value    string
		response domain.Wallet
		err      error
	)

	value = bc.cache.Get(bc.cache.Context(), key).Val()
	if len(value) > 0 {
		if err = json.Unmarshal([]byte(value), &response); err == nil {
			return &response
		}
	}

	return nil
}

func (bc *Score) setCache(crypto domain.Crypto, cache *domain.Wallet) {
	var (
		key = fmt.Sprintf(keyTmpl, crypto.Currency, crypto.Address)
		err error
	)

	if err = bc.cache.Set(bc.cache.Context(), key, cache, ttl).Err(); err != nil {
		bc.
			logger.
			With(
				zap.Error(err),
				zap.String(`key`, key),
			).
			Error(`Could not save balance to cache`)
	}
}
