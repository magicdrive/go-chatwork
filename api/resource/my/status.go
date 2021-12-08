package my

import (
	"net/http"

	json "github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api"
)

// MyStatusResource chatwork api my/status resource.
type MyStatusResource struct {
	ResourceName string
	Client       *api.ChatworkAPIClient
}

// MyStatusData chatwork api resp my/status data.
type MyStatusData struct {
	UnreadRoomNum  int `json:"unread_room_num"`
	MentionRoomNum int `json:"mention_room_num"`
	MytaskRoomNum  int `json:"mytask_room_num"`
	UnreadNum      int `json:"unread_num"`
	MentionNum     int `json:"mention_num"`
	MytaskNum      int `json:"mytask_num"`
}

// NewMyStatus new chatwork api resp my/status resource.
func NewMyStatus(parent string, client *api.ChatworkAPIClient) MyStatusResource {
	data := MyStatusResource{
		ResourceName: parent + `/status`,
		Client:       client,
	}
	return data

}

// Get chatwork api get my/status.
func (c MyStatusResource) Get() (MyStatusData, error) {
	spec := api.APISpec{
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
