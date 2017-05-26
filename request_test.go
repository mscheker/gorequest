package gorequest

import (
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

	resp, body, err := NewRequest(ts.URL)

	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, "GET", resp.Request.Method, "Should equal GET method")
	assert.Equal(t, 200, resp.StatusCode, "Should equal HTTP Status 200 (OK)")

	customers := make([]*TestCustomer, 0)

	err = json.Unmarshal(body, &customers)

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

	options := &Option{
		Url:    ts.URL,
		Method: "GET",
	}

	resp, body, err := NewRequest(options)

	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, "GET", resp.Request.Method, "Should equal GET method")
	assert.Equal(t, 200, resp.StatusCode, "Should equal HTTP Status 200 (OK)")

	customers := make([]*TestCustomer, 0)

	err = json.Unmarshal(body, &customers)

	assert.Nil(t, err, "Should be nil")
	assert.True(t, len(customers) == 2, "Should have two items")
	assert.Equal(t, testCustomers[0], customers[0], "Should be equal")
	assert.Equal(t, testCustomers[1], customers[1], "Should be equal")
}

func TestNewRequestWithOptionsWithoutMethodSpecified(t *testing.T) {
	o := &Option{
		Url: "https://www.google.com",
	}

	resp, body, err := NewRequest(o)

	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, "GET", resp.Request.Method, "Should equal GET method")
	assert.Equal(t, 200, resp.StatusCode, "Should equal HTTP Status 200 (OK)")
	assert.NotEmpty(t, string(body), "Should not be empty")
}

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

func TestNewRequestWithoutURL(t *testing.T) {
	defer func() {
		err := recover().(error)

		assert.NotNil(t, err, "Should not be nil")
		assert.Equal(t, "URL is required", err.Error(), "Should equal error message")
	}()

	o := &Option{
		Url: "",
	}

	NewRequest(o)

	assert.True(t, false, "Should not have completed test")
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

	options := &Option{
		Url: ts.URL,
	}

	resp, body, err := Get(options)

	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, "GET", resp.Request.Method, "Should equal GET method")
	assert.Equal(t, 200, resp.StatusCode, "Should equal HTTP Status 200 (OK)")

	customers := make([]*TestCustomer, 0)

	err = json.Unmarshal(body, &customers)

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

	options := &Option{
		Url:  ts.URL,
		JSON: c1,
	}

	resp, body, err := Post(options)

	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, "POST", resp.Request.Method, "Should equal POST method")
	assert.Equal(t, 201, resp.StatusCode, "Should equal HTTP Status 201 (Created)")
	assert.Equal(t, "Created", string(body), "Should equal body")
	assert.Equal(t, "application/json", options.Headers["Content-Type"], "Should have set Content-Type to application/json")

	assert.True(t, len(testCustomers) == 3, "Should have three items")
	assert.Equal(t, testCustomers[2], c1, "Should be equal")
}

func TestHeadRequest(t *testing.T) {
	options := &Option{
		Url:    "https://www.google.com",
		Method: "HEAD",
	}

	resp, body, err := NewRequest(options)

	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, "HEAD", resp.Request.Method, "Should equal HEAD method")
	assert.Equal(t, 200, resp.StatusCode, "Should equal HTTP Status 200 (OK)")
	assert.Empty(t, string(body), "Should be empty")
}
