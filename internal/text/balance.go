package text

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sepuka/myza/domain"
	button2 "github.com/sepuka/myza/internal/message/button"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
	"time"
)

const (
	addrPattern    = `https://blockchain.info/rawaddr/%s`
	balanceMsgTmpl = `balance is %f BTC (скоро тут появится конвертация в рубли)`
	keyTmpl        = `balance_%d_%s`
)

type (
	BalanceRequestHandler struct {
		logger *zap.SugaredLogger
		client api.HTTPClient
		vkApi  *api.Api
		cache  *redis.Client
	}
)

// NewBalanceRequestHandler creates an object which can check wallet balance
func NewBalanceRequestHandler(
	logger *zap.SugaredLogger,
	client api.HTTPClient,
	vkApi *api.Api,
	cache *redis.Client,
) *BalanceRequestHandler {
	return &BalanceRequestHandler{
		logger: logger,
		client: client,
		vkApi:  vkApi,
		cache:  cache,
	}
}

// Handle handles user's requests
func (b *BalanceRequestHandler) Handle(req domain.TextRequest) error {
	var (
		err          error
		resp         *http.Response
		request      *http.Request
		dumpResponse []byte
		answer       *domain.AddrResponse
		url          = fmt.Sprintf(addrPattern, req.GetMessage())
		keyboard     = button.Keyboard{
			OneTime: true,
			Buttons: button2.ModeChoose(),
		}
	)

	if answer = b.getCache(req); answer == nil {
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

		answer = domain.NewAddrResponse()
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
		b.setCache(req, answer)
	}

	return b.vkApi.SendMessageWithButton(req.GetPeerId(), fmt.Sprintf(balanceMsgTmpl, answer.BalanceToBTC()), keyboard)
}

func (b *BalanceRequestHandler) getCache(req domain.TextRequest) *domain.AddrResponse {
	var (
		key      = fmt.Sprintf(keyTmpl, req.GetPeerId(), req.GetMessage())
		cache    *redis.StringCmd
		value    string
		response domain.AddrResponse
		err      error
	)

	cache = b.cache.Get(b.cache.Context(), key)
	value = cache.Val()
	if len(value) > 0 {
		if err = json.Unmarshal([]byte(value), &response); err == nil {
			return &response
		}
	}

	return nil
}

func (b *BalanceRequestHandler) setCache(req domain.TextRequest, cache *domain.AddrResponse) {
	const (
		ttl = 60 * time.Second
	)

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
