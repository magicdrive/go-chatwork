package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

const ApiHost = "https://api.chatwork.com"
const ApiVersion = "v2"

type ApiSpec struct {
	Credential  string
	Method      string
	ResouceName string
	Params      map[string]string
}

type ToApiSpecIntf interface {
	ToApiSpec() ApiSpec
}

func Call(spec ToApiSpecIntf) ([]byte, error) {
	data := spec.ToApiSpec()

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

func StructToMap(data interface{}) map[string]string {

	result := make(map[string]string)
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		key := elem.Type().Field(i).Name
		value := elem.Field(i).String()
		result[key] = value
	}

	return result
}
