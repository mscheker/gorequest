package request

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var instance *Request

type auth struct {
	Username string
	Password string
	Bearer   string
}

type Option struct {
	Url     string
	Headers map[string]string
	Auth    *auth
	Body    interface{}
}

type Request struct {
	client  *http.Client
	Timeout time.Duration
}

func NewAuth(username, password, bearer string) *auth {
	return &auth{
		Username: username,
		Password: password,
		Bearer:   bearer,
	}
}

func New() *Request {
	r := new(Request)

	r.Timeout = 30 * time.Second

	r.client = &http.Client{
		Timeout: r.Timeout,
	}

	return r
}

func NewRequest(url string) (*http.Response, []byte, error) {
	o := &Option{
		Url: url,
	}

	return Get(o)
}

func (r *Request) Post(o *Option) (*http.Response, []byte, error) {
	return r.doRequest("POST", o)
}

func Post(o *Option) (*http.Response, []byte, error) {
	return getInstance().Post(o)
}

func (r *Request) Put(o *Option) (*http.Response, []byte, error) {
	return r.doRequest("PUT", o)
}

func Put(o *Option) (*http.Response, []byte, error) {
	return getInstance().Put(o)
}

func (r *Request) Get(o *Option) (*http.Response, []byte, error) {
	return r.doRequest("GET", o)
}

func Get(o *Option) (*http.Response, []byte, error) {
	return getInstance().Get(o)
}

func (r *Request) Delete(o *Option) (*http.Response, []byte, error) {
	// REMARKS: Ignore Body - RFC2616
	o.Body = nil

	return r.doRequest("DELETE", o)
}

func Delete(o *Option) (*http.Response, []byte, error) {
	return getInstance().Delete(o)
}

// ********** Private methods/functions **********
// REMARKS: Used internally by non-instance methods
func getInstance() *Request {
	if instance == nil {
		instance = New()
	}

	return instance
}

// REMARKS: The user/pwd can be provided in the URL when doing Basic Authentication (RFC 1738)
func splitUserNamePassword(u string) (usr, pwd string, err error) {
	reg, err := regexp.Compile("^(http|https|mailto)://")

	if err != nil {
		return "", "", err
	}

	s := reg.ReplaceAllString(u, "")

	if reg, err := regexp.Compile("@(.+)"); err != nil {
		return "", "", err
	} else {
		v := reg.ReplaceAllString(s, "")

		c := strings.Split(v, ":")

		if len(c) < 1 {
			return "", "", errors.New("No credentials found in URI")
		}

		return c[0], c[1], nil
	}
}

// REMARKS: Returns a buffer with the body of the request - Content-Type header is set accordingly
func getRequestBody(o *Option) *bytes.Buffer {
	b := reflect.Indirect(reflect.ValueOf(o.Body))
	buff := make([]byte, 0)
	body := new(bytes.Buffer)
	contentType := ""

	switch b.Kind() {
	case reflect.String:
		// REMARKS: This takes care of a JSON serialized string
		buff = []byte(b.String())
		body = bytes.NewBuffer(buff)

		// TODO: Need to set headers accordingly
		contentType = "text/plain"
		break
	case reflect.Struct:
		// TODO: Check the JSON property and use json.Marshal to serialize the struct

		// TODO: Test to ensure that we can safely serialize the body
		if err := binary.Write(body, binary.BigEndian, b); err != nil {
			panic(err)
		}
		break
	}

	// TODO: Change headers property to be a struct ?
	o.Headers["Content-Type"] = contentType

	return body
}

// REMARKS: The Body in the http.Response will be closed when returning a response to the caller
func (r *Request) doRequest(m string, o *Option) (*http.Response, []byte, error) {
	if o.Headers == nil {
		o.Headers = make(map[string]string)
	}
	body := getRequestBody(o)
	req, err := http.NewRequest(m, o.Url, body)

	if err != nil {
		panic(err)
	}

	if o.Auth != nil {
		if o.Auth.Bearer != "" {
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", o.Auth.Bearer))
		} else if o.Auth.Username != "" && o.Auth.Password != "" {
			req.SetBasicAuth(o.Auth.Username, o.Auth.Password)
		}
	} else if usr, pwd, err := splitUserNamePassword(o.Url); err != nil {
		// TODO: Should we panic if an error is returned or silently ignore this - maybe give some warning ?
		//panic(err)
	} else {
		if usr != "" && pwd != "" {
			req.SetBasicAuth(usr, pwd)
		}
	}

	// TODO: Validate headers against known list of headers ?
	// TODO: Ensure headers are only set once
	// TODO: If JSON property set, add Content-Type: application/json if not already set in o.Headers
	for k, v := range o.Headers {
		req.Header.Add(k, v)
	}

	resp, err := r.client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return resp, nil, err
	}

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return resp, nil, err
	} else {
		return resp, body, nil
	}
}
