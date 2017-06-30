package request

import (
	"net/http"
	"time"
)

// TODO: Document c_tors

func NewRequestBuilder() RequestBuilder {
	return &requestBuilder{
		auth:    newAuthNone(),
		headers: make(map[string]string),
		method:  defaultMethod,
		timeout: defaultTimeout,
	}
}

func newHttpClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}

var defaultAuthorization AuthorizationMethod = newAuthNone()
var defaultMethod string = "GET"
var defaultTimeout time.Duration = 30 * time.Second
