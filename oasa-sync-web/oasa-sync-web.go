package oasa_sync_web

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
	"time"
)

const OASA_SERVER = "http://telematics.oasa.gr/api/"

func getRequest(url string, headers map[string]string) (*http.Response, error) {
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

func MakeRequest(action string) (string, error) {
	response, err := getRequest(OASA_SERVER+"?act="+action, map[string]string{
		"Accept-Encoding": "gzip, deflate"})
	if err != nil {
		return "", err
	}

	reader, err := gzip.NewReader(response.Body)

	if err != nil {
		fmt.Printf(err.Error())
		return "", err
	} else {
		defer reader.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		responseStr := buf.String()
		return responseStr, nil
	}
}
