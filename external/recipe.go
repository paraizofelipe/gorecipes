package external

import (
	"net/url"
	"sync"

	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/model"
)

type APIRecipe struct {
	Logger echo.Logger
}

func NewAPIRecipe(logger echo.Logger) APIRecipe {
	return APIRecipe{
		Logger: logger,
	}
}

// Search do requests to API Recipe Puppy
func (a APIRecipe) Search(ingredients string) (recipes model.APIRecipeResponse, err error) {
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

// AsyncSearch do asynchronous requests to API Recipe Puppy
func (a APIRecipe) AsyncSearch(ingredients string, wg *sync.WaitGroup, resultRecipe *model.APIRecipeResponse) {
	var err error

	defer wg.Done()
	if *resultRecipe, err = a.Search(ingredients); err != nil {
		a.Logger.Error(err)
		return
	}
}
