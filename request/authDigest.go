package request

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type authDigest struct {
	parts map[string]string
}

func newAuthDigest(username, password string) AuthorizationMethod {
	digest := &authDigest{
		parts: make(map[string]string),
	}
	digest.parts["username"] = username
	digest.parts["password"] = password

	return digest
}

func (a *authDigest) Configure(request *http.Request) {
	h1 := getMD5(fmt.Sprintf("%s:%s:%s", a.parts["username"], a.parts["realm"], a.parts["password"]))
	h2 := getMD5(fmt.Sprintf("%s:%s", a.parts["method"], a.parts["uri"]))

	nonceCount := 00000001
	nonce := getNOnce()

	response := getMD5(fmt.Sprintf("%s:%s:%v:%s:%s:%s", h1, a.parts["nonce"], nonceCount, nonce, a.parts["qop"], h2))
	authorization := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc="%v", qop="%s", response="%s"`, a.parts["username"], a.parts["realm"], a.parts["nonce"], a.parts["uri"], nonce, nonceCount, a.parts["qop"], response)

	request.Header.Add("Authorization", authorization)

}

func (a *authDigest) setDigestParts(response *http.Response) {
	wwwAuth := response.Header.Get("WWW-Authenticate")

	a.parts["realm"] = getDigestPart("realm", wwwAuth)
	a.parts["nonce"] = getDigestPart("nonce", wwwAuth)
	a.parts["qop"] = getDigestPart("qop", wwwAuth)
	a.parts["uri"] = response.Request.URL.Path
	a.parts["method"] = response.Request.Method
}

func (a *authDigest) getScheme() authScheme {
	return AUTH_DIGEST
}

func getDigestPart(part, wwwAuth string) string {
	v := ""
	headerParts := strings.Split(wwwAuth, ",")

	for _, r := range headerParts {
		if strings.Contains(r, part) {
			v = strings.Split(r, `"`)[1]
		}
	}

	return v
}

func getMD5(v string) string {
	h := md5.New()
	h.Write([]byte(v))

	return hex.EncodeToString(h.Sum(nil))
}

func getNOnce() string {
	b := make([]byte, 8)
	io.ReadFull(rand.Reader, b)

	return fmt.Sprintf("%x", b)[:16]
}
