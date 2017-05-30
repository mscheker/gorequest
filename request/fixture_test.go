package request

import (
	"os"
	"testing"
)

type testJsonStruct struct {
	IntField    int    `json:"intField"`
	StringField string `json:"stringField"`
	BoolField   bool   `json:"boolField"`
}

const (
	POSTMAN_ECHO_ROOT                = "https://postman-echo.com"
	POSTMAN_ECHO_BASIC_AUTH_ENDPOINT = "https://postman-echo.com/basic-auth"

	TEST_TOKEN = "Z29sYW5ndGVzdA=="
)

func TestMain(runner *testing.M) {
	result := -1

	defer func() {
		os.Exit(result)
	}()

	result = runner.Run()
}
