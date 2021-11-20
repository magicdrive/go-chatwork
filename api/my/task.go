package my

import (
	"encoding/json"
	"net/http"

	"github.com/magicdrive/go-chatwork/api"
)

type TasksResource struct {
	ResourceName string
	Credential   string
}

type TaskData struct {
	TaskID int `json:"task_id"`
	Room   struct {
		RoomID   int    `json:"room_id"`
		Name     string `json:"name"`
		IconPath string `json:"icon_path"`
	} `json:"room"`
	AssignedByAccount struct {
		AccountID      int    `json:"account_id"`
		Name           string `json:"name"`
		AvatarImageURL string `json:"avatar_image_url"`
	} `json:"assigned_by_account"`
	MessageID string `json:"message_id"`
	Body      string `json:"body"`
	LimitTime int    `json:"limit_time"`
	Status    string `json:"status"`
	LimitType string `json:"limit_type"`
}

type ParamsGet struct {
	AssignedByAccountId int `json:"assigned_by_account_id"`
	Status              int `json:"status"`
}

const (
	Done = iota
	Open
)

func NewTasks(parent api.MyResource) TasksResource {
	data := TasksResource{
		ResourceName: parent.ResourceName + `/tasks`,
		Credential:   parent.Credential,
	}
	return data

}

func (c TasksResource) Get(params ParamsGet) ([]TaskData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      p,
	}

	result := make([]TaskData, 32)

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}
