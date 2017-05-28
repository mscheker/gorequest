package gorequest

/**
 * TODO: document
 */

import (
  impl "github.com/demianlessa/gorequest/impl"
  model "github.com/demianlessa/gorequest/model"
)

/**
 * Single entry point into the API. A Request instance can only be created 
 * using a RequestBuilder instance, and this is the only public RequestBuilder
 * constructor.
 */
var NewRequestBuilder model.RequestBuilderConstructor = impl.NewRequestBuilder;
