package rooms

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

type RoomLinkResource struct {
	ResourceName string
	Credential   string
}

type RoomLinkData struct {
	Public         bool   `json:"public"`
	URL            string `json:"url"`
	NeedAcceptance bool   `json:"need_acceptance"`
	Description    string `json:"description"`
}

type RoomLinkDeleteData struct {
	Public bool `json:"public"`
}

type RoomLinkParam struct {
	Code           *optional.NullableString `json:"code"`
	Description    *optional.NullableString `json:"description"`
	NeedAcceptance *optional.NullableBool   `json:"need_acceptance"`
}

func NewRoomLinkResource(parent string, credential string) RoomLinkResource {
	data := RoomLinkResource{
		ResourceName: parent + `/%d/link`,
		Credential:   credential,
	}
	return data
}

func (c RoomLinkResource) Get(room_id int) (RoomLinkData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      nil,
	}

	result := RoomLinkData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c RoomLinkResource) Create(room_id int, params RoomLinkParam) (RoomLinkData, error) {
	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      p,
	}

	result := RoomLinkData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

func (c RoomLinkResource) Edit(room_id int, params RoomLinkParam) (RoomLinkData, error) {
	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      p,
	}

	result := RoomLinkData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

func (c RoomLinkResource) Delete(room_id int) (RoomLinkDeleteData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodDelete,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      nil,
	}

	result := RoomLinkDeleteData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}
