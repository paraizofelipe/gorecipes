package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/mocks"
	"github.com/paraizofelipe/gorecipes/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRecipes(t *testing.T) {

	var (
		expected = `{"keywords":["tomato","orange","onion"],"recipes":[{"title":"Tomato \u0026 Orange Cottage Cheese Salad ","ingredients":["balsamic vinaigrette","basil","cottage cheese","orange","red onions","tomato"],"link":"http://www.kraftfoods.com/kf/recipes/tomato-orange-cottage-cheese-54326.aspx","gif":"http://giphy.com"}]}`

		respRecipe model.APIRecipeResponse = model.APIRecipeResponse{
			Title:   "Recipe Puppy",
			Version: 0.1,
			Href:    "http://www.recipepuppy.com/",
			Results: []model.APIRecipe{
				{
					Title:       "Tomato & Orange Cottage Cheese Salad ",
					Href:        "http://www.kraftfoods.com/kf/recipes/tomato-orange-cottage-cheese-54326.aspx",
					Ingredients: "tomato, orange, cottage cheese, red onions, basil, balsamic vinaigrette",
					Thumbnail:   "http://img.recipepuppy.com/636990.jpg",
				},
			},
		}

		mockedRecipeSearcher = &mocks.RecipeSearcher{}
		mockedGifSearcher    = &mocks.GifSearcher{}

		e *echo.Echo = echo.New()
	)

	recipeHandler := Recipe{
		Logger:         e.Logger,
		RecipeSearcher: mockedRecipeSearcher,
		GifSearcher:    mockedGifSearcher,
	}

	t.Run("should return a recipe JSON as a response", func(t *testing.T) {
		mockedRecipeSearcher.Mock = mock.Mock{}
		mockedGifSearcher.Mock = mock.Mock{}

		mockedRecipeSearcher.On("Search", "tomato,orange,onion").Return(respRecipe, nil)
		mockedGifSearcher.On("Search", "Tomato & Orange Cottage Cheese Salad ").Return("http://giphy.com", nil)

		q := make(url.Values)
		q.Set("i", "tomato,orange,onion")
		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, recipeHandler.GetRecipes(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expected+"\n", rec.Body.String())
		}
	})

	t.Run("should return not found error", func(t *testing.T) {
		mockedRecipeSearcher.Mock = mock.Mock{}

		mockedRecipeSearcher.On("Search", "xxxx,xxxx").Return(model.APIRecipeResponse{}, nil)

		q := make(url.Values)
		q.Set("i", "xxxx,xxxx")
		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, recipeHandler.GetRecipes(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Contains(t, rec.Body.String(), "recipe not found")
		}
	})

	t.Run("should return internal server error", func(t *testing.T) {
		mockedRecipeSearcher.Mock = mock.Mock{}

		mockedRecipeSearcher.On("Search", "xxxx,xxxx").Return(model.APIRecipeResponse{}, errors.New(""))

		q := make(url.Values)
		q.Set("i", "xxxx,xxxx")
		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, recipeHandler.GetRecipes(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Contains(t, rec.Body.String(), "internal server error")
		}
	})
}
