package api

import (
	"net/url"
)

type Listing struct {
	Id          int    `json:"id"`
	State       string `json:"state"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type Listings []Listing

type ListingResponse struct {
	Count int      `json:"count"`
	Items Listings `json:"results"`
}

func (l Listings) GetActiveListings() Listings {
	var result []Listing
	for i := range l {
		if l[i].State == "active" {
			result = append(result, l[i])
		}
	}
	return result
}

func (r *listingRequest) AddKeyword(keyword string) {
	r.parameters.Add("keywords", keyword)
}

type listingRequest struct {
	parameters url.Values
}

func NewListingRequest() *listingRequest {
	return &listingRequest{parameters: url.Values{}}
}

func (l *listingRequest) Url() string {
	return urlbase + "/listings/active/?" + l.parameters.Encode()
}
