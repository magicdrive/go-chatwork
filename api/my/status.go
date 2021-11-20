package my

import (
	"encoding/json"
	"net/http"

	"github.com/magicdrive/go-chatwork/api"
)

type StatusResource struct {
	ResourceName string
	Credential   string
}

type StatusData struct {
	UnreadRoomNum  int `json:"unread_room_num"`
	MentionRoomNum int `json:"mention_room_num"`
	MytaskRoomNum  int `json:"mytask_room_num"`
	UnreadNum      int `json:"unread_num"`
	MentionNum     int `json:"mention_num"`
	MytaskNum      int `json:"mytask_num"`
}

func NewStatus(parent api.MyResource) StatusResource {
	data := StatusResource{
		ResourceName: parent.ResourceName + `/status`,
		Credential:   parent.Credential,
	}
	return data

}

func (c StatusResource) Get() (StatusData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      nil,
	}

	result := StatusData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}
