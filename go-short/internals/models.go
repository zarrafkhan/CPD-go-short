package internals

import (
	utils "example/go-short/internals/util"

	"fmt"
	"time"

	uid "github.com/albinj12/unique-id"
	verify "github.com/davidmytton/url-verifier"
)

type Link struct {
	ID        string        `json:"id,omitempty" bson: "id"`
	ShortLink string        `json:"sh, omitempty" bson: "sh"`
	Exp       time.Duration `json: "exp" "bson: "exp"`
}

func ShortenLink(link string) string {
	sh, _ := uid.Generateid("a", 5)
	fmt.Println(sh)
	return sh
}

func SetLink(link string) Link {
	return Link{
		ID:        link,
		ShortLink: ShortenLink(link),
	}
}

func VerifyLink(link string) bool {
	check := false
	v := verify.NewVerifier()
	v.EnableHTTPCheck()
	ret, e := v.Verify(link)
	utils.Check(e)

	if ret.HTTP.IsSuccess {
		check = true
	}

	return check
}
