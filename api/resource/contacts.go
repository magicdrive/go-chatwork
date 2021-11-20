package resource

import (
	"net/http"

	json "github.com/goccy/go-json"
)

type ContactsResource struct {
	ResourceName string
	Credential   string
}

type ContactData struct {
	AccountId        int    `json:"account_id"`
	RoomId           int    `json:"room_id"`
	Name             string `json:"name"`
	ChatworkId       string `json:"chatwork_id"`
	OrganizationId   int    `json:"organizationId"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	AvatarImageUrl   string `json:"avatar_image_url"`
}

func Contacts(credential string) ContactsResource {
	data := ContactsResource{
		ResourceName: `/contacts`,
		Credential:   credential,
	}
	return data
}

func (c ContactsResource) Get() ([]ContactData, error) {
	spec := ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      nil,
	}

	result := make([]ContactData, 32)

	str, err := Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}
