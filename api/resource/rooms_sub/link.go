package rooms_sub

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
)

type LinkResource struct {
	ResourceName string
	Credential   string
}

type LinkData struct {
	Public         bool   `json:"public"`
	URL            string `json:"url"`
	NeedAcceptance bool   `json:"need_acceptance"`
	Description    string `json:"description"`
}

type LinkDeleteData struct {
	Public bool `json:"public"`
}

type LinkParam struct {
	Code           string `json:"code"`
	Description    string `json:"description"`
	NeedAcceptance bool   `json:"need_acceptance"`
}

func NewLinkResource(parent string, credential string) LinkResource {
	data := LinkResource{
		ResourceName: parent + `/%d/link`,
		Credential:   credential,
	}
	return data
}

func (c LinkResource) Get(room_id int) (LinkData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      nil,
	}

	result := LinkData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), result)
	}

	return result, err
}

func (c LinkResource) Create(room_id int, params LinkParam) (LinkData, error) {
	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      p,
	}

	result := LinkData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), result)
	}

	return result, err

}

func (c LinkResource) Edit(room_id int, params LinkParam) (LinkData, error) {
	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      p,
	}

	result := LinkData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), result)
	}

	return result, err

}

func (c LinkResource) Delete(room_id int) (LinkDeleteData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodDelete,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      nil,
	}

	result := LinkDeleteData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), result)
	}

	return result, err

}
