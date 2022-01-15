package domain

import "net/http"

type (
	Gate interface {
		Send(request *http.Request) (resp *http.Response, err error)
	}
)
