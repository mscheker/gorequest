package gorequest

import (
	model "github.com/demianlessa/gorequest/model"
	"net/http"
)

/****************************************************
 * model.AuthorizationMethod implementation
 ****************************************************/

type authNone struct {
}

func newAuthNone() model.AuthorizationMethod {
	return &authNone{}
}

func (a *authNone) Configure(request *http.Request) {
}
