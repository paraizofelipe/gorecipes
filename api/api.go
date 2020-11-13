package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/paraizofelipe/gorecipes/model"
)

const (
	recipeURL string = "http://www.recipepuppy.com/api"
	gifURL    string = "https://api.giphy.com/v1/gifs/search"
)

func makeRequest(method string, url string, result interface{}) error {
	client := http.Client{}
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return errors.New(resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}
	return err
}

// SearchRecipes ---
func SearchRecipes(ingredients string) (recipes model.APIRecipeResponse, err error) {
	url := fmt.Sprintf("%s?i=%s", recipeURL, ingredients)
	if err = makeRequest("GET", url, &recipes); err != nil {
		return
	}
	return
}

// SearchGif ---
func SearchGif(title string) (gif string, err error) {
	var (
		baseURL *url.URL
		resp    map[string]interface{}
	)

	if baseURL, err = url.Parse(gifURL); err != nil {
		return
	}
	params := url.Values{}
	params.Add("api_key", "pPiMNFkdnBt4wGmBiJ9YCryAw3lHJk98")
	params.Add("limit", "1")
	params.Add("q", strings.TrimSpace(title))
	baseURL.RawQuery = params.Encode()

	if err = makeRequest("GET", baseURL.String(), &resp); err != nil {
		return
	}
	gif = resp["data"].([]interface{})[0].(map[string]interface{})["url"].(string)
	return
}
