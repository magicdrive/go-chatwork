package rooms

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

// RoomLinkResource chatwork api rooms/link resouce
type RoomLinkResource struct {
	ResourceName string
	Client   *api.ChatworkApiClient
}

// RoomLinkData chatwork api resp rooms/link data
type RoomLinkData struct {
	Public         bool   `json:"public"`
	URL            string `json:"url"`
	NeedAcceptance bool   `json:"need_acceptance"`
	Description    string `json:"description"`
}

// RoomLinkDeleteData chatwork api resp rooms/link delete data
type RoomLinkDeleteData struct {
	Public bool `json:"public"`
}

// RoomLinkParam chatwork api rooms/link param.
type RoomLinkParam struct {
	Code           *optional.NullableString `json:"code"`
	Description    *optional.NullableString `json:"description"`
	NeedAcceptance *optional.NullableBool   `json:"need_acceptance"`
}

// NewRoomLinkResource new chatwork api rooms/link resouce.
func NewRoomLinkResource(parent string, client *api.ChatworkApiClient) RoomLinkResource {
	data := RoomLinkResource{
		ResourceName: parent + `/%d/link`,
		Client:   client,
	}
	return data
}

// Get chatwork api get rooms/link .
func (c RoomLinkResource) Get(room_id int) (RoomLinkData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      nil,
	}

	result := RoomLinkData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

// Create chatwork api create rooms/link.
func (c RoomLinkResource) Create(room_id int, params RoomLinkParam) (RoomLinkData, error) {
	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      p,
	}

	result := RoomLinkData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

// Edit chatwork api edit rooms/link.
func (c RoomLinkResource) Edit(room_id int, params RoomLinkParam) (RoomLinkData, error) {
	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      p,
	}

	result := RoomLinkData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

// Delete chatwork api delete rooms/link.
func (c RoomLinkResource) Delete(room_id int) (RoomLinkDeleteData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodDelete,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      nil,
	}

	result := RoomLinkDeleteData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}
