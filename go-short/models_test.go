package main

import (
	"fmt"
	"testing"

	I "example/go-short/libraries"

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

func TestDeleteURL(t *testing.T) {
	f := I.DeletURL(collection, "www.google.com")
	assert.Nil(t, f)
}

func TestVerify(t *testing.T) {
	l1 := I.VerifyLink("abc")
	fmt.Println(l1)
	assert.False(t, l1)
}
