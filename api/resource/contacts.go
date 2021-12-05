package resource

import (
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
)

type ContactsResource struct {
	ResourceName string
	Client       *api.ChatworkApiClient
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

func NewContactsResource(client *api.ChatworkApiClient) ContactsResource {
	data := ContactsResource{
		ResourceName: `/contacts`,
		Client:       client,
	}
	return data
}

func (c ContactsResource) List() ([]ContactData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      nil,
	}

	result := make([]ContactData, 0, 32)

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}
