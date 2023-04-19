package oasa_sync_web

import (
	"fmt"
	"net/http"
	"time"
)

func GetRequest(url string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("got error %s", err.Error())
		return nil, err
	}

	if headers != nil && len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err.Error())
		return nil, err
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", response.StatusCode)

	return response, nil
}
