package request

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type request struct {
	request *http.Request
	client  *http.Client
	auth    AuthorizationMethod
}

func newRequest(req *http.Request, client *http.Client, auth AuthorizationMethod) Request {
	return &request{
		request: req,
		client:  client,
		auth:    auth,
	}
}

func (r *request) getUnderlyingRequest() *http.Request {
	return r.request
}

func (r *request) getUnderlyingHttpClient() *http.Client {
	return r.client
}

func (r *request) Do() Response {
	// REMARKS: Delay auth configuration when set to Digest
	if r.auth.getScheme() != AUTH_DIGEST {
		r.auth.Configure(r.request)
	}

	resp, err := r.client.Do(r.request)

	if err != nil {
		panic(err)
	}

	// REMARKS: If we received a 401, check if we need to perform digest auth and resend the request.
	if resp.StatusCode == http.StatusUnauthorized {

		h := resp.Header.Get("WWW-Authenticate")

		authParts := strings.Split(h, " ")

		if len(authParts) > 0 && authParts[0] == "Digest" {
			a, ok := r.auth.(*authDigest)
			if ok {
				a.setDigestParts(resp)
				a.Configure(r.request)

				// REMARKS: Digest authentication has been configured so resend the request.
				if resp, err = r.client.Do(r.request); err != nil {
					panic(err)
				}
			}
		}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return &response{
		body:     body,
		response: resp,
	}
}
