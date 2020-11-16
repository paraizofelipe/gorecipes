package model

// Recipe represents a structure of recipe inside application
type Recipe struct {
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	Gif         string   `json:"gif"`
}

// APIRecipe represents a recipe structure inside external API
type APIRecipe struct {
	Title       string `json:"title"`
	Href        string `json:"href"`
	Ingredients string `json:"ingredients"`
	Thumbnail   string `json:"thumbnail"`
}

// APIRecipeResponse represents the recipe's API response
type APIRecipeResponse struct {
	Title   string      `json:"title"`
	Version float64     `json:"version"`
	Href    string      `json:"href"`
	Results []APIRecipe `json:"results"`
}

// RecipeResponse represents the response of recipe application
type RecipeResponse struct {
	Keywords []string `json:"keywords"`
	Recipes  []Recipe `json:"recipes"`
}
