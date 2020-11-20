package external

import (
	"encoding/json"
	"net/url"

	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/httpclient"
	"github.com/paraizofelipe/gorecipes/model"
)

const recipeURL string = "http://www.recipepuppy.com/api"

type Recipe struct {
	Logger     echo.Logger
	HTTPClient httpclient.Requester
}

func NewRecipe(logger echo.Logger) *Recipe {
	return &Recipe{
		HTTPClient: httpclient.HTTPClient{},
		Logger:     logger,
	}
}

// Search do requests to API Recipe Puppy
func (r Recipe) Search(ingredients string) (recipes model.APIRecipeResponse, err error) {
	var (
		resp    []byte
		baseURL *url.URL
	)
	if baseURL, err = url.Parse(recipeURL); err != nil {
		return
	}

	params := url.Values{}
	params.Add("i", ingredients)
	baseURL.RawQuery = params.Encode()

	r.Logger.Info(baseURL.String())
	if resp, err = r.HTTPClient.MakeRequest("GET", baseURL.String()); err != nil {
		return
	}
	if err = json.Unmarshal(resp, &recipes); err != nil {
		return
	}
	return
}
