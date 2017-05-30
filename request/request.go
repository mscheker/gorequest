package request

import (
	"io/ioutil"
	"net/http"
)

type request struct {
	request *http.Request
}

func newRequest(req *http.Request) Request {
	return &request{
		request: req,
	}
}

func (r *request) getUnderlyingRequest() *http.Request {
	return r.request
}

func (r *request) Do() Response {
	client := getDefaultHttpClient()

	resp, err := client.Do(r.request)

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return &response{
		body:     body,
		response: resp,
	}
}
