package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/sepuka/myza/domain"
	mocks2 "github.com/sepuka/myza/domain/mocks"
	api2 "github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/mocks"
	"github.com/sepuka/vkbotserver/config"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	addr                = `there_is_a_bitcoin_address`
	balanceCacheKeyTmpl = `balance_%d_%s`
)

func TestBalanceRequestHandler_Handle_GotBalanceWhenCacheIsEmpty(t *testing.T) {
	const (
		peerId           = 1
		blockchainAnswer = `{"hash160":"string","address":"string","n_tx":2,"n_unredeemed":0,"total_received":25000,"total_sent":25000,"final_balance":25000,"txs":[{"hash":"string","ver":1,"vin_sz":1,"vout_sz":1,"size":189,"weight":756,"fee":384,"relayed_by":"0.0.0.0","lock_time":0,"tx_index":1,"double_spend":false,"time":1,"block_index":1,"block_height":1,"inputs":[{"sequence":1,"witness":"","script":"string","index":0,"prev_out":{"spent":true,"script":"string","spending_outpoints":[{"tx_index":1,"n":0}],"tx_index":1,"value":25000,"addr":"string","n":0,"type":0}}],"out":[{"type":0,"spent":false,"value":24616,"spending_outpoints":[],"n":0,"tx_index":1,"script":"string","addr":"string"}],"result":-25000,"balance":0},{"hash":"string","ver":2,"vin_sz":1,"vout_sz":2,"size":226,"weight":904,"fee":452,"relayed_by":"0.0.0.0","lock_time":710953,"tx_index":1,"double_spend":false,"time":1,"block_index":1,"block_height":1,"inputs":[{"sequence":1,"witness":"","script":"string","index":0,"prev_out":{"spent":true,"script":"string","spending_outpoints":[{"tx_index":1,"n":0}],"tx_index":1,"value":52218,"addr":"string","n":0,"type":0}}],"out":[{"type":0,"spent":true,"value":25000,"spending_outpoints":[{"tx_index":1,"n":0}],"n":0,"tx_index":1,"script":"string","addr":"string"},{"type":0,"spent":false,"value":26766,"spending_outpoints":[],"n":1,"tx_index":1,"script":"string","addr":"string"}],"result":25000,"balance":25000}]}`
		expectedMsg      = `?????? ???????????? 0.000250 BTC (1 ??????)`
	)

	var (
		client                = mocks.HTTPClient{}
		cache                 = &mocks2.Cache{}
		cacheResponse         = &mocks2.CacheResponse{}
		converterMock         = &mocks2.ExchangeRateConverter{}
		blockchainHttpRequest *http.Request
		vkHttpRequest         *http.Request
		blockchainApiResponse *http.Response
		vkApiResponse         *http.Response
		handler               domain.TextHandler
		logger                = zap.NewNop().Sugar()
		userRequest           = domain.NewTextRequest(peerId, addr)
		vkAnswer, _           = json.Marshal(&api2.Response{})
		wallet                = &domain.Wallet{
			Hash160:       `string`,
			Address:       `string`,
			NTx:           2,
			TotalReceived: 25000,
			TotalSent:     25000,
			FinalBalance:  25000,
		}
		cfg = config.Config{
			Api: config.Api{
				Token: `some_token`,
			},
		}
		rnd      = mocks.Rnder{}
		api      *api2.Api
		err      error
		endpoint = fmt.Sprintf(`%s/%s`, api2.Endpoint, `messages.send`)
		rndId    = int64(777)
		payload  = api2.OutcomeMessage{
			AccessToken: cfg.Api.Token,
			ApiVersion:  api2.Version,
			RandomId:    rndId,
			PeerId:      peerId,
			Message:     expectedMsg,
			Keyboard:    `{"one_time":true,"buttons":[[{"action":{"type":"text","label":"?????????????? ???? ??????????","payload":"{\"command\":\"withdraw\",\"id\":\"\"}"},"color":"primary"}]]}`,
		}
	)

	rnd.On(`Rnd`).Return(rndId)

	balanceCacheKey := fmt.Sprintf(balanceCacheKeyTmpl, peerId, addr)
	cacheResponse.On(`Val`).Return(``)
	cacheResponse.On(`Err`).Return(nil)
	cache.On(`Context`).Return(nil)
	cache.On(`Get`, cache.Context(), balanceCacheKey).Return(cacheResponse)
	cache.On(`Set`, cache.Context(), balanceCacheKey, wallet, ttl).Return(cacheResponse)

	converterMock.On(`Convert`, wallet, domain.Rub).Return(1.0, nil)

	blockchainApiResponse = &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(blockchainAnswer))),
	}
	blockchainHttpRequest, _ = http.NewRequest(`GET`, fmt.Sprintf(addrPattern, addr), nil)
	client.On(`Do`, blockchainHttpRequest).Once().Return(blockchainApiResponse, nil)

	vkApiResponse = &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(vkAnswer)),
	}
	params, _ := query.Values(payload)
	vkHttpRequest, _ = http.NewRequest(`POST`, fmt.Sprintf(`%s?%s`, endpoint, params.Encode()), nil)
	client.On(`Do`, vkHttpRequest).Once().Return(vkApiResponse, nil)

	api = api2.NewApi(logger, cfg, &client, &rnd)

	handler = NewBalanceRequestHandler(logger, &client, api, cache, converterMock)
	err = handler.Handle(userRequest)
	if err != nil {
		t.Errorf(`Got unexpecred error %v`, err)
	}
}

func TestBalanceRequestHandler_Handle_WhenHttpClientFailed(t *testing.T) {
	const (
		peerId = 1
	)
	var (
		client        = mocks.HTTPClient{}
		cache         = &mocks2.Cache{}
		cacheResponse = &mocks2.CacheResponse{}
		converterMock = &mocks2.ExchangeRateConverter{}
		httpRequest   *http.Request
		logger        = zap.NewNop().Sugar()
		handler       domain.TextHandler
		userRequest   = domain.NewTextRequest(peerId, addr)
		cfg           = config.Config{}
		rnd           = mocks.Rnder{}
		api           *api2.Api
		err           error
		expectedErr   = errors.New(`Something went wrong`)
	)

	httpRequest, _ = http.NewRequest(`GET`, fmt.Sprintf(addrPattern, addr), nil)
	client.On(`Do`, httpRequest).Return(nil, expectedErr)

	api = api2.NewApi(logger, cfg, &client, &rnd)

	balanceCacheKey := fmt.Sprintf(balanceCacheKeyTmpl, peerId, addr)
	cacheResponse.On(`Val`).Return(``)
	cache.On(`Context`).Return(nil)
	cache.On(`Get`, cache.Context(), balanceCacheKey).Return(cacheResponse)

	handler = NewBalanceRequestHandler(logger, &client, api, cache, converterMock)
	err = handler.Handle(userRequest)
	if err != expectedErr {
		t.Error(`Expected error was not occured`)
	}
}
