package request

import (
	"net/http"
	"time"
)

type Options struct {
	Url string
}

type Request struct {
	client  *http.Client
	Timeout time.Duration
}

func (r *Request) Get(o *Options) {

}

func Get(o *Options) {

}

func (r *Request) doRequest(m string, o *Options) (*http.Response, []byte, error) {

}
