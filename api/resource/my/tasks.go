package my

import (
	"net/http"

	json "github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

type MyTasksResource struct {
	ResourceName string
	Credential   string
}

type MyTaskData struct {
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

type MyTasksListParam struct {
	AssignedByAccountId *optional.NullableInt `json:"assigned_by_account_id"`
	Status              *optional.NullableInt `json:"status"`
}

var (
	MyTaskStatusDone = optional.Int(1)
	MyTaskStatusOpen = optional.Int(2)
)

func NewMyTasks(parent string, credential string) MyTasksResource {
	data := MyTasksResource{
		ResourceName: parent + `/tasks`,
		Credential:   credential,
	}
	return data

}

func (c MyTasksResource) List(params MyTasksListParam) ([]MyTaskData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      p,
	}

	result := make([]MyTaskData, 0, 32)

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}
