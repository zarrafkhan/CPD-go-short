package libraries

import (
	u "example/go-short/libraries/util"
	"net/url"

	nano "github.com/jaevor/go-nanoid"
)

type Link struct {
	ID        string `bson:"id" json:"id,omitempty"`
	ShortLink string `bson:"shortlink" json:"shortlink,omitempty"`
	// Exp       time.Duration `bson:"exp" json:"exp,omitempty"`
}

// encode link into hash
func EncodeSHA(s string) string {

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
	check := true

	_, err := url.ParseRequestURI(link)
	if err != nil {
		check = false
	}

	return check
}
