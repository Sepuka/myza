package text

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/sepuka/myza/domain"
	api2 "github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/mocks"
	"github.com/sepuka/vkbotserver/config"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	addr = `there_is_a_bitcoin_address`
)

func TestBalanceRequestHandler_Handle_GotBalance(t *testing.T) {
	const (
		balance     = 25000
		expectedMsg = `balance is 0.000250 BTC`
		peerId      = 1
	)

	var (
		client                = mocks.HTTPClient{}
		blockchainHttpRequest *http.Request
		vkHttpRequest         *http.Request
		blockchainApiResponse *http.Response
		vkApiResponse         *http.Response
		handler               domain.TextHandler
		logger                = zap.NewNop().Sugar()
		userRequest           = domain.NewTextRequest(peerId, addr)
		blockchainAnswer, _   = json.Marshal(&domain.AddrResponse{FinalBalance: balance})
		vkAnswer, _           = json.Marshal(&api2.Response{})
		cfg                   = config.Config{
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
		}
	)
	rnd.On(`Rnd`).Return(rndId)

	blockchainApiResponse = &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(blockchainAnswer)),
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

	handler = NewBalanceRequestHandler(logger, &client, api)
	err = handler.Handle(userRequest)
	if err != nil {
		t.Errorf(`Got unexpecred error %v`, err)
	}
}

func TestBalanceRequestHandler_Handle_WhenHttpClientFailed(t *testing.T) {
	var (
		client      = mocks.HTTPClient{}
		httpRequest *http.Request
		logger      = zap.NewNop().Sugar()
		handler     domain.TextHandler
		userRequest = domain.NewTextRequest(0, addr)
		cfg         = config.Config{}
		rnd         = mocks.Rnder{}
		api         *api2.Api
		err         error
		expectedErr = errors.New(`Something went wrong`)
	)
	httpRequest, _ = http.NewRequest(`GET`, fmt.Sprintf(addrPattern, addr), nil)
	client.On(`Do`, httpRequest).Return(nil, expectedErr)

	api = api2.NewApi(logger, cfg, &client, &rnd)

	handler = NewBalanceRequestHandler(logger, &client, api)
	err = handler.Handle(userRequest)
	if err != expectedErr {
		t.Error(`Expected error was not occured`)
	}
}
