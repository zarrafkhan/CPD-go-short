package internals

import (
	"crypto/sha1"
	"encoding/base64"
	u "example/go-short/internals/util"

	"time"

	verify "github.com/davidmytton/url-verifier"
)

const prefix = "go-sh/"

type Link struct {
	ID        string        `json:"id,omitempty" bson: "id"`
	ShortLink string        `json:"sh, omitempty" bson: "sh"`
	Exp       time.Duration `json: "exp" "bson: "exp"`
}

// using SHA1 to encode link into hash
func EncodeSHA(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))[:5] //slice first 5
	return sha
}

func SetLink(link string) Link {
	return Link{
		ID:        link,
		ShortLink: EncodeSHA(link),
	}
}

func VerifyLink(link string) bool {
	check := false
	v := verify.NewVerifier()
	v.EnableHTTPCheck()
	ret, e := v.Verify(link)
	u.Check(e)

	if ret.HTTP.IsSuccess {
		check = true
	}

	return check
}
