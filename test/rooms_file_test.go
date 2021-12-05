package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param/optional"
	"github.com/magicdrive/go-chatwork/api/resource/rooms"
	"github.com/stretchr/testify/assert"
)

func TestGetRoomsFileUpload(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	room_id := 1

	message := optional.String("Hello.")

	dir, _ := os.Getwd()
	file_path := dir + "/testdata/upload_file_test.txt"

	target := client.Rooms().Files()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "file_id": 1234
	}
	`

	httpmock.RegisterResponder(http.MethodPost,
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName, room_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	actual, err := target.Upload(room_id, file_path, message)
	assert.Nil(t, err)

	expected := rooms.RoomFileUploadData{}
	err = json.Unmarshal([]byte(mock_json), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestRoomsFileGet(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	room_id := 1
	file_id := 1
	create_download_flag := optional.NilBool()

	target := client.Rooms().Files()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "file_id":3,
	  "account": {
		"account_id":123,
		"name":"Bob",
		"avatar_image_url": "https://example.com/ico_avatar.png"
	  },
	  "message_id": "22",
	  "filename": "README.md",
	  "filesize": 2232,
	  "upload_time": 1384414750
	}
	`

	httpmock.RegisterResponder(http.MethodGet,
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName+`/%d`, room_id, file_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	actual, err := target.Get(room_id, file_id, create_download_flag)
	assert.Nil(t, err)

	expected := rooms.RoomFileData{}
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestRoomsFileList(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	room_id := 1
	account_id := optional.Int(1)

	target := client.Rooms().Files()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	[
	  {
		"file_id": 3,
		"account": {
		  "account_id": 123,
		  "name": "Bob",
		  "avatar_image_url": "https://example.com/ico_avatar.png"
		},
		"message_id": "22",
		"filename": "README.md",
		"filesize": 2232,
		"upload_time": 1384414750
	  }
	]
	`

	httpmock.RegisterResponder(http.MethodGet,
		fmt.Sprintf(fmt.Sprintf("%s%s", client.Client.AltChatworkApiHost, target.ResourceName), room_id),
		httpmock.NewStringResponder(200, mock_json),
	)

	actual, err := target.List(room_id, account_id)
	assert.Nil(t, err)

	expected := make([]rooms.RoomFileData, 0, 32)
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
