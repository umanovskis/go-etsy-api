package api

import (
	"encoding/json"
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

type listingRequest struct {
	parameters url.Values
	ctx        *EtsyApi
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
	old := r.parameters.Get("keywords")
	if old != "" {
		r.parameters.Set("keywords", old+","+keyword)
	} else {
		r.parameters.Set("keywords", keyword)
	}
}

func (e *EtsyApi) NewListingRequest() *listingRequest {
	return &listingRequest{parameters: url.Values{}, ctx: e}
}

func (l *listingRequest) Url() string {
	return urlbase + "/listings/active/?" + l.parameters.Encode()
}

func (l *listingRequest) Execute() (ListingResponse, error) {
	bytes, err := l.ctx.Request(l)

	if err != nil {
		return ListingResponse{}, err
	}
	var response ListingResponse
	err = json.Unmarshal(bytes, &response)

	if err != nil {
		return ListingResponse{}, err
	}

	return response, nil
}
