package rooms_sub

import (
	"fmt"
	"net/http"
	"strconv"

	json "github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/optional"
)

type MessagesResource struct {
	ResourceName string
	Credential   string
}

type MessageData struct {
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

type MessagePostParam struct {
	Body       string                 `json:"body"`
	SelfUnread *optional.NullableBool `json:"self_unread"`
}

type MessagePostData struct {
	MessageId int `json:"message_id"`
}

type MessageReadStatusData struct {
	UnreadNum int `json:"unread_num"`
	MetionNum int `json:"mention_num"`
}

func NewMessagesResource(parent string, credential string) MessagesResource {
	data := MessagesResource{
		ResourceName: parent + `/%d/messages`,
		Credential:   credential,
	}
	return data

}

func (c MessagesResource) List(room_id string, force optional.NullableBool) ([]MessageData, error) {

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params: map[string]*optional.NullableString{
			"force": force.ToNullableString(),
		},
	}

	result := make([]MessageData, 0, 32)

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c MessagesResource) Post(room_id int, params MessagePostParam) (MessagePostData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      p,
	}

	result := MessagePostData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c MessagesResource) Read(room_id int, message_id *optional.NullableInt) (MessageReadStatusData, error) {

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName+"/read", room_id),
		Params: map[string]*optional.NullableString{
			"message_id": message_id.ToNullableString(),
		},
	}

	result := MessageReadStatusData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}
func (c MessagesResource) Unread(room_id int, message_id *optional.NullableInt) (MessageReadStatusData, error) {
	var _message_id *optional.NullableString
	if message_id.IsPresent() {
		_message_id = optional.String(strconv.FormatInt(message_id.Valid().Get(), 10))
	} else {
		_message_id = optional.NilString()
	}
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName+"/unread", room_id),
		Params: map[string]*optional.NullableString{
			"message_id": _message_id,
		},
	}

	result := MessageReadStatusData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c MessagesResource) Get(room_id int, message_id int) (MessageData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%d", room_id, message_id),
		Params:      nil,
	}

	result := MessageData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

func (c MessagesResource) Edit(room_id int, message_id int, body string) (MessagePostData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%d", room_id, message_id),
		Params:      map[string]*optional.NullableString{"body": optional.String(body)},
	}

	result := MessagePostData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

func (c MessagesResource) Delete(room_id int, message_id int, body string) (MessagePostData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodDelete,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%d", room_id, message_id),
		Params:      nil,
	}

	result := MessagePostData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}
