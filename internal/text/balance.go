package text

import (
	"encoding/json"
	"fmt"
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/vkbotserver/api"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
)

const (
	addrPattern    = `https://blockchain.info/rawaddr/%s`
	balanceMsgTmpl = `balance is %f BTC`
)

type (
	BalanceRequestHandler struct {
		logger *zap.SugaredLogger
		client api.HTTPClient
		vkApi  *api.Api
	}
)

// NewBalanceRequestHandler creates an object which can check wallet balance
func NewBalanceRequestHandler(
	logger *zap.SugaredLogger,
	client api.HTTPClient,
	vkApi *api.Api,
) *BalanceRequestHandler {
	return &BalanceRequestHandler{
		logger: logger,
		client: client,
		vkApi:  vkApi,
	}
}

// Handle handles user's requests
func (b *BalanceRequestHandler) Handle(req domain.TextRequest) error {
	var (
		err          error
		resp         *http.Response
		request      *http.Request
		dumpResponse []byte
		answer       = domain.NewAddrResponse()
		url          = fmt.Sprintf(addrPattern, req.GetMessage())
	)

	if request, err = http.NewRequest(`GET`, url, nil); err != nil {
		b.
			logger.
			With(
				zap.Error(err),
				zap.String(`url`, url),
			).
			Errorf(`Build balance API request error`)

		return err
	}

	if resp, err = b.client.Do(request); err != nil {
		b.
			logger.
			With(
				zap.Error(err),
				zap.String(`url`, url),
			).
			Error(`Send balance API request error`)
		return err
	}

	if dumpResponse, err = httputil.DumpResponse(resp, true); err != nil {
		b.
			logger.
			With(
				zap.Error(err),
				zap.Int64(`size`, resp.ContentLength),
				zap.Int(`code`, resp.StatusCode),
			).
			Errorf(`Dump API response error`)

		return err
	}

	b.
		logger.
		With(
			zap.String(`address`, req.GetMessage()),
			zap.ByteString(`response`, dumpResponse),
		).
		Info(`Balance API response`)

	if err = json.NewDecoder(resp.Body).Decode(answer); err != nil {
		b.
			logger.
			With(
				zap.Error(err),
				zap.ByteString(`response`, dumpResponse),
			).
			Error(`error while decoding Api response`)

		return err
	}

	return b.vkApi.SendMessage(req.GetPeerId(), fmt.Sprintf(balanceMsgTmpl, answer.BalanceToBTC()))
}
