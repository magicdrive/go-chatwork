package rooms

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

// RoomFilesResource chatwork api rooms/files resouce
type RoomFilesResource struct {
	ResourceName string
	Client       *api.ChatworkApiClient
}

// RoomFileData chatwork api resp rooms/file data
type RoomFileData struct {
	FileID  int `json:"file_id"`
	Account struct {
		AccountID      int    `json:"account_id"`
		Name           string `json:"name"`
		AvatarImageURL string `json:"avatar_image_url"`
	} `json:"account"`
	MessageID  string `json:"message_id"`
	Filename   string `json:"filename"`
	Filesize   int    `json:"filesize"`
	UploadTime int    `json:"upload_time"`
}

// RoomFileUploadData chatwork api resp rooms/file upload data.
type RoomFileUploadData struct {
	Public *optional.NullableInt `json:"public"`
}

// NewRoomFilesResource new chatwork api rooms/files resouce.
func NewRoomFilesResource(parent string, client *api.ChatworkApiClient) RoomFilesResource {
	data := RoomFilesResource{
		ResourceName: parent + `/%d/files`,
		Client:       client,
	}
	return data
}

// List chatwork api get rooms/files list.
func (c RoomFilesResource) List(room_id int, account_id *optional.NullableInt) ([]RoomFileData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params: map[string]*optional.NullableString{
			"account_id": account_id.ToNullableString(),
		},
	}

	result := make([]RoomFileData, 0, 32)

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

// Upload chatwork api upload rooms/files.
func (c RoomFilesResource) Upload(room_id int, filepath string, message *optional.NullableString) (RoomFileUploadData, error) {

	file_entity, err := os.Open(filepath)
	if err != nil {
		return RoomFileUploadData{}, err
	}
	params := map[string]io.Reader{
		"file": file_entity,
	}
	if message.IsPresent() {
		s := message.Valid().Get()
		params["message"] = strings.NewReader(s)
	}

	spec := api.ApiSpecMultipart{
		Credential:  c.Client.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      params,
	}

	result := RoomFileUploadData{}

	str, err := c.Client.CallMultipart(spec)
	if err == nil {
		json.Unmarshal([]byte(str), result)
	}

	return result, err
}

// Upload chatwork api get rooms/file.
func (c RoomFilesResource) Get(room_id int, file_id int, create_download_flag *optional.NullableBool) (RoomFileData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName+`/%d`, room_id, file_id),
		Params: map[string]*optional.NullableString{
			"create_download_flag": create_download_flag.ToNullableString(),
		},
	}

	result := RoomFileData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err

}
