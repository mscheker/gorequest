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
	}
}

func getDefaultHttpClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: defaultTimeout,
		}
	}

	return httpClient
}

func getHttpClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout * time.Second,
	}
}

var httpClient *http.Client
var defaultAuthorization AuthorizationMethod = newAuthNone()
var defaultMethod string = "GET"
var defaultTimeout time.Duration = 30 * time.Second
