package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	json "github.com/goccy/go-json"
)

const ApiHost = "https://api.chatwork.com"
const ApiVersion = "v2"

type ApiSpec struct {
	Credential  string
	Method      string
	ResouceName string
	Params      map[string]string
}

const (
	ChatworkBoolFalse = iota
	ChatworkBoolTrue
)

func Call(data ApiSpec) ([]byte, error) {

	req := HttpRequest(data)
	req.Header.Set("X-ChatworkToken", data.Credential)

	client := new(http.Client)

	if resp, err := client.Do(req); err != nil {
		return []byte{}, err
	} else {

		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			return []byte{}, err
		} else {
			return body, nil
		}
	}
}

func HttpRequest(data ApiSpec) *http.Request {

	endpoint := fmt.Sprintf("%s/%s/%s", ApiHost, ApiVersion, data.ResouceName)

	if data.Method == http.MethodGet {

		req, _ := http.NewRequest(data.Method, endpoint, nil)

		params := req.URL.Query()
		for key := range data.Params {
			params.Add(key, data.Params[key])
		}

		req.URL.RawQuery = params.Encode()
		return req

	} else {

		params := url.Values{}
		for key := range data.Params {
			params.Add(key, data.Params[key])
		}

		req, _ := http.NewRequest(data.Method, endpoint, bytes.NewBufferString(params.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		return req
	}
}

func JsonToMap(data []byte) (map[string]string, error) {
	result := make(map[string]string)
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
