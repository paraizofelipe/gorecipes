package external

import (
	"net/url"
	"strings"
	"sync"

	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/settings"
)

type APIGif struct {
	Logger echo.Logger
}

func NewAPIGif(logger echo.Logger) APIGif {
	return APIGif{
		Logger: logger,
	}
}

// Search do requests to API Giphy
func (a APIGif) Search(title string) (gif string, err error) {
	var (
		baseURL *url.URL
		resp    map[string]interface{}
	)

	if baseURL, err = url.Parse(gifURL); err != nil {
		return
	}

	params := url.Values{}
	params.Add("api_key", settings.GiphyToken)
	params.Add("limit", "1")
	params.Add("q", strings.TrimSpace(title))
	baseURL.RawQuery = params.Encode()

	if err = makeRequest("GET", baseURL.String(), &resp); err != nil {
		return
	}
	return resp["data"].([]interface{})[0].(map[string]interface{})["url"].(string), nil
}

// AsyncSearch do asynchronous requests to API Giphy
func (a APIGif) AsyncSearch(title string, wg *sync.WaitGroup, resultGif *string) {
	var err error

	defer wg.Done()
	if *resultGif, err = a.Search(title); err != nil {
		a.Logger.Error(err)
		return
	}
}
