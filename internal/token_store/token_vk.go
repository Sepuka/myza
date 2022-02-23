package token_store

import (
	"encoding/json"
	"fmt"
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/errors"
	domain2 "github.com/sepuka/vkbotserver/domain"
	"go.uber.org/zap"
	"time"
)

const (
	vkTokenKeyTmpl = `myza_vk_auth_token_%s`
)

type (
	tokenStore struct {
		cache  domain.Cache
		logger *zap.SugaredLogger
	}
)

func NewTokenStore(
	cache domain.Cache,
	logger *zap.SugaredLogger,
) *tokenStore {
	return &tokenStore{
		cache:  cache,
		logger: logger,
	}
}

func (s *tokenStore) GetToken(authCookie string) (authToken interface{}, err error) {
	var (
		value string
		key   = fmt.Sprintf(vkTokenKeyTmpl, authCookie)
		token domain2.OauthVkTokenResponse
	)

	value = s.cache.Get(s.cache.Context(), key).Val()
	if len(value) > 0 {
		if err = json.Unmarshal([]byte(value), &token); err == nil {
			return &token, nil
		}
	}

	return nil, errors.NewOauthTokenError(`could not fetch VK token from store`, nil)
}

func (s *tokenStore) SetToken(authToken interface{}) (cookie string, err error) {
	var (
		cookieValue = `this_is_a_secret`
		key         = fmt.Sprintf(vkTokenKeyTmpl, cookieValue)
		cache       []byte
		ttl         = 600 * time.Second
	)

	if cache, err = json.Marshal(authToken); err != nil {
		return ``, errors.NewOauthTokenError(`could not save VK token to store`, err)
	}

	if err = s.cache.Set(s.cache.Context(), key, cache, ttl).Err(); err != nil {
		s.
			logger.
			With(
				zap.Error(err),
				zap.String(`key`, key),
			).
			Error(`Could not save auth token to store`)
		return ``, errors.NewOauthTokenError(`could not save VK token to store`, err)
	}

	return cookieValue, nil
}
