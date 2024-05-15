// business logic
package internals

import (
	uniqueid "github.com/albinj12/unique-id"
	//"go.mongodb.org/mongo-driver/mongo"
	// urlverifier "github.com/davidmytton/url-verifier"
)

// ID gen
var og, _ = uniqueid.Generateid("a")
var sh, _ = uniqueid.Generateid("a", 5)

type Url struct {
	ID   string `json:"id"`
	SH   string `json:"sh"`
	OG   string `json:"og"`
	Flag bool   `json:"flag"`
}

type Store interface {
	SaveURL(u Url) error //void
	GetID(ID string) (*Url, error)
	GetShID(SH string) (*Url, error)
}

func Shorten(link string) Url {
	return Url{
		ID:   og,
		SH:   sh,
		OG:   link,
		Flag: true,
	}
}
