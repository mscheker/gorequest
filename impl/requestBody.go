package gorequest

import (
	"bytes"
	"encoding/json"
	"errors"
	model "gorequest/model"
	"reflect"
)

/****************************************************
 * model.RequestBody implementation
 ****************************************************/

type requestBody struct {
	contentType string
	data *bytes.Buffer
}

func newJsonBody(data interface{}) model.RequestBody {

	var buffer *bytes.Buffer

	indirect := reflect.Indirect(reflect.ValueOf(data))

	switch indirect.Kind() {
	case reflect.String:
		buffer = bytes.NewBuffer([]byte(indirect.String()))
		break
	case reflect.Struct:
		if rawBytes, err := json.Marshal(indirect.Interface()); err == nil {
			buffer = bytes.NewBuffer(rawBytes)
		} else {
			panic(err)
		}
		break
	default:
		panic(errors.New("Can only serialize a string or a struct as JSON content."))
	}

	return &requestBody{
		contentType: "application/json",
		data: buffer,
	}
}

func (b *requestBody) ContentType() string {
	return b.contentType
}

func (b *requestBody) RawData() *bytes.Buffer {
	return b.data
}
