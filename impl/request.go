package gorequest

import (
	model "gorequest/model"
	"io/ioutil"
	"net/http"
)

/****************************************************
 * model.Request implementation
 ****************************************************/

type request struct {
	request *http.Request
}

func newRequest(req *http.Request) model.Request {
	return &request{
		request: req,
	}
}

func (r *request) Do() model.Response {

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
		body: body,
		response: resp,
	}
}
