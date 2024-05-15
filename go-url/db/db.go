package db

import (
	"log"

	uniqueid "github.com/albinj12/unique-id"
)

// ID gen
var og, _ = uniqueid.Generateid("a")
var sh, _ = uniqueid.Generateid("a", 5)

type Url struct {
	ID   string
	SH   string
	OG   string
	Flag bool
}

type Store interface {
	SaveURL(u Url) error //void
	GetID(ID string) (*Url, error)
	GetShID(SH string) (*Url, error)
}

func Shorten(full string) Url {
	return Url{
		ID:   og,
		SH:   sh,
		OG:   full,
		Flag: true,
	}
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
