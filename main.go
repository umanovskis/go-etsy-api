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

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	var response api.ListingResponse
	json.Unmarshal(body, &response)
	for _, l := range response.Items {
		fmt.Println(l.Title + " -- SOLD BY -- " + l.Shop.Name)
                fmt.Println(l.Url)
                fmt.Println(l.User.Feedback)
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
