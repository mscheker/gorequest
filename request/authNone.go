package request

import "net/http"

type authNone struct {
}

func newAuthNone() AuthorizationMethod {
	return &authNone{}
}

func (a *authNone) Configure(request *http.Request) {

}
