package main

import (
	i "example/go-short/internals"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewShort(t *testing.T) {
	l1 := i.Shorten("https://example.com/")
	assert.Equal(t, "https://example.com/", l1.OG)
}
