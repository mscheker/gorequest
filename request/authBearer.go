package request

import (
	"fmt"
	"net/http"
)

type authBearer struct {
	token string
}

func newAuthBearer(token string) AuthorizationMethod {
	return &authBearer{
		token: token,
	}
}

func (a *authBearer) Configure(request *http.Request) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.token))
}

func (a *authBearer) getScheme() authScheme {
	return AUTH_BEARER
}
