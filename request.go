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

// ***********************************************
// ************* Convenience Methods *************
// ***********************************************
var builder r.RequestBuilder

func Get(url string) r.Response {
	return getInstance().WithMethod("GET").WithUrl(url).Build().Do()
}

func PostText(url, data string) r.Response {
	return getInstance().WithMethod("POST").WithUrl(url).WithTextBody(data).Build().Do()
}

func PostJson(url string, data interface{}) r.Response {
	return getInstance().WithMethod("POST").WithUrl(url).WithJsonBody(data).Build().Do()
}

func PutText(url, data string) r.Response {
	return getInstance().WithMethod("PUT").WithUrl(url).WithTextBody(data).Build().Do()
}

func PutJson(url string, data interface{}) r.Response {
	return getInstance().WithMethod("PUT").WithUrl(url).WithJsonBody(data).Build().Do()
}

func Delete(url string) r.Response {
	return getInstance().WithMethod("DELETE").WithUrl(url).Build().Do()
}

func Head(url string) r.Response {
	return getInstance().WithMethod("HEAD").WithUrl(url).Build().Do()
}

// ***********************************************
// ********** Private methods/functions **********
// ***********************************************

func getInstance() r.RequestBuilder {
	if builder == nil {
		builder = NewRequestBuilder()
	}

	return builder
}
