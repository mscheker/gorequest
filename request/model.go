package request

import (
	"bytes"
	"net/http"
	"time"
)

type authScheme int

const (
	AUTH_NONE authScheme = iota
	AUTH_BASIC
	AUTH_BEARER
	AUTH_DIGEST
)

// TODO: Document interfaces

type Request interface {
	Do() Response
	getUnderlyingRequest() *http.Request
	getUnderlyingHttpClient() *http.Client
}

type Response interface {
	Body() []byte
	Response() *http.Response
}

type AuthorizationMethod interface {
	Configure(request *http.Request)
	getScheme() authScheme
}

type RequestBody interface {
	ContentType() string
	RawData() *bytes.Buffer
}

type RequestBuilder interface {
	Build() Request
	WithTextBody(data string) RequestBuilder
	WithJsonBody(data interface{}) RequestBuilder
	WithRFC1738(url string) RequestBuilder
	WithHeader(name, value string) RequestBuilder
	WithMethod(method string) RequestBuilder
	WithUrl(url string) RequestBuilder
	WithBasicAuth(username, password string) RequestBuilder
	WithBearerAuth(token string) RequestBuilder
	WithDigestAuth(username, password string) RequestBuilder
	WithTimeout(timeout time.Duration) RequestBuilder
	WithCheckRedirect(checkRedirect func(req *http.Request, via []*http.Request) error) RequestBuilder
}

type RequestBuilderConstructor func() RequestBuilder
