package resource

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
	rooms "github.com/magicdrive/go-chatwork/api/resource/rooms_sub"
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

type RoomsCreateParam struct {
	Description        string `json:"description"`
	IconPreset         string `json:"icon_preset"`
	Link               bool   `json:"link"`
	LinkCode           string `json:"link_code"`
	LinkNeedAcceptance bool   `json:"link_need_acceptance"`
	MembersAdminIds    []int  `json:"members_admin_ids"`
	MembersMemberIds   []int  `json:"members_member_ids"`
	MembersReadonlyIds []int  `json:"members_readonly_ids"`
	Name               string `json:"name"`
}

type RoomsUpdateParam struct {
	Description string `json:"description"`
	IconPreset  string `json:"icon_preset"`
	Name        string `json:"name"`
}

const (
	IconPresetGroup    = "group"
	IconPresetCheck    = "check"
	IconPresetDocument = "document"
	IconPresetMeeting  = "meeting"
	IconPresetEvent    = "event"
	IconPresetProject  = "project"
	IconPresetBusiness = "business"
	IconPresetStudy    = "study"
	IconPresetSecurity = "security"
	IconPresetStar     = "star"
	IconPresetIdea     = "idea"
	IconPresetHeart    = "heart"
	IconPresetMagcup   = "magcup"
	IconPresetBeer     = "beer"
	IconPresetMusic    = "music"
	IconPresetSports   = "sports"
	IconPresetTravel   = "travel"
)

const (
	RoomLeave = iota
	RoomDelete
)

var (
	_roomAction = []string{"leave", "delete"}
)

func NewRoomsResource(credential string) MyResource {
	data := MyResource{
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

func (c RoomsResource) Create(params RoomsCreateParam) error {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPost,
		ResouceName: c.ResourceName,
		Params:      p,
	}

	result := make([]RoomData, 0, 32)

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return err
}

func (c RoomsResource) Get(room_id int) (RoomData, error) {

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPost,
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

func (c RoomsResource) Update(room_id int, params RoomsUpdateParam) error {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, room_id),
		Params:      p,
	}

	_, err := api.Call(spec)

	return err

}

func (c RoomsResource) Delete(room_id int, action int) error {

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, room_id),
		Params:      map[string]string{"action": _roomAction[action]},
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
