package rooms_sub

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
)

type FilesResource struct {
	ResourceName string
	Credential   string
}

type FileData struct {
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

type FileUploadData struct {
	Public bool `json:"public"`
}

func NewFilesResource(parent string, credential string) FilesResource {
	data := FilesResource{
		ResourceName: parent + `/%d/files`,
		Credential:   credential,
	}
	return data
}

func (c FilesResource) List(room_id int) ([]FileData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params:      nil,
	}

	result := make([]FileData, 0, 32)

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), result)
	}

	return result, err
}

func (c FilesResource) Upload(room_id int, filepath string, message string) (FileUploadData, error) {

	file_entity, err := os.Open(filepath)
	if err != nil {
		return FileUploadData{}, err
	}

	spec := api.ApiSpecMultipart{
		Credential:  c.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, room_id),
		Params: map[string]io.Reader{
			"file":    file_entity,
			"message": strings.NewReader("hello world!"),
		},
	}

	result := FileUploadData{}

	str, err := api.CallMultipart(spec)
	if err == nil {
		json.Unmarshal([]byte(str), result)
	}

	return result, err
}

func (c FilesResource) Get(room_id int, file_id int, create_download_flag int) (FileData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf(c.ResourceName+"/%d", room_id, file_id),
		Params:      map[string]string{"create_download_flag": strconv.Itoa(create_download_flag)},
	}

	result := FileData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), result)
	}

	return result, err

}
