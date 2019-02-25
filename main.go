package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/umanovskis/go-etsy-api/api"
	"os"
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

	req := api.NewListingRequest()
	body, err := ctx.HttpRequest(req.Url())

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	var response api.ListingResponse
	json.Unmarshal(body, &response)
	fmt.Printf("%s listings total\n", response.Count)
	for _, l := range response.Items {
		fmt.Println(l.Title + " ---- " + l.Url)
	}
	active := response.Items.GetActiveListings()
	fmt.Printf("%s active listings decoded\n", len(active))
	req.AddKeyword("goat")
	fmt.Println(req.Url())

	body, err = ctx.HttpRequest(req.Url())
	json.Unmarshal(body, &response)
	for _, l := range response.Items {
		fmt.Println(l.Title + " ---- " + l.State)
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
