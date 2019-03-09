package api

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Listing struct {
	Id          int       `json:"id"`
	State       string    `json:"state"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	Price       string    `json:"price"`
	Currency    string    `json:"currency"`
	Shop        Shop      `json:"Shop"`
	User        User      `json:"User"`
	Image       MainImage `json:"MainImage"`
}

type Listings []Listing

type ListingResponse struct {
	Count int      `json:"count"`
	Items Listings `json:"results"`
}

type Accepter interface {
	Accept(Listing) bool
}

type listingRequest struct {
	parameters url.Values
	ctx        *EtsyApi
}

func (l Listings) Filter(filter Accepter) Listings {
	var result []Listing
	for i := range l {
		if filter.Accept(l[i]) {
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

func (r *listingRequest) SetPageNumber(page int) {
	r.parameters.Set("page", strconv.Itoa(page))
}

func (r *listingRequest) SetLimit(limit int) {
	r.parameters.Set("limit", strconv.Itoa(limit))
}

func (e *EtsyApi) NewListingRequest() *listingRequest {
	return &listingRequest{parameters: url.Values{"includes": []string{"Shop,User,MainImage"}}, ctx: e}
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

func (l *Listing) GetFeedback() FeedbackInfo {
	return l.User.Feedback
}
