package rooms_test

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

	roomID := 1

	message := optional.String("Hello.")

	dir, _ := os.Getwd()
	filePath := dir + "/testdata/upload_file_test.txt"

	target := client.Rooms().Files()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "file_id": 1234
	}
	`

	httpmock.RegisterResponder(http.MethodPost,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName, roomID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	actual, err := target.Upload(roomID, filePath, message)
	assert.Nil(t, err)

	expected := rooms.RoomFileUploadData{}
	err = json.Unmarshal([]byte(mockBody), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestRoomsFileGet(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	roomID := 1
	fileID := 1
	createDownloadFlag := optional.NilBool()

	target := client.Rooms().Files()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
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
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName+`/%d`, roomID, fileID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	actual, err := target.Get(roomID, fileID, createDownloadFlag)
	assert.Nil(t, err)

	expected := rooms.RoomFileData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestRoomsFileList(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	roomID := 1
	accountID := optional.Int(1)

	target := client.Rooms().Files()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
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
		fmt.Sprintf(fmt.Sprintf("%s%s", client.Client.APIEndpoint, target.ResourceName), roomID),
		httpmock.NewStringResponder(200, mockBody),
	)

	actual, err := target.List(roomID, accountID)
	assert.Nil(t, err)

	expected := make([]rooms.RoomFileData, 0, 32)
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
