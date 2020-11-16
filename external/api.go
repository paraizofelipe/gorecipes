package external

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

const (
	recipeURL string = "http://www.recipepuppy.com/api"
	gifURL    string = "https://api.giphy.com/v1/gifs/search"
)

// ExternalAPI ---
type API struct {
	Recipe APIRecipe
	Gif    APIGif
}

func New(logger echo.Logger) API {
	return API{
		Recipe: NewAPIRecipe(logger),
		Gif:    NewAPIGif(logger),
	}
}

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
