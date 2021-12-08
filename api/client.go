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

// APIDefaultHost chatwork api default host.
const APIDefaultHost = "https://api.chatwork.com"

// APIVersion chatwork api version string.
const APIVersion = "v2"

// APISpec chatwork api spec difinition type.
type APISpec struct {
	Credential  string
	Method      string
	ResouceName string
	Params      map[string]*optional.NullableString
}

// APISpecMultipart chatwork api spec difinition type for multipart.
type APISpecMultipart struct {
	Credential  string
	Method      string
	ResouceName string
	Params      map[string]io.Reader
}

// ChatworkAPIClient chatwork api client type.
type ChatworkAPIClient struct {
	Credential         string
	HTTPClient         *http.Client
	AltChatworkAPIHost string
	APIEndpoint        string
}

// NewChatworkAPIClient new chatwork api client.
func NewChatworkAPIClient(credential string, client *http.Client, altChatworkAPIHost string) *ChatworkAPIClient {
	var endpoint string
	if altChatworkAPIHost == "" {
		endpoint = fmt.Sprintf("%s/%s", APIDefaultHost, APIVersion)
	} else {
		endpoint = fmt.Sprintf("%s/%s", altChatworkAPIHost, APIVersion)
	}
	return &ChatworkAPIClient{
		Credential:         credential,
		HTTPClient:         client,
		AltChatworkAPIHost: altChatworkAPIHost,
		APIEndpoint:        endpoint,
	}
}

// CallMultipart call chatwork api with multipart.
func (c *ChatworkAPIClient) CallMultipart(data APISpecMultipart) ([]byte, error) {

	req, err := c.HTTPRequestMultipart(data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-ChatworkToken", data.Credential)

	var client *http.Client

	if c.HTTPClient == nil {
		client = new(http.Client)
	} else {
		client = c.HTTPClient
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
func (c *ChatworkAPIClient) Call(data APISpec) ([]byte, error) {

	req := c.HTTPRequest(data)
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

// HTTPRequest build chatwork api httprequest.
func (c *ChatworkAPIClient) HTTPRequest(data APISpec) *http.Request {

	endpoint := fmt.Sprintf("%s%s", c.APIEndpoint, data.ResouceName)

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

// HTTPRequestMultipart build chatwork api httprequest with multipart.
func (c *ChatworkAPIClient) HTTPRequestMultipart(data APISpecMultipart) (*http.Request, error) {

	endpoint := fmt.Sprintf("%s%s", c.APIEndpoint, data.ResouceName)
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

// JSONToMap convert Json string bytes to map[stirng]*optional.NullableString.
func JSONToMap(data []byte) (map[string]*optional.NullableString, error) {
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
