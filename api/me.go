package api

import (
	"net/http"

	json "github.com/goccy/go-json"
)

type MeResource struct {
	ResourceName string
	Credential   string
}

type MeData struct {
	AccountID        int    `json:"account_id"`
	RoomID           int    `json:"room_id"`
	Name             string `json:"name"`
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	Title            string `json:"title"`
	URL              string `json:"url"`
	Introduction     string `json:"introduction"`
	Mail             string `json:"mail"`
	TelOrganization  string `json:"tel_organization"`
	TelExtension     string `json:"tel_extension"`
	TelMobile        string `json:"tel_mobile"`
	Skype            string `json:"skype"`
	Facebook         string `json:"facebook"`
	Twitter          string `json:"twitter"`
	AvatarImageURL   string `json:"avatar_image_url"`
	LoginMail        string `json:"login_mail"`
}

func NewMe(credential string) ContactsResource {
	data := ContactsResource{
		ResourceName: `/me`,
		Credential:   credential,
	}
	return data
}

func (c MeResource) Get() (MeData, error) {
	spec := ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      nil,
	}

	result := MeData{}

	str, err := Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), result)
	}

	return result, err
}
