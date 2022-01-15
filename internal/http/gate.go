package http

import (
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
)

type (
	Gate struct {
		client *http.Client
		logger *zap.SugaredLogger
	}
)

func NewGate(
	client *http.Client,
	logger *zap.SugaredLogger,
) *Gate {
	return &Gate{
		client: client,
		logger: logger,
	}
}

func (g *Gate) Send(request *http.Request) (resp *http.Response, err error) {
	if resp, err = g.client.Do(request); err != nil {
		g.
			logger.
			With(
				zap.Error(err),
				zap.String(`url`, request.URL.String()),
			).
			Error(`Could not send HTTP request`)

		return nil, err
	}

	var dumpResponse []byte
	if dumpResponse, err = httputil.DumpResponse(resp, true); err != nil {
		g.
			logger.
			With(
				zap.Error(err),
				zap.Int64(`size`, resp.ContentLength),
				zap.Int(`code`, resp.StatusCode),
			).
			Errorf(`Dump API response error`)

		return nil, err
	}

	g.
		logger.
		With(
			zap.String(`url`, request.URL.String()),
			zap.ByteString(`response`, dumpResponse),
		).
		Info(`API response`)

	return resp, nil
}
