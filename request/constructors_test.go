package request

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRequestBuilderWithDefaults(t *testing.T) {
	builder := NewRequestBuilder()
	builder.WithUrl(POSTMAN_ECHO_ROOT)

	r1 := builder.Build()

	assert.NotNil(t, r1, "Should not be nil")

	r2 := r1.getUnderlyingRequest()

	assert.NotNil(t, r2, "Should not be nil")
	assert.Equal(t, "GET", r2.Method, "Should equal GET method")
	assert.True(t, len(r2.Header) == 0, "Should have empty header map")
	assert.Empty(t, r2.Header.Get("Authorization"), "Should not have set authorization header")
}

func TestDefaultHttpClient(t *testing.T) {
	client := getDefaultHttpClient()

	assert.NotNil(t, client, "Should not be nil")
	assert.Equal(t, 30*time.Second, client.Timeout, "Should equal 30 seconds timeout")
}

func TestHttpClientWithCustomTimeout(t *testing.T) {
	client := getHttpClient(45)

	assert.NotNil(t, client, "Should not be nil")
	assert.Equal(t, 45*time.Second, client.Timeout, "Should equal 45 seconds timeout")
}
