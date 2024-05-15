package main

import (
	"example/go-url/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewShort(t *testing.T) {
	l1 := db.Shorten("https://example.com/")
	assert.Equal(t, "https://example.com/", l1.OG)
}
