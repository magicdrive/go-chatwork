package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetMe(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	target := client.Me()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "account_id": 123,
	  "room_id": 322,
	  "name": "John Smith",
	  "chatwork_id": "tarochatworkid",
	  "organization_id": 101,
	  "organization_name": "Hello Company",
	  "department": "Marketing",
	  "title": "CMO",
	  "url": "http://mycompany.example.com",
	  "introduction": "Self Introduction",
	  "mail": "taro@example.com",
	  "tel_organization": "XXX-XXXX-XXXX",
	  "tel_extension": "YYY-YYYY-YYYY",
	  "tel_mobile": "ZZZ-ZZZZ-ZZZZ",
	  "skype": "myskype_id",
	  "facebook": "myfacebook_id",
	  "twitter": "mytwitter_id",
	  "avatar_image_url": "https://example.com/abc.png",
	  "login_mail": "account@example.com"
	}
	`

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s%s", api.ApiEndpoint, target.ResourceName),
		httpmock.NewStringResponder(http.StatusOK, mock_json),
	)

	actual, err := target.Get()
	assert.Nil(t, err)

	expected := resource.MeData{}
	err = json.Unmarshal([]byte(mock_json), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
