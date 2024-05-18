package libraries

import (
	u "example/go-short/libraries/util"

	"time"

	verify "github.com/davidmytton/url-verifier"
	nano "github.com/jaevor/go-nanoid"
)

type Link struct {
	ID        string        `bson:"id" json:"id,omitempty"`
	ShortLink string        `bson:"shortlink" json:"shortlink,omitempty"`
	Exp       time.Duration `bson:"exp" json:"exp,omitempty"`
}

// using SHA1 to encode link into hash
func EncodeSHA(s string) string {
	// h := sha1.New()
	// h.Write([]byte(s))
	// sha := base64.URLEncoding.EncodeToString(h.Sum(nil))[:5] //slice first 5

	sha, e := nano.Standard(5)
	uid := sha()
	u.Check(e)

	return uid
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
