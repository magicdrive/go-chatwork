package rooms

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/optional"
)

type MembersResource struct {
	ResourceName string
	Credential   string
}

type MemberData struct {
	AccountID        int    `json:"account_id"`
	Role             string `json:"role"`
	Name             string `json:"name"`
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	AvatarImageURL   string `json:"avatar_image_url"`
}

type MembersRoomAuthorityData struct {
	Admin    []int `json:"admin"`
	Member   []int `json:"member"`
	Readonly []int `json:"readonly"`
}

type MembersUpdateParam struct {
	MembersAdminIds    []int                   `json:"members_admin_ids"`
	MembersMemberIds   []*optional.NullableInt `json:"members_member_ids"`
	MembersReadonlyIds []*optional.NullableInt `json:"members_readonly_ids"`
}

func NewMembersResource(parent string, credential string) MembersResource {
	data := MembersResource{
		ResourceName: parent + `/%d/members`,
		Credential:   credential,
	}
	return data
}

func (c MembersResource) List(room_id int, force *optional.NullableBool) ([]MemberData, error) {

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params: map[string]*optional.NullableString{
			"force": force.ToNullableString(),
		},
	}

	result := make([]MemberData, 0, 32)

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c MembersResource) Update(room_id int, params MembersUpdateParam) (MembersRoomAuthorityData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JsonToMap(b)

	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      p,
	}

	result := MembersRoomAuthorityData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}
