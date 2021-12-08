package rooms

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

// RoomTasksResource chatwork api rooms/tasks resouce.
type RoomTasksResource struct {
	ResourceName string
	Client       *api.ChatworkAPIClient
}

// RoomTaskData chatwork api resp rooms/task data.
type RoomTaskData struct {
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

// RoomTasksListParam chatwork api rooms/task request param.
type RoomTasksListParam struct {
	AccountID           *optional.NullableInt    `json:"account_id"`
	AssignedByAccountID *optional.NullableInt    `json:"assigned_by_account_id"`
	Status              *optional.NullableString `json:"status"`
}

// RoomTaskPostParam chatwork api post rooms/task request param.
type RoomTaskPostParam struct {
	Body      string                   `json:"body"`
	Limit     *optional.NullableInt64  `json:"limit"`
	LimitType *optional.NullableString `json:"limit_type"`
	ToIDs     *param.IntArrayParam     `json:"to_ids"`
}

// RoomTaskPostData chatwork api resp rooms/task post data.
type RoomTaskPostData struct {
	TaskID []int `json:"task_ids"`
}

var (
	// RoomTaskLimitTypeNone binding room/task limit type param "none".
	RoomTaskLimitTypeNone = optional.String("none")
	// RoomTaskLimitTypeDate binding room/task limit type param "date".
	RoomTaskLimitTypeDate = optional.String("date")
	// RoomTaskLimitTypeTime binding room/task limit type param "time".
	RoomTaskLimitTypeTime = optional.String("time")
)
var (
	// RoomTaskBodyStatusOpen binding room/task body status param "open".
	RoomTaskBodyStatusOpen = optional.String("open")
	// RoomTaskBodyStatusDone binding room/task body status param "done".
	RoomTaskBodyStatusDone = optional.String("done")
)

// NewRoomTasksResource new chatwork api rooms/task resouce.
func NewRoomTasksResource(parent string, client *api.ChatworkAPIClient) RoomTasksResource {
	data := RoomTasksResource{
		ResourceName: parent + `/%d/tasks`,
		Client:       client,
	}
	return data

}

// List chatwork api get rooms/tasks list.
func (c RoomTasksResource) List(roomID int, params RoomTasksListParam) ([]RoomTaskData, error) {
	b, _ := json.Marshal(params)
	p, _ := api.JSONToMap(b)

	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, roomID),
		Params:      p,
	}

	result := make([]RoomTaskData, 0, 32)

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

// Create chatwork api create rooms/task.
func (c RoomTasksResource) Create(roomID int, params RoomTaskPostParam) (RoomTaskPostData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JSONToMap(b)

	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, roomID),
		Params:      p,
	}

	result := RoomTaskPostData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

// Get chatwork api get rooms/task.
func (c RoomTasksResource) Get(roomID int, taskID int) (RoomTaskData, error) {
	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%d", roomID, taskID),
		Params:      nil,
	}

	result := RoomTaskData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

// Update chatwork api update rooms/task.
func (c RoomTasksResource) Update(roomID int, taskID int, body *optional.NullableString) (RoomTaskPostData, error) {
	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%d/status", roomID, taskID),
		Params:      map[string]*optional.NullableString{"body": body},
	}

	result := RoomTaskPostData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}
