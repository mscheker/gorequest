package request

import (
	"io/ioutil"
	"net/http"
)

type request struct {
	request *http.Request
	client  *http.Client
}

func newRequest(req *http.Request, client *http.Client) Request {
	return &request{
		request: req,
		client:  client,
	}
}

func (r *request) getUnderlyingRequest() *http.Request {
	return r.request
}

func (r *request) getUnderlyingHttpClient() *http.Client {
	return r.client
}

func (r *request) Do() Response {
	resp, err := r.client.Do(r.request)

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
