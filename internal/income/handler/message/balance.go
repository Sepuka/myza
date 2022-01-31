package message

import (
	"encoding/json"
	"fmt"
	"github.com/sepuka/myza/domain"
	button2 "github.com/sepuka/myza/internal/income/button"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
	"time"
)

const (
	addrPattern    = `https://blockchain.info/rawaddr/%s`
	balanceMsgTmpl = `Ваш баланс %f BTC (%d руб)`
	keyTmpl        = `balance_%d_%s`
	ttl            = 60 * time.Second
)

type (
	BalanceRequestHandler struct {
		logger    *zap.SugaredLogger
		client    api.HTTPClient
		vkApi     *api.Api
		cache     domain.Cache
		converter domain.ExchangeRateConverter
	}
)

// NewBalanceRequestHandler creates an object which can check wallet balance
func NewBalanceRequestHandler(
	logger *zap.SugaredLogger,
	client api.HTTPClient,
	vkApi *api.Api,
	cache domain.Cache,
	converter domain.ExchangeRateConverter,
) *BalanceRequestHandler {
	return &BalanceRequestHandler{
		logger:    logger,
		client:    client,
		vkApi:     vkApi,
		cache:     cache,
		converter: converter,
	}
}

// Handle handles user's requests
func (b *BalanceRequestHandler) Handle(req domain.TextRequest) error {
	var (
		err          error
		resp         *http.Response
		request      *http.Request
		dumpResponse []byte
		wallet       *domain.Wallet
		amount       float64
		answer       string
		url          = fmt.Sprintf(addrPattern, req.GetMessage())
		keyboard     = button.Keyboard{
			OneTime: true,
			Buttons: button2.Buttons(),
		}
	)

	// TODO replace it feature toggling
	if req.GetPeerId() == 557404793 {
		keyboard.Buttons = button2.ButtonsWithAddr()
	}

	if wallet = b.getCache(req); wallet == nil {
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

		wallet = domain.NewWallet()
		if err = json.NewDecoder(resp.Body).Decode(wallet); err != nil {
			b.
				logger.
				With(
					zap.Error(err),
					zap.ByteString(`response`, dumpResponse),
				).
				Error(`error while decoding Api response`)

			return err
		}
		b.setCache(req, wallet)
	}

	amount, err = b.converter.Convert(wallet, req.FiatCurrency)
	answer = fmt.Sprintf(balanceMsgTmpl, wallet.BalanceToBTC(), int(amount))

	return b.vkApi.SendMessageWithButton(req.GetPeerId(), answer, keyboard)
}

func (b *BalanceRequestHandler) getCache(req domain.TextRequest) *domain.Wallet {
	var (
		key      = fmt.Sprintf(keyTmpl, req.GetPeerId(), req.GetMessage())
		value    string
		response domain.Wallet
		err      error
	)

	value = b.cache.Get(b.cache.Context(), key).Val()
	if len(value) > 0 {
		if err = json.Unmarshal([]byte(value), &response); err == nil {
			return &response
		}
	}

	return nil
}

func (b *BalanceRequestHandler) setCache(req domain.TextRequest, cache *domain.Wallet) {
	var (
		key = fmt.Sprintf(keyTmpl, req.GetPeerId(), req.GetMessage())
		err error
	)

	if err = b.cache.Set(b.cache.Context(), key, cache, ttl).Err(); err != nil {
		b.
			logger.
			With(
				zap.Error(err),
				zap.String(`key`, key),
			).
			Error(`Could not save balance to cache`)
	}
}
