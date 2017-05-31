package gorequest

/**
 * TODO: document
 */

import (
	r "github.com/mscheker/gorequest/request"
)

/**
 * Single entry point into the API. A Request instance can only be created
 * using a RequestBuilder instance, and this is the only public RequestBuilder
 * constructor.
 */
var NewRequestBuilder r.RequestBuilderConstructor = r.NewRequestBuilder
