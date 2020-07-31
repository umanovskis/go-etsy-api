package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Guest struct {
	Id          string `json:"guest_id"`
	CheckoutUrl string `json:"checkout_url"`
	ApiCtx      *EtsyApi
}

type GuestResponse struct {
	Count  int     `json:"count"`
	Guests []Guest `json:"results"`
}

type guestGeneratorRequester struct{}

type guestCartGeneratorRequester struct {
	Guest Guest
}

type GuestCart struct {
	Id      int    `json:"guest_id"`
	Total   string `json:"total"`
	GuestId string
	ApiCtx  *EtsyApi
}

type GuestCartResponse struct {
	Count int         `json:"count"`
	Carts []GuestCart `json:"results"`
}

type AddListingToCartData struct {
	GuestId   string `json:"guest_id"`
	ListingId int    `json:"listing_id"`
}

func (e *EtsyApi) GenerateGuest() (Guest, error) {
	g := guestGeneratorRequester{}
	bytes, err := e.Request(g)
	if err != nil {
		return Guest{}, err
	}
	s := string(bytes)
	fmt.Printf(s)
	var response GuestResponse
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return Guest{}, err
	}
	if len(response.Guests) == 0 {
		return Guest{}, errors.New("Empty list of guests returned!")
	}
	res := response.Guests[0]
	res.ApiCtx = e
	return res, nil
}

func (g guestGeneratorRequester) Url() string {
	return urlbase + "/guests/generator"
}

func (r guestCartGeneratorRequester) Url() string {
	return urlbase + fmt.Sprintf("/guests/%s/carts", r.Guest.Id)
}

func (e *EtsyApi) GetGuestCart(g Guest) (GuestCart, error) {
	r := guestCartGeneratorRequester{Guest: g}
	bytes, err := e.Request(r)
	if err != nil {
		return GuestCart{}, err
	}
	var response GuestCartResponse
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return GuestCart{}, err
	}
	if len(response.Carts) == 0 {
		return GuestCart{}, errors.New("Empty list of carts returned!")
	}
	cart := response.Carts[0]
	cart.GuestId = g.Id
	cart.ApiCtx = e
	return response.Carts[0], nil
}

func (cart *Guest) AddListingToCart(l Listing) {
	data := AddListingToCartData{ListingId: l.Id, GuestId: cart.Id}
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	p := addToCartPoster{PostData: bytes, GuestId: cart.Id}
	response, err := cart.ApiCtx.Post(&p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(response))

}

type addToCartPoster struct {
	PostData []byte
	GuestId  string
}

func (p *addToCartPoster) Url() string {
	return urlbase + fmt.Sprintf("/guests/%s/carts", p.GuestId)
}

func (p *addToCartPoster) Data() []byte {
	return p.PostData
}
