package gorequest

import (
	model "github.com/demianlessa/gorequest/model"
	"net/http"
	"time"
)

/**
 * This constructor is the entry point into the implementation. 
 */
func NewRequestBuilder() model.RequestBuilder {
	return &requestBuilder{
		auth: newAuthNone(),
		headers: make(map[string]string),
		method: defaultMethod,
	}
}

func getDefaultHttpClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: defaultTimeout,
		}
	}
	return httpClient;
}

func getHttpClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout*time.Second,
	}
}

var httpClient *http.Client
var defaultAuthorization model.AuthorizationMethod = newAuthNone()
var defaultMethod string = "GET"
var defaultTimeout time.Duration = 30*time.Second
