package main

import (
	"testing"

	I "example/go-short/internals"

	"github.com/stretchr/testify/assert"
)

var Client, collection = I.SetupMongo()

func TestShortenLink(t *testing.T) {
	l1 := I.SetLink("https://example.com/")
	assert.Equal(t, "https://example.com/", l1.ID)
}

func TestFindURL(t *testing.T) {
	f, _ := I.GetLinkFromShort(collection, "MFaJz")
	assert.Equal(t, "www.google.com", f)
}
