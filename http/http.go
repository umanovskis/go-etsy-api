package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpRequest(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
        fmt.Println(url)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
		return nil, fmt.Errorf("HTTP request returned status %d", resp.StatusCode)
	}

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
