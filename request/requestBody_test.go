package request

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJsonBodyPanic(t *testing.T) {
	defer func() {
		err := recover().(error)

		assert.NotNil(t, err, "Should not be nil")
		assert.Equal(t, "Can only serialize a string or a struct as JSON content.", err.Error(), "Should equal error message")
	}()

	newJsonBody(10)

	assert.True(t, false, "Should not have completed test")
}

func TestNewJsonBodyWithString(t *testing.T) {
	b := "Hello World"

	body := newJsonBody(b)

	assert.NotNil(t, body, "Should not be nil")
	assert.Equal(t, "application/json", body.ContentType(), "Should equal Content-Type")
	assert.Equal(t, b, body.RawData().String(), "Should equal RawData")
}

func TestNewJsonBodyWithStruct(t *testing.T) {
	testJsonData := &testJsonStruct{
		IntField:    10,
		StringField: "Hello World",
		BoolField:   true,
	}
	b, err := json.Marshal(testJsonData)

	assert.Nil(t, err, "Should be nil")

	body := newJsonBody(testJsonData)

	assert.NotNil(t, body, "Should not be nil")
	assert.Equal(t, "application/json", body.ContentType(), "Should equal Content-Type")
	assert.Equal(t, string(b), body.RawData().String(), "Should equal RawData")

	var r *testJsonStruct

	e := json.Unmarshal(body.RawData().Bytes(), &r)

	assert.Nil(t, e, "Should be nil")

	assert.Equal(t, testJsonData.IntField, r.IntField, "Should be equal")
	assert.Equal(t, testJsonData.StringField, r.StringField, "Should be equal")
	assert.Equal(t, testJsonData.BoolField, r.BoolField, "Should be equal")
}
