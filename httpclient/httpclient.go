package httpclient

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type HTTPClient struct{}

type Requester interface {
	MakeRequest(method string, url string) (content []byte, err error)
}

func (h HTTPClient) MakeRequest(method string, url string) (content []byte, err error) {
	client := http.Client{}
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return content, errors.New(resp.Status)
	}

	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	return
}
