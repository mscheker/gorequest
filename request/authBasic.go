package request

import "net/http"

type authBasic struct {
	username string
	password string
}

func newAuthBasic(username, password string) AuthorizationMethod {
	return &authBasic{
		username: username,
		password: password,
	}
}

func (a *authBasic) Configure(request *http.Request) {
	request.SetBasicAuth(a.username, a.password)
}

func (a *authBasic) getScheme() authScheme {
	return AUTH_BASIC
}
