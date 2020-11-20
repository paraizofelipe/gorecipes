package external

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/httpclient"
	"github.com/paraizofelipe/gorecipes/settings"
)

const gifURL string = "https://api.giphy.com/v1/gifs/search"

type Gif struct {
	Logger     echo.Logger
	HTTPClient httpclient.Requester
}

func NewGif(logger echo.Logger) *Gif {
	return &Gif{
		HTTPClient: httpclient.HTTPClient{},
		Logger:     logger,
	}
}

// Search do requests to API Giphy
func (g Gif) Search(title string) (gif string, err error) {
	var (
		result  map[string]interface{}
		baseURL *url.URL
		resp    []byte
	)

	if baseURL, err = url.Parse(gifURL); err != nil {
		return
	}

	params := url.Values{}
	params.Add("api_key", settings.GiphyToken)
	params.Add("limit", "1")
	params.Add("q", strings.TrimSpace(title))
	baseURL.RawQuery = params.Encode()

	g.Logger.Info(baseURL.String())
	if resp, err = g.HTTPClient.MakeRequest("GET", baseURL.String()); err != nil {
		return
	}
	if err = json.Unmarshal(resp, &result); err != nil {
		return
	}
	return result["data"].([]interface{})[0].(map[string]interface{})["url"].(string), nil
}
