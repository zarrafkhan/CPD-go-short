package main

import (
	"example/go-short/internals/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewShort(t *testing.T) {
	l1 := models.Shorten("https://example.com/")
	assert.Equal(t, "https://example.com/", l1.OG)
}
