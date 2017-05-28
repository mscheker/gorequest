package gorequest

import (
	"fmt"
	model "github.com/demianlessa/gorequest/model"
	"net/http"
)

/****************************************************
 * model.AuthorizationMethod implementation
 ****************************************************/

type authBearer struct {
	token string
}

func newAuthBearer(token string) model.AuthorizationMethod {
	return &authBearer{
		token: token,
	}
}

func (a *authBearer) Configure(request *http.Request) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.token))
}
