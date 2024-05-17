package main

import (
	"testing"

	i "example/go-short/internals"

	"github.com/stretchr/testify/assert"
)

func TestShortenLink(t *testing.T) {
	l1 := i.SetLink("https://example.com/")
	assert.Equal(t, "https://example.com/", l1.ID)
}
