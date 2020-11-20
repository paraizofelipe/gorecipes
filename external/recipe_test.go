package external

import (
	"net/url"
	"testing"

	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRecipeSearch(t *testing.T) {

	var (
		JSONRecipe                 = `{"title":"Recipe Puppy","version":0.1,"href":"http:\/\/www.recipepuppy.com\/","results":[{"title":"Tomato & Orange Cottage Cheese Salad \r\n\t\t\n","href":"http:\/\/www.kraftfoods.com\/kf\/recipes\/tomato-orange-cottage-cheese-54326.aspx","ingredients":"tomato, orange, cottage cheese, red onions, basil, balsamic vinaigrette","thumbnail":"http:\/\/img.recipepuppy.com\/636990.jpg"},{"title":"A Citrus Salsa Recipe","href":"http:\/\/www.grouprecipes.com\/19982\/a-citrus-salsa.html","ingredients":"lemon, lemon juice, orange, orange juice, tomato, sugar","thumbnail":"http:\/\/img.recipepuppy.com\/169997.jpg"},{"title":"\nDip For Taco Chips Recipe\n\n","href":"http:\/\/cookeatshare.com\/recipes\/dip-for-taco-chips-998","ingredients":"cream cheese, tomato, cherry pepper, salsa, monterey jack cheese, orange","thumbnail":"http:\/\/img.recipepuppy.com\/855798.jpg"},{"title":"Kidney Bean and Orange Salad \r\n\t\t\n","href":"http:\/\/www.kraftfoods.com\/kf\/recipes\/kidney-bean-orange-salad-74387.aspx","ingredients":"red kidney beans, tomato, orange, green pepper, green onion, italian dressing, almonds","thumbnail":"http:\/\/img.recipepuppy.com\/638084.jpg"},{"title":"Fennel &amp; Tomato Salad With Orange Dressing","href":"http:\/\/www.recipezaar.com\/Fennel-Tomato-Salad-With-Orange-Dressing-218195","ingredients":"balsamic vinegar, brown sugar, fennel, orange, vegetable oil, soy sauce, tomato","thumbnail":"http:\/\/img.recipepuppy.com\/558239.jpg"},{"title":"Simple tomato, orange, grapefruit salad","href":"http:\/\/www.nibbledish.com\/people\/Dorara\/recipes\/simple-tomato-orange-grapefruit-salad","ingredients":"orange, grapefruit, tomato, basil, olive oil, black pepper, salt","thumbnail":"http:\/\/img.recipepuppy.com\/509643.jpg"},{"title":"Tomato & Orange Salad with Feta \r\n\t\t\r\n\t\r\n\t\t\r\n\t\r\n\t\t\r\n\t\r\n\t\r\n\r\n","href":"http:\/\/www.kraftfoods.com\/kf\/recipes\/tomato-orange-salad-feta-50786.aspx","ingredients":"tomato, orange, feta cheese, basil, olive oil, balsamic vinegar, salt, black pepper","thumbnail":"http:\/\/img.recipepuppy.com\/636530.jpg"},{"title":"Fruit Salsa","href":"http:\/\/allrecipes.com\/Recipe\/Fruit-Salsa\/Detail.aspx","ingredients":"avocado, cilantro, garlic salt, jalapeno, kiwi, orange, red onions, tomato","thumbnail":"http:\/\/img.recipepuppy.com\/25121.jpg"},{"title":"Halibut Steamed with Oranges, Tomatoes, and Olives","href":"http:\/\/www.epicurious.com\/recipes\/food\/views\/Halibut-Steamed-with-Oranges-Tomatoes-and-Olives-107632","ingredients":"feta cheese, garlic, olive oil, orange, tomato, red onions, halibut fillets, white wine","thumbnail":"http:\/\/img.recipepuppy.com\/107330.jpg"},{"title":"Cod Fish","href":"http:\/\/www.cooks.com\/rec\/view\/0,1917,158188-245200,00.html","ingredients":"dill weed, cod, orange, tomato","thumbnail":""}]}`
		e               *echo.Echo = echo.New()
		mockedRequester            = &mocks.Requester{}
	)

	t.Run("should return a object APIRecipeResponse with list of recipes", func(t *testing.T) {
		recipe := Recipe{
			Logger:     e.Logger,
			HTTPClient: mockedRequester,
		}

		testURL, _ := url.Parse(recipeURL)
		params := url.Values{}
		params.Add("i", "tomato,orange,onion")
		testURL.RawQuery = params.Encode()

		mockedRequester.Mock = mock.Mock{}
		mockedRequester.On("MakeRequest", "GET", testURL.String()).Return([]byte(JSONRecipe), nil)

		recipes, err := recipe.Search("tomato,orange,onion")
		assert.NoError(t, err)
		assert.Greater(t, len(recipes.Results), 0)
	})
}
