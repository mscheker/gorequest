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

## Table of Contents

## Convenience Methods

There are methods for each different HTTP Verb; these methods are similar to NewRequest() but the method field is set for you:

* request.Get() - Defaults to method: "GET"
* request.Post() - Defaults to method: "POST"
* request.Put() - Defaults to method: "PUT"
* request.Delete() - Defaults to method: "DELETE"
* request.Head() - Defaults to method: "HEAD"

## Credits
