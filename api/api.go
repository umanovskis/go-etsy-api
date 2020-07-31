package api

import (
	"fmt"
	"github.com/umanovskis/go-etsy-api/http"
	"io/ioutil"
	"os/user"
	"strings"
)

const urlbase = "https://openapi.etsy.com/v2/"

type EtsyApi struct {
	apiKey string
}

type Requester interface {
	Url() string
}

type Poster interface {
	Url() string
	Data() []byte
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

func (e *EtsyApi) Request(r Requester) ([]byte, error) {
	return http.HttpRequest(e.authenticate(r.Url()))
}

func (e *EtsyApi) Post(p Poster) ([]byte, error) {
	fmt.Println(p.Url())
	return http.HttpPost(e.authenticate(p.Url()), p.Data())
}
