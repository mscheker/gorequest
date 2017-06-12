# GoRequest - ![Build Status](https://travis-ci.org/mscheker/gorequest.svg?branch=master)
Simplified HTTP Client inspired by the Node.js [request](https://github.com/request/request) module

## Disclaimer
Not all functionality from the original Node.js request module has been yet ported. WIP

## Description
GoRequest makes HTTP requests simple and idiomatic. It was implemented and designed to be simple when making HTTP calls.

## Installation
```
$ go get github.com/mscheker/gorequest
```

## Table of Contents
[Simple to Use](#simple-to-use)
* [With Convenience Methods](#with-convenience-methods)
* [With Request Builder](#with-request-builder)

[Request Builder Methods](#request-builder-methods)

[Authentication](#authentication)
* [Basic Authentication](#basic-authentication)
* [Bearer Authentication](#bearer-authentication)

[Convenience Methods](#convenience-methods)

[Credits](#credits)

## Simple to Use
## With Convenience Methods
```go
package main

import (
	"fmt"

	request "github.com/mscheker/gorequest"
)

func main() {
	resp := request.Get("https://www.google.com")
	
	fmt.Printf("Body: %s \n\r", string(resp.Body()))
	fmt.Printf("Status: %s \n\r", resp.Response().Status)
}
```

## With Request Builder
If you need more control when making requests, the package exposes a constructor for a `RequestBuilder` object.
```go
package main

import (
	"fmt"

	request "github.com/mscheker/gorequest"
)

func main() {
	resp := request.NewRequestBuilder().WithUrl("https://www.google.com").Build().Do()
	
	fmt.Printf("Body: %s \n\r", string(resp.Body()))
	fmt.Printf("Status: %s \n\r", resp.Response().Status)
}
```

## Request Builder Methods
When building a request, the only required option is the URL; the method will default to `GET` if none is specified.

* `Url` - Fully qualified URL.
* `Method` - HTTP method (Defaults to "GET").
* `Headers` - HTTP headers (Defaults to an empty map).
* `Body` - Entity body for POST and PUT requests. Must be a string or struct. If JSON is true, Body must be a JSON serializable struct or a valid JSON formatted string. `Body is ignored for GET and DELETE requests`.
* `JSON` - A JSON serializable struct or a valid JSON formatted string. Sets the Body to a JSON representation of the data and sets the `Content-Type header to application/json`. If set to true, it will attempt to serialize the Body.
* `Auth` - A struct containing values for `username` and `password`, and `bearer` token.

## Authentication
The builder exposes various methods for the different authentication mechanisms that are supported:
* Basic Authentication
* Bearer Authentication

### Basic Authentication
Basic authentication is supported, and it is set when a `username` and `password` are provided as part of the `WithBasicAuth` method.
```go
package main

import (
	"fmt"

	request "github.com/mscheker/gorequest"
)

func main() {
	resp := request.NewRequestBuilder().WithUrl("https://postman-echo.com/basic-auth").WithBasicAuth("postman", "password").Build().Do()

	fmt.Printf("Body: %s \n\r", string(resp.Body()))
	fmt.Printf("Status: %s \n\r", resp.Response().Status)
}
```
You can also specify basic authentication using the URL itself, as detailed in [RFC 1738](http://www.ietf.org/rfc/rfc1738.txt).
```go
package main

import (
	"fmt"

	request "github.com/mscheker/gorequest"
)

func main() {
	resp := request.NewRequestBuilder().WithUrl("https://postman:password@postman-echo.com/basic-auth").Build().Do()

	fmt.Printf("Body: %s \n\r", string(resp.Body()))
	fmt.Printf("Status: %s \n\r", resp.Response().Status)
}
```

### Bearer Authentication
Bearer authentication is supported, and it is set when the `bearer` value is provided as part of the `WithBearerAuth` method.
```go
package main

import (
	"fmt"

	request "github.com/mscheker/gorequest"
)

func main() {
	resp := request.NewRequestBuilder().WithUrl("https://your_endpoint").WithBearerAuth("your_bearer_token").Build().Do()

	fmt.Printf("Body: %s \n\r", string(resp.Body()))
	fmt.Printf("Status: %s \n\r", resp.Response().Status)
}
```

## Convenience Methods

There are methods for each different HTTP Verb; these methods are similar to NewRequest() but the method field is set for you:

* request.Get() - Defaults to method: "GET"
* request.Post() - Defaults to method: "POST"
* request.Put() - Defaults to method: "PUT"
* request.Delete() - Defaults to method: "DELETE"
* request.Head() - Defaults to method: "HEAD"

## Credits
* [Postman Echo](https://docs.postman-echo.com) for providing a service to test REST clients, API calls, and various auth mechanisms.
* To the team behind the Node.js [request](https://github.com/request/request) module for implementing a robust yet simple to use library which is the inspiration for this package.
