package rooms

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/optional"
)

type TasksResource struct {
	ResourceName string
	Credential   string
}

type TaskData struct {
	TaskID  int `json:"task_id"`
	Account struct {
		AccountID      int    `json:"account_id"`
		Name           string `json:"name"`
		AvatarImageURL string `json:"avatar_image_url"`
	} `json:"account"`
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

type TasksListParam struct {
	AccountID           *optional.NullableInt    `json:"account_id"`
	AssignedByAccountId *optional.NullableInt    `json:"assigned_by_account_id"`
	Status              *optional.NullableString `json:"status"`
}

type TaskPostParam struct {
	Body      string                   `json:"body"`
	Limit     *optional.NullableInt    `json:"limit"`
	LimitType *optional.NullableString `json:"limit_type"`
	ToIds     []int                    `json:"to_ids"`
}

type TaskPostData struct {
	TaskId []int `json:"task_ids"`
}

var (
	TaskLimitTypeNone = optional.String("none")
	TaskLimitTypeDate = optional.String("date")
	TaskLimitTypeTime = optional.String("time")
)
var (
	TaskBodyStatusOpen = optional.String("open")
	TaskBodyStatusDone = optional.String("done")
)

func NewTasksResource(parent string, credential string) TasksResource {
	data := TasksResource{
		ResourceName: parent + `/%d/tasks`,
		Credential:   credential,
	}
	return data

}

func (c TasksResource) List(room_id int, params TasksListParam) ([]TaskData, error) {
	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      p,
	}

	result := make([]TaskData, 0, 32)

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c TasksResource) Create(room_id int, params TaskPostParam) (TaskPostData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      p,
	}

	result := TaskPostData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c TasksResource) Get(room_id int, task_id int) (TaskData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%d", room_id, task_id),
		Params:      nil,
	}

	result := TaskData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

func (c TasksResource) Update(room_id int, task_id int, body *optional.NullableString) (TaskPostData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%d/status", room_id, task_id),
		Params:      map[string]*optional.NullableString{"body": body},
	}

	result := TaskPostData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}
