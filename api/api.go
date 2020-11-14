package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/paraizofelipe/gorecipes/model"
	"github.com/paraizofelipe/gorecipes/settings"
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
	var baseURL *url.URL
	if baseURL, err = url.Parse(recipeURL); err != nil {
		return
	}
	params := url.Values{}
	params.Add("i", ingredients)
	baseURL.RawQuery = params.Encode()
	if err = makeRequest("GET", baseURL.String(), &recipes); err != nil {
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
	params.Add("api_key", settings.GyphToken)
	params.Add("limit", "1")
	params.Add("q", strings.TrimSpace(title))
	baseURL.RawQuery = params.Encode()

	if err = makeRequest("GET", baseURL.String(), &resp); err != nil {
		return
	}
	return resp["data"].([]interface{})[0].(map[string]interface{})["url"].(string), nil
}

func AsyncSearchRecipes(ingredients string) (ch chan model.APIRecipeResponse, err error) {
	go func() {
		defer close(ch)
		recipe, err := SearchRecipes(ingredients)
		if err != nil {
			log.Println(err)
			return
		}
		ch <- recipe
	}()

	return
}

func AsyncSearchGif(title string) (ch chan string, err error) {
	go func() {
		defer close(ch)
		gif, err := SearchGif(title)
		if err != nil {
			log.Println(err)
			return
		}
		ch <- gif
	}()
	return
}