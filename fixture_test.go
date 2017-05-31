package gorequest

import (
	"os"
	"testing"
)

const (
	POSTMAN_ECHO_ROOT                = "https://postman-echo.com"
	POSTMAN_ECHO_BASIC_AUTH_ENDPOINT = "https://postman-echo.com/basic-auth"

	POSTMAN_ECHO_POST_ENDPOINT   = "https://postman-echo.com/post"
	POSTMAN_ECHO_GET_ENDPOINT    = "https://postman-echo.com/get"
	POSTMAN_ECHO_PUT_ENDPOINT    = "https://postman-echo.com/put"
	POSTMAN_ECHO_PATCH_ENDPOINT  = "https://postman-echo.com/patch"
	POSTMAN_ECHO_DELETE_ENDPOINT = "https://postman-echo.com/delete"

	TEST_TOKEN = "Z29sYW5ndGVzdA=="
)

func TestMain(runner *testing.M) {
	result := -1

	defer func() {
		os.Exit(result)
	}()

	result = runner.Run()
}
