package handler

import (
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/externalapi"
	"github.com/paraizofelipe/gorecipes/model"
)

type RecipeResponse struct {
	Keywords []string       `json:"keywords"`
	Recipes  []model.Recipe `json:"recipes"`
}

// APIRecipeToRecipe ---
func (h Handler) APIRecipeToRecipe(apiRecipes []model.APIRecipe) (recipes []model.Recipe, err error) {
	var (
		wg sync.WaitGroup
	)

	recipes = make([]model.Recipe, len(apiRecipes), len(apiRecipes))

	for index, apiRecipe := range apiRecipes {
		recipes[index].Title = apiRecipe.Title
		recipes[index].Link = apiRecipe.Href

		recipes[index].Ingredients = strings.Split(apiRecipe.Ingredients, ",")
		sort.Strings(recipes[index].Ingredients)

		wg.Add(1)
		go externalapi.AsyncSearchGif(recipes[index].Title, &wg, &recipes[index].Gif)
	}
	wg.Wait()

	return
}

// sls ---
func (h Handler) Recipes(c echo.Context) (err error) {
	var (
		recipes    []model.Recipe
		respRecipe model.APIRecipeResponse
	)

	h.Logger.Info("Teste")
	ingredients := c.QueryParam("i")
	if respRecipe, err = externalapi.SearchRecipes(ingredients); err != nil {
		h.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	if len(respRecipe.Results) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "recipe not found"})
	}
	if recipes, err = h.APIRecipeToRecipe(respRecipe.Results); err != nil {
		h.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, recipes)
}
