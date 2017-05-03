package request

import (
	"io/ioutil"
	"net/http"
	"time"
)

var instance *Request

type Options struct {
	Url string
}

type Request struct {
	client  *http.Client
	Timeout time.Duration
}

func New() *Request {
	r := new(Request)

	r.Timeout = 30 * time.Second

	r.client = &http.Client{
		Timeout: r.Timeout,
	}

	return r
}

func (r *Request) Get(o *Options) (*http.Response, []byte, error) {
	return r.doRequest("GET", o)
}

func Get(o *Options) (*http.Response, []byte, error) {
	return getInstance().doRequest("GET", o)
}

// ********** Private methods/functions **********
// REMARKS: Used internally by non-instance methods
func getInstance() *Request {
	if instance == nil {
		instance = New()
	}

	return instance
}

func (r *Request) doRequest(m string, o *Options) (*http.Response, []byte, error) {
	req, err := http.NewRequest(m, o.Url, nil)

	if err != nil {
		panic(err)
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
