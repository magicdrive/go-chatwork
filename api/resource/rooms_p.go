package resource

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
	rooms "github.com/magicdrive/go-chatwork/api/resource/rooms"
	"github.com/magicdrive/go-chatwork/optional"
)

type RoomsResource struct {
	ResourceName string
	Credential   string
}

type RoomData struct {
	RoomID         int    `json:"room_id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Role           string `json:"role"`
	Sticky         bool   `json:"sticky"`
	UnreadNum      int    `json:"unread_num"`
	MentionNum     int    `json:"mention_num"`
	MytaskNum      int    `json:"mytask_num"`
	MessageNum     int    `json:"message_num"`
	FileNum        int    `json:"file_num"`
	TaskNum        int    `json:"task_num"`
	IconPath       string `json:"icon_path"`
	LastUpdateTime int    `json:"last_update_time"`
}
type RoomCreateData struct {
	RoomID int `json:"room_id"`
}

type RoomUpdateData RoomsCreateParam

type RoomsCreateParam struct {
	Description        *optional.NullableString `json:"description"`
	IconPreset         *optional.NullableString `json:"icon_preset"`
	Link               *optional.NullableBool   `json:"link"`
	LinkCode           *optional.NullableString `json:"link_code"`
	LinkNeedAcceptance *optional.NullableBool   `json:"link_need_acceptance"`
	MembersAdminIds    []int                    `json:"members_admin_ids"`
	MembersMemberIds   []*optional.NullableInt  `json:"members_member_ids"`
	MembersReadonlyIds []*optional.NullableInt  `json:"members_readonly_ids"`
	Name               string                   `json:"name"`
}

type RoomsUpdateParam struct {
	Description *optional.NullableString `json:"description"`
	IconPreset  *optional.NullableString `json:"icon_preset"`
	Name        *optional.NullableString `json:"name"`
}

var (
	IconPresetGroup    = optional.String("group")
	IconPresetCheck    = optional.String("check")
	IconPresetDocument = optional.String("document")
	IconPresetMeeting  = optional.String("meeting")
	IconPresetEvent    = optional.String("event")
	IconPresetProject  = optional.String("project")
	IconPresetBusiness = optional.String("business")
	IconPresetStudy    = optional.String("study")
	IconPresetSecurity = optional.String("security")
	IconPresetStar     = optional.String("star")
	IconPresetIdea     = optional.String("idea")
	IconPresetHeart    = optional.String("heart")
	IconPresetMagcup   = optional.String("magcup")
	IconPresetBeer     = optional.String("beer")
	IconPresetMusic    = optional.String("music")
	IconPresetSports   = optional.String("sports")
	IconPresetTravel   = optional.String("travel")
)

const (
	RoomLeave  = "leave"
	RoomDelete = "delete"
)

func NewRoomsResource(credential string) RoomsResource {
	data := RoomsResource{
		ResourceName: `/rooms`,
		Credential:   credential,
	}
	return data
}

func (c RoomsResource) List() ([]RoomData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      nil,
	}

	result := make([]RoomData, 0, 32)

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c RoomsResource) Create(params RoomsCreateParam) (RoomCreateData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPost,
		ResouceName: c.ResourceName,
		Params:      p,
	}

	result := RoomCreateData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c RoomsResource) Get(room_id int) (RoomData, error) {

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, room_id),
		Params:      nil,
	}

	result := RoomData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

func (c RoomsResource) Update(room_id int, params RoomsUpdateParam) (RoomUpdateData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, room_id),
		Params:      p,
	}

	result := RoomUpdateData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

func (c RoomsResource) Delete(room_id int, action string) error {

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodDelete,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, room_id),
		Params: map[string]*optional.NullableString{
			"action_type": optional.String(action),
		},
	}

	_, err := api.Call(spec)

	return err
}

func (c RoomsResource) Files() rooms.FilesResource {
	return rooms.NewFilesResource(c.ResourceName, c.Credential)
}

func (c RoomsResource) Link() rooms.LinkResource {
	return rooms.NewLinkResource(c.ResourceName, c.Credential)
}

func (c RoomsResource) Members() rooms.MembersResource {
	return rooms.NewMembersResource(c.ResourceName, c.Credential)
}

func (c RoomsResource) Message() rooms.MessagesResource {
	return rooms.NewMessagesResource(c.ResourceName, c.Credential)
}

func (c RoomsResource) Tasks() rooms.TasksResource {
	return rooms.NewTasksResource(c.ResourceName, c.Credential)
}
