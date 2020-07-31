package http

import (
	"bytes"
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

func HttpPost(url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
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
