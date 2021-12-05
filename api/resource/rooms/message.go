package rooms

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

type RoomMessagesResource struct {
	ResourceName string
	Client       *api.ChatworkApiClient
}

type RoomMessageData struct {
	MessageID string `json:"message_id"`
	Account   struct {
		AccountID      int    `json:"account_id"`
		Name           string `json:"name"`
		AvatarImageURL string `json:"avatar_image_url"`
	} `json:"account"`
	Body       string `json:"body"`
	SendTime   int    `json:"send_time"`
	UpdateTime int    `json:"update_time"`
}

type RoomMessagePostParam struct {
	Body       string                 `json:"body"`
	SelfUnread *optional.NullableBool `json:"self_unread"`
}

type RoomMessagePostData struct {
	MessageId string `json:"message_id"`
}

type RoomMessageReadStatusData struct {
	UnreadNum int `json:"unread_num"`
	MetionNum int `json:"mention_num"`
}

func NewRoomMessagesResource(parent string, client *api.ChatworkApiClient) RoomMessagesResource {
	data := RoomMessagesResource{
		ResourceName: parent + `/%d/messages`,
		Client:   client,
	}
	return data

}

func (c RoomMessagesResource) List(room_id int, force *optional.NullableBool) ([]RoomMessageData, error) {

	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params: map[string]*optional.NullableString{
			"force": force.ToNullableString(),
		},
	}

	result := make([]RoomMessageData, 0, 32)

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c RoomMessagesResource) Post(room_id int, params RoomMessagePostParam) (RoomMessagePostData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      p,
	}

	result := RoomMessagePostData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c RoomMessagesResource) Read(room_id int, message_id *optional.NullableString) (RoomMessageReadStatusData, error) {

	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName+"/read", room_id),
		Params: map[string]*optional.NullableString{
			"message_id": message_id,
		},
	}

	result := RoomMessageReadStatusData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}
func (c RoomMessagesResource) Unread(room_id int, message_id *optional.NullableString) (RoomMessageReadStatusData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName+"/unread", room_id),
		Params: map[string]*optional.NullableString{
			"message_id": message_id,
		},
	}

	result := RoomMessageReadStatusData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c RoomMessagesResource) Get(room_id int, message_id string) (RoomMessageData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%s", room_id, message_id),
		Params:      nil,
	}

	result := RoomMessageData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

func (c RoomMessagesResource) Edit(room_id int, message_id string, body string) (RoomMessagePostData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%s", room_id, message_id),
		Params:      map[string]*optional.NullableString{"body": optional.String(body)},
	}

	result := RoomMessagePostData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c RoomMessagesResource) Delete(room_id int, message_id string) (RoomMessagePostData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodDelete,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%s", room_id, message_id),
		Params:      nil,
	}

	result := RoomMessagePostData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}
