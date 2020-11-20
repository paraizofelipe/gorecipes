package external

import (
	"github.com/paraizofelipe/gorecipes/model"
)

type RecipeSearcher interface {
	Search(ingredients string) (model.APIRecipeResponse, error)
}

type GifSearcher interface {
	Search(title string) (gif string, err error)
}

type External struct {
	RecipeSearcher RecipeSearcher
	GifSearcher    GifSearcher
}
