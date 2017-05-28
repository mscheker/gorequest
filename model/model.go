package gorequest

/**
 * TODO: describe this API, motivation, give examples, etc.
 */

import (
	"bytes"
	"net/http"
)

/**
 *  TODO: describe this interface.
 */
type Request interface {
	Do() Response
}

/**
 *  TODO: describe this interface.
 */
type Response interface {
	Body() []byte
	Response() *http.Response
}

/**
 *  TODO: describe this interface.
 */
type AuthorizationMethod interface {
	Configure(request *http.Request)
}

/**
 *  TODO: describe this interface.
 */
type RequestBody interface {
	ContentType() string
	RawData() *bytes.Buffer
}

/**
 * TODO: describe this interface.
 */
type RequestBuilder interface {
	Build() Request
	WithBasicAuth(user string, password string) RequestBuilder
	WithBearerAuth(token string) RequestBuilder
	WithBody(body RequestBody) RequestBuilder
	WithCustomAuth(auth AuthorizationMethod) RequestBuilder
	WithHeader(name, value string) RequestBuilder
	WithMethod(method string) RequestBuilder
	WithUrl(url string) RequestBuilder
}

/**
 * Defines a constructor type that returns a default RequestBuilder instance.
 */
type RequestBuilderConstructor func() RequestBuilder;
