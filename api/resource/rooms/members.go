package rooms

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

type RoomMembersResource struct {
	ResourceName string
	Credential   string
}

type RoomMemberData struct {
	AccountID        int    `json:"account_id"`
	Role             string `json:"role"`
	Name             string `json:"name"`
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	AvatarImageURL   string `json:"avatar_image_url"`
}

type RoomMembersAuthorityData struct {
	Admin    []int `json:"admin"`
	Member   []int `json:"member"`
	Readonly []int `json:"readonly"`
}

type RoomMembersUpdateParam struct {
	MembersAdminIds    *param.IntArrayParam       `json:"members_admin_ids"`
	MembersMemberIds   *optional.NullableIntArray `json:"members_member_ids"`
	MembersReadonlyIds *optional.NullableIntArray `json:"members_readonly_ids"`
}

func NewRoomMembersResource(parent string, credential string) RoomMembersResource {
	data := RoomMembersResource{
		ResourceName: parent + `/%d/members`,
		Credential:   credential,
	}
	return data
}

func (c RoomMembersResource) List(room_id int, force *optional.NullableBool) ([]RoomMemberData, error) {

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params: map[string]*optional.NullableString{
			"force": force.ToNullableString(),
		},
	}

	result := make([]RoomMemberData, 0, 32)

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c RoomMembersResource) Update(room_id int, params RoomMembersUpdateParam) (RoomMembersAuthorityData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      p,
	}

	result := RoomMembersAuthorityData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}
