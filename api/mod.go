package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func Call(spec_data ToApiSpecIntf) (string, error) {
	data := spec_data.ToApiSpec()

	req, _ := http.NewRequest(
		data.Method, fmt.Sprintf("%s/%s/%s", ApiHost, ApiVersion, data.ResouceName), nil)
	req.Header.Set("X-ChatworkToken", data.Credential)

	//TODO:  param & methods
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	return string(byteArray), nil
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
