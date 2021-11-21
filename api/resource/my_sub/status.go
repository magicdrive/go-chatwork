package my_sub

import (
	"net/http"

	json "github.com/goccy/go-json"

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

func NewStatus(parent string, credential string) StatusResource {
	data := StatusResource{
		ResourceName: parent + `/status`,
		Credential:   credential,
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