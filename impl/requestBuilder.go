package gorequest

import (
	"bytes"
	model "gorequest/model"
	"errors"
	"net/http"
	"strings"
)

/****************************************************
 * model.RequestBuilder implementation
 ****************************************************/

type requestBuilder struct {
	auth    	model.AuthorizationMethod
	body    	model.RequestBody
	headers 	map[string]string
	method  	string
	url     	string
}

func (b *requestBuilder) Build() model.Request {

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

	// delegate the authorization configuration
	b.auth.Configure(req)

	// set request headers
	for k, v := range b.headers {
		// do not override headers set previously
		if req.Header.Get(k) != "" {
			continue
		}
		req.Header.Add(k, v)
	}

	return newRequest(req)
}

func (b *requestBuilder) WithBasicAuth(user string, password string) model.RequestBuilder {
	b.auth = newAuthBasic(user, password)
	return b
}

func (b *requestBuilder) WithBearerAuth(token string) model.RequestBuilder {
	b.auth = newAuthBearer(token)
	return b
}

func (b *requestBuilder) WithBody(body model.RequestBody) model.RequestBuilder {
	b.body = body
	return b
}

func (b *requestBuilder) WithCustomAuth(auth model.AuthorizationMethod) model.RequestBuilder {
	if auth != nil {
		b.auth = auth
	} else {
		b.auth = newAuthNone()
	}
	return b
}

func (b *requestBuilder) WithHeader(name, value string) model.RequestBuilder {
	b.headers[name] = value
	return b
}

func (b *requestBuilder) WithMethod(method string) model.RequestBuilder {
	b.method = method
	return b
}

func (b *requestBuilder) WithUrl(url string) model.RequestBuilder {
	b.url = url
	return b
}

func (b *requestBuilder) validate() {

	if strings.Trim(b.url, " ") == "" {
		panic(errors.New("URL is required"))
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
