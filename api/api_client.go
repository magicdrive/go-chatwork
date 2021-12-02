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
	"github.com/magicdrive/go-chatwork/optional"
)

const ApiHost = "https://api.chatwork.com"
const ApiVersion = "v2"

var ApiEndpoint = fmt.Sprintf("%s/%s", ApiHost, ApiVersion)

type ApiSpec struct {
	Credential  string
	Method      string
	ResouceName string
	Params      map[string]*optional.NullableString
}

type ApiSpecMultipart struct {
	Credential  string
	Method      string
	ResouceName string
	Params      map[string]io.Reader
}

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

	endpoint := fmt.Sprintf("%s/%s%s", ApiHost, ApiVersion, data.ResouceName)

	if data.Method == http.MethodGet {

		req, _ := http.NewRequest(data.Method, endpoint, nil)

		if data.Params != nil {
			params := req.URL.Query()
			for key := range data.Params {
				if data.Params[key].IsPresent() {
					params.Add(key, data.Params[key].Valid().Get())
				}
			}

			req.URL.RawQuery = params.Encode()

		}
		return req

	} else {

		if data.Params != nil {
			params := url.Values{}
			for key := range data.Params {
				if data.Params[key].IsPresent() {
					params.Add(key, data.Params[key].Valid().Get())
				}
			}

			req, _ := http.NewRequest(data.Method, endpoint, bytes.NewBufferString(params.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			return req

		} else {

			req, _ := http.NewRequest(data.Method, endpoint, nil)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			return req
		}
	}
}

func JsonToMap(data []byte) (map[string]*optional.NullableString, error) {
	result := make(map[string]*optional.NullableString)
	unmarshaled := make(map[string]*optional.NullableString)

	if err := json.Unmarshal([]byte(data), &unmarshaled); err != nil {
		return nil, err
	}

	for key := range unmarshaled {
		if unmarshaled[key] != nil {
			result[key] = unmarshaled[key]
		}
	}
	return result, nil
}

func HttpRequestMultipart(data ApiSpecMultipart) (*http.Request, error) {

	endpoint := fmt.Sprintf("%s/%s", ApiEndpoint, data.ResouceName)
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
