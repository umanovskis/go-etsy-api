package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ctx = CreateApiCtx("testkey")

func TestNewListing(t *testing.T) {
	l := ctx.NewListingRequest()
	assert.Equal(t, urlbase+"/listings/active/?includes=Shop%2CUser%2CMainImage", l.Url())

	l2 := ctx.NewListingRequest()
	assert.Equal(t, l.Url(), l2.Url())
}

func TestAddKeyword(t *testing.T) {
	l := ctx.NewListingRequest()
	l.AddKeyword("goats")
	assert.Contains(t, l.Url(), "keywords=goats")
}

func TestMultipleKeywordsHaveCommas(t *testing.T) {
	l := ctx.NewListingRequest()
	l.AddKeyword("goats")
	l.AddKeyword("sheep")
	joined := "keywords=goats%2Csheep"
	assert.Contains(t, l.parameters.Encode(), joined)
}

var l1 = Listing{Id: 2, State: "active", Title: "chocolate bar", Description: "tastes good"}
var l2 = Listing{Id: 5, State: "active", Title: "antique ps/2 cable", Description: "does not taste good"}
var list = Listings{l1, l2}

type FilterAcceptAll struct{}
type FilterAcceptNone struct{}
type FilterAcceptId5 struct{}

func (f FilterAcceptAll) Accept(l Listing) bool {
	return true
}

func (f FilterAcceptNone) Accept(l Listing) bool {
	return false
}

func (f FilterAcceptId5) Accept(l Listing) bool {
	return l.Id == 5
}

func TestListingFilters(t *testing.T) {
	old := list
	filtered := old.Filter(FilterAcceptAll{})
	assert.Equal(t, old, filtered)

	filtered = old.Filter(FilterAcceptNone{})
	assert.Empty(t, filtered)

	filtered = old.Filter(FilterAcceptId5{})
	assert.Len(t, filtered, 1)
	assert.Equal(t, filtered[0], l2)
}
