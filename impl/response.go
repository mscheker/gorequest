package gorequest

import (
	"net/http"
)

/****************************************************
 * api.Response implementation
 ****************************************************/

type response struct {
	body []byte
	response *http.Response
}

func (r *response) Body() []byte {
	return r.body
}

func (r *response) Response() *http.Response {
	return r.response
}
