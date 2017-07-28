package gorequest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method, "Should equal request method")

		resp.WriteHeader(http.StatusOK)
		fmt.Fprintf(resp, "Hello World")
	}))
	defer ts.Close()

	r := NewRequestBuilder().WithUrl(ts.URL).Build().Do()

	assert.NotNil(t, r, "Should not be nil")
	assert.Equal(t, http.StatusOK, r.Response().StatusCode, "Should equal HTTP Status 200 (OK)")
	assert.Equal(t, "Hello World", string(r.Body()), "Should equal response body")
}

func TestGetRequestConvenienceMethod(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method, "Should equal request method")

		resp.WriteHeader(http.StatusOK)
		fmt.Fprintf(resp, "Hello World")
	}))
	defer ts.Close()

	r := Get(ts.URL)

	assert.NotNil(t, r, "Should not be nil")
	assert.Equal(t, http.StatusOK, r.Response().StatusCode, "Should equal HTTP Status 200 (OK)")
	assert.Equal(t, "Hello World", string(r.Body()), "Should equal response body")
}

func TestPostRequestWithTextBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "POST", req.Method, "Should equal request method")
		assert.Equal(t, "text/plain", req.Header.Get("Content-Type"), "Should equal Content-Type header")

		defer req.Body.Close()

		if b, err := ioutil.ReadAll(req.Body); err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(resp, err.Error())
		} else {
			resp.WriteHeader(http.StatusOK)
			fmt.Fprintf(resp, string(b))
		}
	}))
	defer ts.Close()

	r := NewRequestBuilder().WithMethod("POST").WithUrl(ts.URL).WithTextBody("Hello World").Build().Do()

	assert.NotNil(t, r, "Should not be nil")
	assert.Equal(t, http.StatusOK, r.Response().StatusCode, "Should equal HTTP Status 200 (OK)")
	assert.Equal(t, "Hello World", string(r.Body()), "Should equal request body")
}

func TestPostTextRequestConvenienceMethod(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "POST", req.Method, "Should equal request method")
		assert.Equal(t, "text/plain", req.Header.Get("Content-Type"), "Should equal Content-Type header")

		defer req.Body.Close()

		if b, err := ioutil.ReadAll(req.Body); err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(resp, err.Error())
		} else {
			resp.WriteHeader(http.StatusOK)
			fmt.Fprintf(resp, string(b))
		}
	}))
	defer ts.Close()

	r := PostText(ts.URL, "Hello World")

	assert.NotNil(t, r, "Should not be nil")
	assert.Equal(t, http.StatusOK, r.Response().StatusCode, "Should equal HTTP Status 200 (OK)")
	assert.Equal(t, "Hello World", string(r.Body()), "Should equal request body")
}

func TestPostJsonRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "POST", req.Method, "Should equal request method")
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"), "Should equal Content-Type header")

		var s *testJsonStruct

		decoder := json.NewDecoder(req.Body)

		if err := decoder.Decode(&s); err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(resp, err.Error())
		} else {
			assert.Equal(t, 10, s.IntField, "Should equal IntField")
			assert.Equal(t, "Hello World", s.StringField, "Should equal StringField")
			assert.True(t, s.BoolField, "Should be true")

			resp.WriteHeader(http.StatusOK)
			fmt.Fprintf(resp, "OK")
		}
	}))
	defer ts.Close()

	testJsonData := &testJsonStruct{
		IntField:    10,
		StringField: "Hello World",
		BoolField:   true,
	}

	r := NewRequestBuilder().WithMethod("POST").WithUrl(ts.URL).WithJsonBody(testJsonData).Build().Do()

	assert.NotNil(t, r, "Should not be nil")
	assert.Equal(t, http.StatusOK, r.Response().StatusCode, "Should equal HTTP Status 200 (OK)")
}

//func TestPutRequest(t *testing.T) {
//  assert.True(t, false, "Not Implemented")
//}

//func TestDeleteRequest(t *testing.T) {
//  assert.True(t, false, "Not Implemented")
//}

func TestBasicAuthentication(t *testing.T) {
	r := NewRequestBuilder().WithMethod("GET").WithUrl(POSTMAN_ECHO_BASIC_AUTH_ENDPOINT).WithBasicAuth("postman", "password").Build().Do()

	assert.NotNil(t, r, "Should not be nil")
	assert.Equal(t, http.StatusOK, r.Response().StatusCode, "Should equal HTTP Status 200 (OK)")
}

func TestBasicAuthenticationWithRFC1738(t *testing.T) {
	basicAuthEndpoint := "https://postman:password@postman-echo.com/basic-auth"
	r := NewRequestBuilder().WithMethod("GET").WithRFC1738(basicAuthEndpoint).Build().Do()

	assert.NotNil(t, r, "Should not be nil")
	assert.Equal(t, http.StatusOK, r.Response().StatusCode, "Should equal HTTP Status 200 (OK)")
}

func TestDigestAuthentication(t *testing.T) {
	r := NewRequestBuilder().WithMethod("GET").WithUrl("https://postman-echo.com/digest-auth").WithDigestAuth("postman", "password").Build().Do()

	assert.NotNil(t, r, "Should not be nil")
	assert.Equal(t, http.StatusOK, r.Response().StatusCode, "Should equal HTTP Status 200 (OK)")
}

func TestWithCheckRedirect(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Location", POSTMAN_ECHO_GET_ENDPOINT)
		resp.WriteHeader(http.StatusMovedPermanently)

		fmt.Fprintf(resp, "")
	}))
	defer ts.Close()

	redirectPol := func(req *http.Request, via []*http.Request) error {
		assert.Equal(t, "I am going to be redirected", via[0].Header.Get("TestHeader"), "Should equal request header")

		return nil
	}

	b := NewRequestBuilder().WithUrl(ts.URL).WithCheckRedirect(redirectPol)
	r := b.WithHeader("TestHeader", "I am going to be redirected").Build().Do()

	assert.NotNil(t, r, "Should not be nil")
	assert.Equal(t, http.StatusOK, r.Response().StatusCode, "Should equal HTTP Status 200 (OK)")
	assert.Equal(t, POSTMAN_ECHO_GET_ENDPOINT, r.Response().Request.URL.String(), "Should equal redirect URL")
}
