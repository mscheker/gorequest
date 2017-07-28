package request

import (
	"bytes"
	"errors"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// TODO: Document

type requestBuilder struct {
	auth          AuthorizationMethod
	body          RequestBody
	headers       map[string]string
	method        string
	url           string
	timeout       time.Duration
	checkRedirect func(req *http.Request, via []*http.Request) error
}

func (b *requestBuilder) WithUrl(url string) RequestBuilder {
	b.url = url

	return b
}

// REMARKS: The user/pwd can be provided in the URL when doing Basic Authentication (RFC 1738)
func (b *requestBuilder) WithRFC1738(url string) RequestBuilder {
	u, p, e := splitUserNamePassword(url)

	// TODO: Panic ?
	if e != nil {
		panic(e)
	}

	b.url = url

	return b.WithBasicAuth(u, p)
}

func (b *requestBuilder) WithTextBody(data string) RequestBuilder {
	b.body = newTextBody(data)

	return b
}

func (b *requestBuilder) WithJsonBody(data interface{}) RequestBuilder {
	b.body = newJsonBody(data)

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

func (b *requestBuilder) WithTimeout(timeout time.Duration) RequestBuilder {
	b.timeout = timeout

	return b
}

func (b *requestBuilder) WithCheckRedirect(checkRedirect func(req *http.Request, via []*http.Request) error) RequestBuilder {
	b.checkRedirect = checkRedirect

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

	// REMARKS: Initialize HTTP Client
	client := newHttpClient(b.timeout)

	if b.checkRedirect != nil {
		client.CheckRedirect = b.checkRedirect
	}

	return newRequest(req, client)
}

// ***********************************************
// ********** Private methods/functions **********
// ***********************************************

func (b *requestBuilder) validate() {
	if strings.Trim(b.url, " ") == "" {
		panic(errors.New("URL is required."))
	}

	// REMARKS: Validate method and synchronize the body.
	// REMARKS: For the time being, the Body of a GET request will be ignored. For more information, read below or refer to the HTTP Specification.
	// REMARKS: There is a lot of ambiguity to suggest that most servers won't inspect the body of a GET request. Clients like Postman disable the Body tab when performing a GET request.
	// Ref: https://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html.
	// Section 4.3: A message-body MUST NOT be included in a request if the specification of the request method (section 5.1.1) does not allow sending an entity-body in requests.
	// Section 5.2: The exact resource identified by an Internet request is determined by examining both the Request-URI and the Host header field.
	// Section 9.3: The GET method means retrieve whatever information (in the form of an entity) is identified by the Request-URI.
	// REMARKS: Ignore Body of a DELETE and HEAD requests - RFC2616.
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

// REMARKS: The user/pwd can be provided in the URL when doing Basic Authentication (RFC 1738)
func splitUserNamePassword(url string) (usr, pwd string, err error) {
	reg, err := regexp.Compile("^(http|https|mailto)://")

	if err != nil {
		return "", "", err
	}

	s := reg.ReplaceAllString(url, "")

	if !strings.Contains(s, "@") {
		return "", "", errors.New("No credentials found in URI")
	}

	if reg, err := regexp.Compile("@(.+)"); err != nil {
		return "", "", err
	} else {
		v := reg.ReplaceAllString(s, "")

		c := strings.Split(v, ":")

		if len(c) < 2 {
			return "", "", errors.New("No credentials found in URI")
		}

		return c[0], c[1], nil
	}
}
