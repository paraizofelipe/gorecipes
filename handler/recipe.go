package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/paraizofelipe/gorecipes/api"
	"github.com/paraizofelipe/gorecipes/model"
)

type RecipeResponse struct {
	Keywords []string       `json:"keywords"`
	Recipes  []model.Recipe `json:"recipes"`
}

// RecipeHandler ---
func (h Handler) RecipeHandler(w http.ResponseWriter, r *http.Request) {
	router := NewRouter(h.Logger)
	router.Add(`recipes\/?$`, http.MethodGet, h.getRecipes())

	router.ServeHTTP(w, r)
}

// APIRecipeToRecipe ---
func (h Handler) APIRecipeToRecipe(apiRecipes []model.APIRecipe) (recipes []model.Recipe, err error) {
	var (
		recipe model.Recipe
		wg     sync.WaitGroup
	)

	for _, apiRecipe := range apiRecipes {
		recipe.Title = apiRecipe.Title
		recipe.Link = apiRecipe.Href

		recipe.Ingredients = strings.Split(apiRecipe.Ingredients, ",")
		sort.Strings(recipe.Ingredients)

		wg.Add(1)
		go api.AsyncSearchGif(recipe.Title, &wg, &recipe.Gif)

		recipes = append(recipes, recipe)
	}
	wg.Wait()

	return
}

func (h Handler) getRecipes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err         error
			recipes     []model.Recipe
			response    RecipeResponse
			respRecipe  model.APIRecipeResponse
			ingredients string
		)

		ctx := r.Context()
		w.Header().Set("Content-Type", "application/json")

		ingredients = r.URL.Query().Get("i")
		if respRecipe, err = api.SearchRecipes(ingredients); err != nil {
			log.Println(err)
			http.Error(w, "error fetching recipe", http.StatusInternalServerError)
			return
		}
		if len(respRecipe.Results) == 0 {
			log.Println(err)
			http.Error(w, "recipes not found", http.StatusNotFound)
			return
		}
		if recipes, err = h.APIRecipeToRecipe(respRecipe.Results); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Recipes = recipes
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		r = r.WithContext(ctx)
	}
}
