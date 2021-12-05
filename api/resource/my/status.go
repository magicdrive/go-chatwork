package my

import (
	"net/http"

	json "github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api"
)

type MyStatusResource struct {
	ResourceName string
	Client       *api.ChatworkApiClient
}

type MyStatusData struct {
	UnreadRoomNum  int `json:"unread_room_num"`
	MentionRoomNum int `json:"mention_room_num"`
	MytaskRoomNum  int `json:"mytask_room_num"`
	UnreadNum      int `json:"unread_num"`
	MentionNum     int `json:"mention_num"`
	MytaskNum      int `json:"mytask_num"`
}

func NewMyStatus(parent string, client *api.ChatworkApiClient) MyStatusResource {
	data := MyStatusResource{
		ResourceName: parent + `/status`,
		Client:       client,
	}
	return data

}

func (c MyStatusResource) Get() (MyStatusData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      nil,
	}

	result := MyStatusData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}
