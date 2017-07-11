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
* [Digest Authentication](#digest-authentication)

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

* `WithUrl` - Fully qualified URL.
* `WithRFC1738` - Full qualified URL with `username` and `password` for `Basic Authentication`.
* `WithMethod` - HTTP method (Defaults to "GET").
* `WithHeader` - HTTP header (Defaults to an empty map).
* `WithTextBody` - Body for POST and PUT requests. Must be a string. `Content-Type` header is set to `text/plain`.
* `WithJsonBody` - Body for POST and PUT requests. Must be a valid JSON formatted string or a JSON serializable struct. `Content-Type` header is set to `application/json`.
* `WithBasicAuth` - Generates a Base64 encoded string from the `username` and `password` specified, and sets the `Authorization` header to `Basic <encoded_string>` accordingly.
* `WithBearerAuth` - Sets the `Authorization` header to `Bearer <your_bearer_token>` accordingly.
* `WithDigestAuth` - Generates the necessary MD5 hash and nonce values, and sets the hashed Digest `Authorization` header accordingly before resending the request.
* `WithTimeout` - Sets the time limit for requests made by the HTTP client. Defaults to `30 seconds`.
* `Build` - Builds a request object with the specified options. Will panic if a `URL` has not been set.

```
Note: Body is ignored for GET, DELETE and HEAD requests.
```

## Authentication
The builder exposes various methods for the different authentication mechanisms that are supported:
* Basic
* Bearer
* Digest

### Basic Authentication
Basic authentication is supported, and it is set when a `username` and `password` is provided as part of the `WithBasicAuth` method.
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
    resp := request.NewRequestBuilder().WithRFC1738("https://postman:password@postman-echo.com/basic-auth").Build().Do()

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

### Digest Authentication
Digest authentication is supported, and it is set when a `username` and `password` is provided as part of the `WithDigestAuth` method. Upon detecting a `401 (Unauthorized)` in the initial response, the request is sent again with the hashed Digest Authorization header.
```go
package main

import (
    "fmt"

    request "github.com/mscheker/gorequest"
)

func main() {
    resp := request.NewRequestBuilder().WithMethod("GET").WithUrl("https://postman-echo.com/digest-auth").WithDigestAuth("postman", "password").Build().Do()

    fmt.Printf("Body: %s \n\r", string(resp.Body()))
    fmt.Printf("Status: %s \n\r", resp.Response().Status)
}
```

## Convenience Methods

There are methods for each different HTTP Verb; the method field is set for you. In the PostText, PostJson, PutText and PutJson methods, the Content-Type header is set accordingly:

* request.Get() - Defaults to method: "GET".
* request.PostText() - Defaults to method: "POST" and Content-Type: "text/plain".
* request.PostJson() - Defaults to method: "POST" and Content-Type: "application/json".
* request.PutText() - Defaults to method: "PUT" and Content-Type: "text/plain".
* request.PutJson() - Defaults to method: "PUT" and Content-Type: "application/json".
* request.Delete() - Defaults to method: "DELETE".
* request.Head() - Defaults to method: "HEAD".

## Credits
* [Postman Echo](https://docs.postman-echo.com) for providing a service to test REST clients, API calls, and various auth mechanisms.
* To the team behind the Node.js [request](https://github.com/request/request) module for implementing a robust yet simple to use library which is the inspiration for this package.
* To [Demian Lessa](https://github.com/demianlessa) for contributing with the revised design and refactoring of the package.
