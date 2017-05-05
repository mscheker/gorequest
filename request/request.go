package request

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

func (r *Request) Get(o *Option) (*http.Response, []byte, error) {
	return r.doRequest("GET", o)
}

func Get(o *Option) (*http.Response, []byte, error) {
	return getInstance().doRequest("GET", o)
}

func (r *Request) Delete(o *Option) (*http.Response, []byte, error) {
	return r.doRequest("DELETE", o)
}

func Delete(o *Option) (*http.Response, []byte, error) {
	return getInstance().doRequest("DELETE", o)
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

// REMARKS: The Body in the http.Response will be closed when returning a response to the caller
func (r *Request) doRequest(m string, o *Option) (*http.Response, []byte, error) {
	req, err := http.NewRequest(m, o.Url, nil)

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

	resp, err := r.client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return resp, nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return resp, nil, err
	}

	return resp, body, nil
}
