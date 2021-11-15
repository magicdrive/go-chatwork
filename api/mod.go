package api

import "reflect"

const ApiHost = "https://api.chatwork.com"
const ApiVersion = "v2"

type ApiSpec struct {
	Credential  string
	Method      string
	ResouceName string
	Params      map[string]string
}

func call(spec ApiSpec) (string, error) {
	return "", nil

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
