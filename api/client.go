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

	"github.com/magicdrive/go-chatwork/api/param/optional"
)

// ApiDefaultHost chatwork api default host.
const ApiDefaultHost = "https://api.chatwork.com"

// ApiVersion chatwork api version string.
const ApiVersion = "v2"

// ApiSpec chatwork api spec difinition type.
type ApiSpec struct {
	Credential  string
	Method      string
	ResouceName string
	Params      map[string]*optional.NullableString
}

// ApiSpecMultipart chatwork api spec difinition type for multipart.
type ApiSpecMultipart struct {
	Credential  string
	Method      string
	ResouceName string
	Params      map[string]io.Reader
}

// ChatworkApiClient chatwork api client type.
type ChatworkApiClient struct {
	Credential         string
	HttpClient         *http.Client
	AltChatworkApiHost string
	ApiEndpoint        string
}

// NewChatworkApiClient new chatwork api client.
func NewChatworkApiClient(credential string, client *http.Client, alt_chatwork_api_host string) *ChatworkApiClient {
	var endpoint string
	if alt_chatwork_api_host == "" {
		endpoint = fmt.Sprintf("%s/%s", ApiDefaultHost, ApiVersion)
	} else {
		endpoint = fmt.Sprintf("%s/%s", alt_chatwork_api_host, ApiVersion)
	}
	return &ChatworkApiClient{
		Credential:         credential,
		HttpClient:         client,
		AltChatworkApiHost: alt_chatwork_api_host,
		ApiEndpoint:        endpoint,
	}
}

// CallMultipart call chatwork api with multipart.
func (c *ChatworkApiClient) CallMultipart(data ApiSpecMultipart) ([]byte, error) {

	req, err := c.HttpRequestMultipart(data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-ChatworkToken", data.Credential)

	var client *http.Client

	if c.HttpClient == nil {
		client = new(http.Client)
	} else {
		client = c.HttpClient
	}

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

// Call call chatwork api.
func (c *ChatworkApiClient) Call(data ApiSpec) ([]byte, error) {

	req := c.HttpRequest(data)
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

// HttpRequest build chatwork api httprequest.
func (c *ChatworkApiClient) HttpRequest(data ApiSpec) *http.Request {

	endpoint := fmt.Sprintf("%s%s", c.ApiEndpoint, data.ResouceName)

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

// HttpRequest build chatwork api httprequest with multipart.
func (c *ChatworkApiClient) HttpRequestMultipart(data ApiSpecMultipart) (*http.Request, error) {

	endpoint := fmt.Sprintf("%s%s", c.ApiEndpoint, data.ResouceName)
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

// JsonToMap convert Json string bytes to map[stirng]*optional.NullableString.
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
