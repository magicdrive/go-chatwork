package api

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

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

type ApiSpecMultipart struct {
	Credential  string
	Method      string
	ResouceName string
	Params      map[string]io.Reader
}

const (
	ChatworkBoolFalse = iota
	ChatworkBoolTrue
)

func CallMultipart(data ApiSpecMultipart) ([]byte, error) {

	req, err := HttpRequestMultipart(data)
	if err != nil {
		return nil, err
	}
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

func HttpRequestMultipart(data ApiSpecMultipart) (*http.Request, error) {

	endpoint := fmt.Sprintf("%s/%s/%s", ApiHost, ApiVersion, data.ResouceName)
	values := data.Params

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	defer writer.Close()

	for key, r := range values {
		var fieldw io.Writer
		var err error
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}

		if x, ok := r.(*os.File); ok {
			if fieldw, err = writer.CreateFormFile(key, x.Name()); err != nil {
				return nil, err
			}
		} else {
			if fieldw, err = writer.CreateFormField(key); err != nil {
				return nil, err
			}
		}
		if _, err := io.Copy(fieldw, r); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(data.Method, endpoint, &b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, err
}
