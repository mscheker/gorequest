package gorequest

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

const (
	user    = "test"
	pass    = "12345"
	hash    = "Z29sYW5ndGVzdA=="
	testUrl = "https://www.google.com"
)

type TestCustomer struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type TestOption struct {
	Url    string
	Method string
}

func newTestCustomer(id int, firstName, lastName string) *TestCustomer {
	return &TestCustomer{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
}

var testCustomers = make([]*TestCustomer, 0)
var testRouter *mux.Router
var testRecorder *httptest.ResponseRecorder

func init() {
	testCustomers = append(testCustomers,
		newTestCustomer(1, "John", "Doe"),
		newTestCustomer(2, "Jane", "Doe"))
}

func TestValidateSingleInstance(t *testing.T) {
	i1 := getDefaultHttpClient()
	i2 := getDefaultHttpClient()

	assert.NotNil(t, i1, "Should not be nil")
	assert.NotNil(t, i2, "Should not be nil")
	assert.True(t, i1 == i2, "Should be the same instance")
}

func TestValidateMultipleInstances(t *testing.T) {
	i1 := getDefaultHttpClient()
	httpClient = nil
	i2 := getDefaultHttpClient()

	assert.NotNil(t, i1, "Should not be nil")
	assert.NotNil(t, i2, "Should not be nil")
	assert.False(t, i1 == i2, "Should be different instances")
}

func TestValidateNewAuth(t *testing.T) {
	auth := newAuthBearer(hash)
	
	assert.Equal(t, hash, auth.(*authBearer).token, "Should equal token")
}

func TestValidateDefaultHttpClientTimeout(t *testing.T) {
	r := getDefaultHttpClient()
	assert.Equal(t, defaultTimeout, r.Timeout, "Should use the default timeout")
}

func TestValidateOverridingHttpClientTimeout(t *testing.T) {
	r := getHttpClient(45)
	assert.Equal(t, 45*time.Second, r.Timeout, "Should use the specified timeout: 45 seconds")
}

/*
func TestSplitUserNamePassword(t *testing.T) {
	// REMARKS: The user/pwd can be provided in the URL when doing Basic Authentication (RFC 1738)
	url := "https://testuser:testpass12345@mysite.com"

	usr, pwd, err := splitUserNamePassword(url)

	assert.Equal(t, "testuser", usr, "Should equal username")
	assert.Equal(t, "testpass12345", pwd, "Should equal password")
	assert.Nil(t, err, "Should be nil")
}

func TestSplitUserNamePasswordNoCredentialsFound(t *testing.T) {
	url := "https://mysite.com"

	usr, pwd, err := splitUserNamePassword(url)

	assert.Empty(t, usr, "Should be empty")
	assert.Empty(t, pwd, "Should be empty")
	assert.EqualError(t, err, "No credentials found in URI")

	url = "https://@mysite.com"

	u, p, e := splitUserNamePassword(url)

	assert.Empty(t, u, "Should be empty")
	assert.Empty(t, p, "Should be empty")
	assert.EqualError(t, e, "No credentials found in URI")
}
*/

func TestNewRequestWithUrl(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if json, err := json.Marshal(testCustomers); err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(resp, err.Error())
		} else {
			fmt.Fprintf(resp, string(json))
		}
	}))

	defer ts.Close()

	response := NewRequestBuilder().WithUrl(ts.URL).Build().Do()

	assert.Equal(t, "GET", response.Response().Request.Method, "Should equal GET method")
	assert.Equal(t, 200, response.Response().StatusCode, "Should equal HTTP Status 200 (OK)")

	customers := make([]*TestCustomer, 0)

	err := json.Unmarshal(response.Body(), &customers)

	assert.Nil(t, err, "Should be nil")
	assert.True(t, len(customers) == 2, "Should have two items")
	assert.Equal(t, testCustomers[0], customers[0], "Should be equal")
	assert.Equal(t, testCustomers[1], customers[1], "Should be equal")
}

func TestNewRequestWithOptions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if json, err := json.Marshal(testCustomers); err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(resp, err.Error())
		} else {
			fmt.Fprintf(resp, string(json))
		}
	}))

	defer ts.Close()

	response := NewRequestBuilder().WithUrl(ts.URL).WithMethod("GET").Build().Do()

	assert.Equal(t, "GET", response.Response().Request.Method, "Should equal GET method")
	assert.Equal(t, 200, response.Response().StatusCode, "Should equal HTTP Status 200 (OK)")

	customers := make([]*TestCustomer, 0)

	err := json.Unmarshal(response.Body(), &customers)

	assert.Nil(t, err, "Should be nil")
	assert.True(t, len(customers) == 2, "Should have two items")
	assert.Equal(t, testCustomers[0], customers[0], "Should be equal")
	assert.Equal(t, testCustomers[1], customers[1], "Should be equal")
}

func TestNewRequestWithOptionsWithoutMethodSpecified(t *testing.T) {
	response := NewRequestBuilder().WithUrl("https://www.google.com").Build().Do()
	assert.Equal(t, "GET", response.Response().Request.Method, "Should equal GET method")
	assert.Equal(t, 200, response.Response().StatusCode, "Should equal HTTP Status 200 (OK)")
	assert.NotEmpty(t, string(response.Body()), "Should not be empty")
}

/*
func TestNewRequestPanicWhenInvalidArgumentType(t *testing.T) {
	defer func() {
		err := recover().(error)
		assert.NotNil(t, err, "Should not be nil")
		assert.Equal(t, "Invalid argument type", err.Error(), "Should equal error message")
	}()

	o := 10

	NewRequest(o)

	assert.True(t, false, "Should not have completed test")
}

func TestNewRequestPanicWhenInvalidStructType(t *testing.T) {
	defer func() {
		err := recover().(error)

		assert.NotNil(t, err, "Should not be nil")
		assert.Equal(t, "Type was *gorequest.TestOption but expected *gorequest.Option", err.Error(), "Should equal error message")
	}()

	o := &TestOption{
		Url:    "https://www.google.com",
		Method: "GET",
	}

	NewRequest(o)

	assert.True(t, false, "Should not have completed test")
}
*/

func TestNewRequestWithoutURL(t *testing.T) {
	defer func() {
		err := recover().(error)

		assert.NotNil(t, err, "Should not be nil")
		assert.Equal(t, "URL is required", err.Error(), "Should equal error message")
	}()

	NewRequestBuilder().Build()

	assert.True(t, false, "Should not have completed test")
}

func TestBasicAuthentication(t *testing.T) {
	response := NewRequestBuilder().WithUrl("https://postman-echo.com/basic-auth").WithBasicAuth("postman", "password").Build().Do()

	basicAuth := base64.StdEncoding.EncodeToString([]byte("postman:password"))
	basicAuth = "Basic " + basicAuth

	assert.Equal(t, "GET", response.Response().Request.Method, "Should equal GET method")
	assert.Equal(t, 200, response.Response().StatusCode, "Should equal HTTP Status 200 (OK)")
	assert.Equal(t, "{\"authenticated\":true}", string(response.Body()), "Should equal body")
	assert.Equal(t, basicAuth, response.Response().Request.Header.Get("Authorization"), "Should equal Authorization header")
}

func TestGetRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if json, err := json.Marshal(testCustomers); err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(resp, err.Error())
		} else {
			fmt.Fprintf(resp, string(json))
		}
	}))

	defer ts.Close()

	response := NewRequestBuilder().WithUrl(ts.URL).Build().Do()

	assert.Equal(t, "GET", response.Response().Request.Method, "Should equal GET method")
	assert.Equal(t, 200, response.Response().StatusCode, "Should equal HTTP Status 200 (OK)")

	customers := make([]*TestCustomer, 0)

	err := json.Unmarshal(response.Body(), &customers)

	assert.Nil(t, err, "Should be nil")
	assert.True(t, len(customers) == 2, "Should have two items")
	assert.Equal(t, testCustomers[0], customers[0], "Should be equal")
	assert.Equal(t, testCustomers[1], customers[1], "Should be equal")
}

func TestPostRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		var customer *TestCustomer

		decoder := json.NewDecoder(req.Body)

		if err := decoder.Decode(&customer); err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(resp, err.Error())
		} else {
			testCustomers = append(testCustomers, customer)

			resp.WriteHeader(http.StatusCreated)
			fmt.Fprintf(resp, "Created")
		}
	}))

	defer ts.Close()

	c1 := &TestCustomer{
		Id:        3,
		FirstName: "PostTest",
		LastName:  "PostTest",
	}

	response := NewRequestBuilder().WithUrl(ts.URL).WithMethod("POST").WithBody(newJsonBody(c1)).Build().Do()

	assert.Equal(t, "POST", response.Response().Request.Method, "Should equal POST method")
	assert.Equal(t, 201, response.Response().StatusCode, "Should equal HTTP Status 201 (Created)")
	assert.Equal(t, "Created", string(response.Body()), "Should equal body")
	assert.Equal(t, "application/json", response.Response().Request.Header.Get("Content-Type"), "Should have set Content-Type to application/json")

	assert.True(t, len(testCustomers) == 3, "Should have three items")
	assert.Equal(t, testCustomers[2], c1, "Should be equal")
}

func TestHeadRequest(t *testing.T) {
	response := NewRequestBuilder().WithUrl("https://www.google.com").WithMethod("HEAD").Build().Do()

	assert.Equal(t, "HEAD", response.Response().Request.Method, "Should equal HEAD method")
	assert.Equal(t, 200, response.Response().StatusCode, "Should equal HTTP Status 200 (OK)")
	assert.Empty(t, string(response.Body()), "Should be empty")
}
