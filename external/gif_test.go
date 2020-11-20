package external

import (
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/mocks"
	"github.com/paraizofelipe/gorecipes/settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGifSearch(t *testing.T) {

	var (
		JSONGif = `{
			"data": [
				{
					"url": "https://giphy.com/gifs/8itapp-hungry-hangry-8it-KbAbh6CbK4r5dyaZ4i"
				}
			]
		}`
		e               *echo.Echo = echo.New()
		mockedRequester            = &mocks.Requester{}
	)

	t.Run("should return a string with link to a gif", func(t *testing.T) {
		gif := Gif{
			Logger:     e.Logger,
			HTTPClient: mockedRequester,
		}

		testURL, _ := url.Parse(gifURL)

		params := url.Values{}
		params.Add("api_key", settings.GiphyToken)
		params.Add("limit", "1")
		params.Add("q", strings.TrimSpace(""))
		testURL.RawQuery = params.Encode()

		mockedRequester.Mock = mock.Mock{}
		mockedRequester.On("MakeRequest", "GET", testURL.String()).Return([]byte(JSONGif), nil)

		gifLink, err := gif.Search("")
		assert.NoError(t, err)
		assert.Equal(t, "https://giphy.com/gifs/8itapp-hungry-hangry-8it-KbAbh6CbK4r5dyaZ4i", gifLink)
	})
}
