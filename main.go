package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/umanovskis/go-etsy-api/api"
)

func main() {
	key := api.GetApiKey()
	if key == "" {
		fmt.Println("No API key found, type key below:")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			key = scanner.Text()
		}
	}
	ctx := api.CreateApiCtx(key)

	req := ctx.NewListingRequest()
	req.AddKeyword("goat")
	body, err := ctx.Request(req)
	body = body

	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	var response api.ListingResponse
	json.Unmarshal(body, &response)
	/*		for _, l := range response.Items {
				fmt.Println(l.Title + " -- SOLD BY -- " + l.Shop.Name)
				fmt.Println(l.Url)
				fmt.Println(l.User.Feedback)
			}
	*/
	first := response.Items[0]
	fmt.Println("---- FIRST ---- ")
	//fmt.Println(first.Title)
	// fmt.Println(first.Image.Url)

	guest, err := ctx.GenerateGuest()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Got a guest!")
		fmt.Printf("%+v\n", guest)
	}
	fmt.Printf("Guest checkout link is %s\n", guest.CheckoutUrl)
	guest.AddListingToCart(first)
	cart, err := ctx.GetGuestCart(guest)
	if err != nil {
		fmt.Println("Error adding item to cart...")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Got a cart!")
		fmt.Printf("%s", cart)
	}

	/*
		url = url.AddKeyword("goat")
		req, _ = http.NewRequest(http.MethodGet, url, nil)
		resp, _ = client.Do(req)
		body, _ = ioutil.ReadAll(resp.Body)

		json.Unmarshal(body, &response)
		fmt.Printf("%s listings found\n", response.Count)
	*/
}
