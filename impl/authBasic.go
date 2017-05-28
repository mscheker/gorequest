package gorequest

import (
	model "gorequest/model"
	"net/http"
)

/****************************************************
 * model.AuthorizationMethod implementation
 ****************************************************/

type authBasic struct {
	password string
	user string
}

func newAuthBasic(user string, password string) model.AuthorizationMethod {
	return &authBasic{
		user: user,
		password: password,
	}
}

func (a *authBasic) Configure(request *http.Request) {
	request.SetBasicAuth(a.user, a.password)
}
