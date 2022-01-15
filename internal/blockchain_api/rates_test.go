package blockchain_api

import (
	"encoding/json"
	"github.com/sepuka/myza/domain"
	mocks2 "github.com/sepuka/myza/domain/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

var (
	rates = domain.Rates{
		domain.Rub: &domain.Rate{
			M:      3274646.11,
			Last:   3274646.11,
			Sell:   3274646.11,
			Buy:    3274646.11,
			Symbol: `RUB`,
		},
	}
)

func TestExchangeConverter_Convert_WithCache(t *testing.T) {
	var (
		logger        = zap.NewNop().Sugar()
		cache         = &mocks2.Cache{}
		cacheResponse = &mocks2.CacheResponse{}
		gate          = &mocks2.Gate{}
		wallet        = &domain.Wallet{
			Hash160:       `string`,
			Address:       `string`,
			NTx:           2,
			TotalReceived: 25000,
			TotalSent:     25000,
			FinalBalance:  25000,
		}
		converter   *ExchangeConverter
		cachedRates []byte
		result      float64
		err         error
	)

	cachedRates, err = json.Marshal(rates)
	assert.Nil(t, err)

	cache.On(`Context`).Return(nil)
	cacheResponse.On(`Val`).Return(string(cachedRates))
	cache.On(`Get`, cache.Context(), exchangeRateKeyTmpl).Return(cacheResponse, nil)

	converter = NewExchangeConverter(logger, cache, gate)
	result, err = converter.Convert(wallet, domain.Rub)
	assert.Nil(t, err)
	assert.Equal(t, 818.6615275, result)
}
