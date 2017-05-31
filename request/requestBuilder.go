package request

import (
	"bytes"
	"errors"
	"net/http"
	"strings"
)

// TODO: Document

type requestBuilder struct {
	auth    AuthorizationMethod
	body    RequestBody
	headers map[string]string
	method  string
	url     string
}

func (b *requestBuilder) WithUrl(url string) RequestBuilder {
	b.url = url

	return b
}

//func (b *requestBuilder) WithBody(body RequestBody) RequestBuilder {
//	b.body = body

//	return b
//}

func (b *requestBuilder) WithTextBody(data string) RequestBuilder {
	b.body = newTextBody(data)

	return b
}

func (b *requestBuilder) WithHeader(key, value string) RequestBuilder {
	b.headers[key] = value

	return b
}

func (b *requestBuilder) WithMethod(method string) RequestBuilder {
	b.method = method

	return b
}

func (b *requestBuilder) WithBasicAuth(username, password string) RequestBuilder {
	b.auth = newAuthBasic(username, password)

	return b
}

func (b *requestBuilder) WithBearerAuth(token string) RequestBuilder {
	b.auth = newAuthBearer(token)

	return b
}

func (b *requestBuilder) Build() Request {
	b.validate()

	var body *bytes.Buffer = &bytes.Buffer{}

	if b.body != nil {
		body = b.body.RawData()
		b.headers["Content-Type"] = b.body.ContentType()
	}

	req, err := http.NewRequest(b.method, b.url, body)

	if err != nil {
		panic(err)
	}

	b.auth.Configure(req)

	for k, v := range b.headers {
		if req.Header.Get(k) != "" {
			continue
		}

		req.Header.Add(k, v)
	}

	return newRequest(req)
}

func (b *requestBuilder) validate() {

	if strings.Trim(b.url, " ") == "" {
		panic(errors.New("URL is required."))
	}

	// validate method and synchronize the body
	switch strings.ToUpper(b.method) {
	case "POST":
		b.method = "POST"
	case "PUT":
		b.method = "PUT"
	case "DELETE":
		b.method = "DELETE"
		b.body = nil
	case "HEAD":
		b.method = "HEAD"
		b.body = nil
	default:
		b.method = "GET"
		b.body = nil
	}
}
