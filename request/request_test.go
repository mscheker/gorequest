package request

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	user    = "test"
	pass    = "12345"
	hash    = "Z29sYW5ndGVzdA=="
	testUrl = "https://www.google.com"
)

func TestValidateSingleInstance(t *testing.T) {
	i1 := getInstance()
	i2 := getInstance()

	assert.NotNil(t, i1, "Should not be nil")
	assert.NotNil(t, i2, "Should not be nil")
	assert.True(t, i1 == i2, "Should be the same instance")
}

func TestValidateMultipleInstances(t *testing.T) {
	i1 := getInstance()
	instance = nil
	i2 := getInstance()

	assert.NotNil(t, i1, "Should not be nil")
	assert.NotNil(t, i2, "Should not be nil")
	assert.False(t, i1 == i2, "Should be different instances")
}

func TestValidateNewAuth(t *testing.T) {
	auth := NewAuth(user, pass, hash)

	assert.Equal(t, user, auth.Username, "Should equal username")
	assert.Equal(t, pass, auth.Password, "Should equal password")
	assert.Equal(t, hash, auth.Bearer, "Should equal token")
}

func TestValidateDefaultHttpClientTimeout(t *testing.T) {
	r := New()

	assert.Equal(t, 30*time.Second, r.client.Timeout, "Should default to 30 seconds")
}

func TestValidateOverridingHttpClientTimeout(t *testing.T) {
	// REMARKS: Override timeout value to 45 seconds
	r := New(45)

	assert.Equal(t, 45*time.Second, r.client.Timeout, "Should equals 45 seconds")
}
