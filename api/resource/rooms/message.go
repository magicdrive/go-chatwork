package rooms

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

// RoomMessagesResource chatwork api rooms/message resouce.
type RoomMessagesResource struct {
	ResourceName string
	Client       *api.ChatworkAPIClient
}

// RoomMessageData chatwork api rooms/message resp data.
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

// RoomMessagePostParam chatwork api rooms/message post param.
type RoomMessagePostParam struct {
	Body       string                 `json:"body"`
	SelfUnread *optional.NullableBool `json:"self_unread"`
}

// RoomMessagePostData chatwork api resp rooms/message post data.
type RoomMessagePostData struct {
	MessageID string `json:"message_id"`
}

// RoomMessageReadStatusData chatwork api resp rooms/message read status data.
type RoomMessageReadStatusData struct {
	UnreadNum int `json:"unread_num"`
	MetionNum int `json:"mention_num"`
}

// NewRoomMessagesResource new chatwork api rooms/messages resouce.
func NewRoomMessagesResource(parent string, client *api.ChatworkAPIClient) RoomMessagesResource {
	data := RoomMessagesResource{
		ResourceName: parent + `/%d/messages`,
		Client:   client,
	}
	return data

}

// List chatwork api get rooms/messages list.
func (c RoomMessagesResource) List(roomID int, force *optional.NullableBool) ([]RoomMessageData, error) {

	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, roomID),
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

// Post chatwork api post rooms/message.
func (c RoomMessagesResource) Post(roomID int, params RoomMessagePostParam) (RoomMessagePostData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JSONToMap(b)

	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, roomID),
		Params:      p,
	}

	result := RoomMessagePostData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

// Read chatwork api mark as read to rooms/message.
func (c RoomMessagesResource) Read(roomID int, messageID *optional.NullableString) (RoomMessageReadStatusData, error) {

	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName+"/read", roomID),
		Params: map[string]*optional.NullableString{
			"message_id": messageID,
		},
	}

	result := RoomMessageReadStatusData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

// Unread chatwork api mark as unread to rooms/message.
func (c RoomMessagesResource) Unread(roomID int, messageID *optional.NullableString) (RoomMessageReadStatusData, error) {
	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName+"/unread", roomID),
		Params: map[string]*optional.NullableString{
			"message_id": messageID,
		},
	}

	result := RoomMessageReadStatusData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

// Get chatwork api get rooms/message.
func (c RoomMessagesResource) Get(roomID int, messageID string) (RoomMessageData, error) {
	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%s", roomID, messageID),
		Params:      nil,
	}

	result := RoomMessageData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

// Edit chatwork api edit rooms/message.
func (c RoomMessagesResource) Edit(roomID int, messageID string, body string) (RoomMessagePostData, error) {
	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%s", roomID, messageID),
		Params:      map[string]*optional.NullableString{"body": optional.String(body)},
	}

	result := RoomMessagePostData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

// Delete chatwork api delete rooms/message.
func (c RoomMessagesResource) Delete(roomID int, messageID string) (RoomMessagePostData, error) {
	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodDelete,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%s", roomID, messageID),
		Params:      nil,
	}

	result := RoomMessagePostData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}
