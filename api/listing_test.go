package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewListing(t *testing.T) {
	l := NewListingRequest()
	assert.Equal(t, urlbase+"/listings/active/?", l.Url())

	l2 := NewListingRequest()
	assert.Equal(t, l.Url(), l2.Url())
}

func TestAddKeyword(t *testing.T) {
	l := NewListingRequest()
	l.AddKeyword("goats")
	assert.Contains(t, l.Url(), "keywords=goats")
}

func TestMultipleKeywordsHaveCommas(t *testing.T) {
	l := NewListingRequest()
	l.AddKeyword("goats")
	l.AddKeyword("sheep")
	joined := "keywords=goats%2Csheep"
	assert.Contains(t, l.parameters.Encode(), joined)
}
