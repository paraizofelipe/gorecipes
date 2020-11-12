package handler

import (
	"log"
	"net/http"
)

// RecipesHandler ---
func (h Handler) RecipeHandler(w http.ResponseWriter, r *http.Request) {
	router := NewRouter(h.Logger)
	router.Add(`recipes\/?$`, http.MethodGet, h.getRecipes())

	router.ServeHTTP(w, r)
}

func (h Handler) getRecipes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pathParams string
		ctx := r.Context()
		w.Header().Set("Content-Type", "application/json")

		pathParams = ctx.Value("ingredients").(string)
		log.Printf(" --- PARAMS %s", pathParams)
		w.WriteHeader(http.StatusOK)
		r = r.WithContext(ctx)
	}
}
