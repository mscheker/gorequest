package request

import (
	"bytes"
	"net/http"
)

// TODO: Document interfaces

type Request interface {
	Do() Response
	getUnderlyingRequest() *http.Request
}

type Response interface {
	Body() []byte
	Response() *http.Response
}

type AuthorizationMethod interface {
	Configure(request *http.Request)
}

type RequestBody interface {
	ContentType() string
	RawData() *bytes.Buffer
}

type RequestBuilder interface {
	Build() Request
	WithBody(body RequestBody) RequestBuilder
	WithHeader(name, value string) RequestBuilder
	WithMethod(method string) RequestBuilder
	WithUrl(url string) RequestBuilder
}

type RequestBuilderConstructor func() RequestBuilder
