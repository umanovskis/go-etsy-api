package api

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"strings"
)

const urlbase = "https://openapi.etsy.com/v2/"

type EtsyApi struct {
	apiKey string
}

type RequestUrl interface {
	Url()
}

func CreateApiCtx(apiKey string) *EtsyApi {
	return &EtsyApi{apiKey: apiKey}
}

func GetApiKey() string {
	user, _ := user.Current()
	home := user.HomeDir
	file := fmt.Sprintf("%s/.etsyapikey", home)
	key, err := ioutil.ReadFile(file)

	if err != nil {
		return ""
	}
	return strings.Trim(string(key), "\n\r")
}

func (e *EtsyApi) authenticate(url string) string {
	if !strings.Contains(url, "?") {
		return url + "?api_key=" + e.apiKey
	}
	return url + "&api_key=" + e.apiKey
}
