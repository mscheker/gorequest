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
* [With URL](#with-url---defaults-to-method-get)
* [With Options](#with-options)

[Options](#options)

[Authentication](#authentication)
* [Basic Authentication](#basic-authentication)
* [Bearer Authentication](#bearer-authentication)

[Convenience Methods](#convenience-methods)

## Simple to Use
### With URL - Defaults to method: GET
```go
package main

import (
	"fmt"
	
	request "github.com/mscheker/gorequest"
)

func main() {
	resp, body, err := request.NewRequest("https://www.google.com")
	
	fmt.Printf("Response: %v \n\r", resp)
	fmt.Printf("Body: %s \n\r", string(body))
	fmt.Printf("Error: %v \n\r", err)
}
```

### With Options
```go
package main

import (
	"fmt"

	request "github.com/mscheker/gorequest"
)

func main() {
	options := &request.Option{
		Url:    "https://www.google.com",
		Method: "GET",
	}
	resp, body, err := request.NewRequest(options)

	fmt.Printf("Response: %v \n\r", resp)
	fmt.Printf("Body: %s \n\r", string(body))
	fmt.Printf("Error: %v \n\r", err)
}
```

## Options
```go
func NewRequest(val interface{}) (*http.Response, []byte, error) {...}
```
The argument can be either a URL or an options struct. The only required option is URL; all others are optional.
```go
type Option struct {
	Url     string
	Headers map[string]string
	Auth    *auth
	Body    interface{}
	JSON    interface{}
	Method  string
}
```
* `Url` - Fully qualified URL.
* `Method` - HTTP method (Defaults to "GET").
* `Headers` - HTTP headers (Defaults to an empty map).
* `Body` - Entity body for POST and PUT requests. Must be a string or struct. If JSON is true, Body must be a JSON serializable struct or a valid JSON formatted string. `Body is ignored for GET and DELETE requests`.
* `JSON` - A JSON serializable struct or a valid JSON formatted string. Sets the Body to a JSON representation of the data and sets the `Content-Type header to application/json`. If set to true, it will attempt to serialize the Body.
* `Auth` - A struct containing values for `username` and `password`, and `bearer` token.

## Authentication
If passed as an option, `Auth` is a struct containing the values:
* `Username`
* `Password`
* `Bearer` (Optional)
```go
// username, password, bearer
func NewAuth(vals ...string) *auth {...}
```

### Basic Authentication
Basic authentication is supported, and it is set when a `username` and `password` are provided as part of the `Auth` option.
```go
package main

import (
	"fmt"

	request "github.com/mscheker/gorequest"
)

func main() {
	options := &request.Option{
		Url:    "https://postman-echo.com/basic-auth",
		Method: "GET",
		Auth:   request.NewAuth("postman", "password"),
	}
	resp, body, err := request.NewRequest(options)

	fmt.Printf("Response: %v \n\r", resp)
	fmt.Printf("Body: %s \n\r", string(body))
	fmt.Printf("Error: %v \n\r", err)
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
	options := &request.Option{
		Url:    "https://postman:password@postman-echo.com/basic-auth",
		Method: "GET",
	}
	resp, body, err := request.NewRequest(options)

	fmt.Printf("Response: %v \n\r", resp)
	fmt.Printf("Body: %s \n\r", string(body))
	fmt.Printf("Error: %v \n\r", err)
}
```

### Bearer Authentication
Bearer authentication is supported, and it is set when the `bearer` value is provided as part of the `Auth` option.
```go
package main

import (
	"fmt"

	request "github.com/mscheker/gorequest"
)

func main() {
	options := &request.Option{
		Url:    "https://your_endpoint",
		Method: "GET",
		Auth:   request.NewAuth("", "", "your_bearer_token"),
	}
	resp, body, err := request.NewRequest(options)

	fmt.Printf("Response: %v \n\r", resp)
	fmt.Printf("Body: %s \n\r", string(body))
	fmt.Printf("Error: %v \n\r", err)
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
