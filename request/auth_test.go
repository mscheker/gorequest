package request

import (
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthNone(t *testing.T) {
	req, err := http.NewRequest("GET", POSTMAN_ECHO_ROOT, nil)

	assert.Nil(t, err, "Should be nil")

	auth := newAuthNone()
	auth.Configure(req)

	assert.Empty(t, req.Header.Get("Authorization"), "Should not have set authorization header")
}

func TestBasicAuth(t *testing.T) {
	req, err := http.NewRequest("GET", POSTMAN_ECHO_BASIC_AUTH_ENDPOINT, nil)

	assert.Nil(t, err, "Should be nil")

	auth := newAuthBasic("postman", "password")
	auth.Configure(req)

	basicAuthString := base64.StdEncoding.EncodeToString([]byte("postman:password"))
	basicAuthString = "Basic " + basicAuthString

	assert.Equal(t, basicAuthString, req.Header.Get("Authorization"), "Should equal authorization header value")
}

func TestBearerAuth(t *testing.T) {
	req, err := http.NewRequest("GET", POSTMAN_ECHO_ROOT, nil)

	assert.Nil(t, err, "Should be nil")

	auth := newAuthBearer(TEST_TOKEN)
	auth.Configure(req)

	bearerAuthString := "Bearer " + TEST_TOKEN

	assert.Equal(t, bearerAuthString, req.Header.Get("Authorization"), "Should equal authorization header value")
}
