package request

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRequestBuilderPanicWithNoUrl(t *testing.T) {
	defer func() {
		err := recover().(error)

		assert.NotNil(t, err, "Should not be nil")
		assert.Equal(t, "URL is required.", err.Error(), "Should equal error message")
	}()

	builder := NewRequestBuilder()
	builder.Build()

	assert.True(t, false, "Should not have completed test")
}

func TestRequestBuilderWithUrl(t *testing.T) {
	r1 := NewRequestBuilder().WithUrl(POSTMAN_ECHO_ROOT).Build()

	assert.NotNil(t, r1, "Should not be nil")

	r2 := r1.getUnderlyingRequest()

	assert.NotNil(t, r2, "Should not be nil")
	assert.Equal(t, POSTMAN_ECHO_ROOT, r2.URL.String(), "Should equal URL")
}

func TestRequestBuilderWithDefaults(t *testing.T) {
	r1 := NewRequestBuilder().WithUrl(POSTMAN_ECHO_ROOT).Build()

	assert.NotNil(t, r1, "Should not be nil")

	r2 := r1.getUnderlyingRequest()

	assert.NotNil(t, r2, "Should not be nil")
	assert.Equal(t, "GET", r2.Method, "Should equal GET method")
	assert.Empty(t, r2.Header.Get("Authorization"), "Should not have set authorization header")
}

func TestRequestBuilderWithDefaultTimeout(t *testing.T) {
	r1 := NewRequestBuilder().WithUrl(POSTMAN_ECHO_ROOT).Build()

	assert.NotNil(t, r1, "Should not be nil")

	c := r1.getUnderlyingHttpClient()

	assert.NotNil(t, c, "Should not be nil")
	assert.Equal(t, 30*time.Second, c.Timeout, "Should equal 30 seconds")
}

func TestRequestBuilderWithTimeout(t *testing.T) {
	r1 := NewRequestBuilder().WithUrl(POSTMAN_ECHO_ROOT).WithTimeout(45 * time.Second).Build()

	assert.NotNil(t, r1, "Should not be nil")

	c := r1.getUnderlyingHttpClient()

	assert.NotNil(t, c, "Should not be nil")
	assert.Equal(t, 45*time.Second, c.Timeout, "Should equal 45 seconds")
}
