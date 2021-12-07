package resource

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/magicdrive/go-chatwork/api/param/optional"
	rooms "github.com/magicdrive/go-chatwork/api/resource/rooms"
)

// RoomsResource chatwork api rooms resource
type RoomsResource struct {
	ResourceName string
	Client       *api.ChatworkApiClient
}

// RoomsData chatwork api room resp data.
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

// RoomsCreateData chatwork api rooms create resp data.
type RoomsCreateData struct {
	RoomID int `json:"room_id"`
}

// RoomsUpdateData chatwork api rooms update resp data.
type RoomUpdateData RoomsCreateData

// RoomsCreateParam chatwork api rooms create request param.
type RoomsCreateParam struct {
	Description        *optional.NullableString   `json:"description"`
	IconPreset         *optional.NullableString   `json:"icon_preset"`
	Link               *optional.NullableBool     `json:"link"`
	LinkCode           *optional.NullableString   `json:"link_code"`
	LinkNeedAcceptance *optional.NullableBool     `json:"link_need_acceptance"`
	MembersAdminIds    *param.IntArrayParam       `json:"members_admin_ids"`
	MembersMemberIds   *optional.NullableIntArray `json:"members_member_ids"`
	MembersReadonlyIds *optional.NullableIntArray `json:"members_readonly_ids"`
	Name               string                     `json:"name"`
}

// RoomsUpdateParam chatwork api rooms update request param.
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

// NewRoomsResource new chatwork api rooms resource.
func NewRoomsResource(client *api.ChatworkApiClient) RoomsResource {
	data := RoomsResource{
		ResourceName: `/rooms`,
		Client:       client,
	}
	return data
}

// List chatwork api get rooms list.
func (c RoomsResource) List() ([]RoomData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      nil,
	}

	result := make([]RoomData, 0, 32)

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

// Create chatwork api create room.
func (c RoomsResource) Create(params RoomsCreateParam) (RoomsCreateData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPost,
		ResouceName: c.ResourceName,
		Params:      p,
	}

	result := RoomsCreateData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

// Create chatwork api get room data.
func (c RoomsResource) Get(room_id int) (RoomData, error) {

	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, room_id),
		Params:      nil,
	}

	result := RoomData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

// Create chatwork api update room data.
func (c RoomsResource) Update(room_id int, params RoomsUpdateParam) (RoomUpdateData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, room_id),
		Params:      p,
	}

	result := RoomUpdateData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}

// Create chatwork api delete room.
func (c RoomsResource) Delete(room_id int, action string) error {

	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodDelete,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, room_id),
		Params: map[string]*optional.NullableString{
			"action_type": optional.String(action),
		},
	}

	_, err := c.Client.Call(spec)

	return err
}

// Files chatwork api rooms/files resource.
func (c RoomsResource) Files() rooms.RoomFilesResource {
	return rooms.NewRoomFilesResource(c.ResourceName, c.Client)
}

// Link chatwork api rooms/link resource.
func (c RoomsResource) Link() rooms.RoomLinkResource {
	return rooms.NewRoomLinkResource(c.ResourceName, c.Client)
}

// Members chatwork api rooms/members resource.
func (c RoomsResource) Members() rooms.RoomMembersResource {
	return rooms.NewRoomMembersResource(c.ResourceName, c.Client)
}

// Message chatwork api rooms/message resource.
func (c RoomsResource) Message() rooms.RoomMessagesResource {
	return rooms.NewRoomMessagesResource(c.ResourceName, c.Client)
}

// Tasks chatwork api rooms/tasks resource.
func (c RoomsResource) Tasks() rooms.RoomTasksResource {
	return rooms.NewRoomTasksResource(c.ResourceName, c.Client)
}
