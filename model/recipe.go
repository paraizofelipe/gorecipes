package model

type Recipe struct {
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	Gif         string   `json:"gif"`
}

// APIRecipe ---
type APIRecipe struct {
	Title       string `json:"title"`
	Href        string `json:"href"`
	Ingredients string `json:"ingredients"`
	Thumbnail   string `json:"thumbnail"`
}

// APIRecipeResponse ---
type APIRecipeResponse struct {
	Title   string      `json:"title"`
	Version float64     `json:"version"`
	Href    string      `json:"href"`
	Results []APIRecipe `json:"results"`
}
