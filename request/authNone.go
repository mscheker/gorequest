package request

import "net/http"

type authNone struct {
}

func newAuthNone() AuthorizationMethod {
	return &authNone{}
}

func (a *authNone) Configure(request *http.Request) {

}

func (a *authNone) getScheme() authScheme {
	return AUTH_NONE
}
