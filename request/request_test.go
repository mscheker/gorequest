package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
