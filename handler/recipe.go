package handler

import (
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/external"
	"github.com/paraizofelipe/gorecipes/model"
)

type Recipe struct {
	Logger         echo.Logger
	GifSearcher    external.GifSearcher
	RecipeSearcher external.RecipeSearcher
}

func NewRecipe(logger echo.Logger) RecipeHandler {
	return Recipe{
		Logger:         logger,
		GifSearcher:    external.NewGif(logger),
		RecipeSearcher: external.NewRecipe(logger),
	}
}

// APIRecipeToRecipe convert a list of ApiRecipe in a struct Recipe
func (h Recipe) APIRecipeToRecipe(apiRecipes []model.APIRecipe) (recipes []model.Recipe, err error) {
	var (
		wg sync.WaitGroup
	)

	recipes = make([]model.Recipe, len(apiRecipes))

	for index, apiRecipe := range apiRecipes {
		recipes[index].Title = apiRecipe.Title
		recipes[index].Link = apiRecipe.Href

		apiRecipe.Ingredients = strings.ReplaceAll(apiRecipe.Ingredients, ", ", ",")
		recipes[index].Ingredients = strings.Split(apiRecipe.Ingredients, ",")
		sort.Strings(recipes[index].Ingredients)

		wg.Add(1)
		go func(title string, wg *sync.WaitGroup, resultGif *string) {
			defer wg.Done()
			if *resultGif, err = h.GifSearcher.Search(title); err != nil {
				h.Logger.Error(err)
				return
			}
		}(recipes[index].Title, &wg, &recipes[index].Gif)
	}
	wg.Wait()

	return
}

// Recipes handles requests made to the endpoint /recipes
func (h Recipe) GetRecipes(c echo.Context) (err error) {
	var (
		recipes       []model.Recipe
		recipeResp    model.RecipeResponse
		recipeAPIResp model.APIRecipeResponse
	)

	ingredients := c.QueryParam("i")
	if recipeAPIResp, err = h.RecipeSearcher.Search(ingredients); err != nil {
		h.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	if len(recipeAPIResp.Results) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "recipe not found"})
	}
	if recipes, err = h.APIRecipeToRecipe(recipeAPIResp.Results); err != nil {
		h.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	recipeResp = model.RecipeResponse{
		Keywords: strings.Split(ingredients, ","),
		Recipes:  recipes,
	}

	return c.JSON(http.StatusOK, recipeResp)
}
