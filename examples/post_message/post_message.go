package examples

import (
	"fmt"

	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

func main() {

	client := chatwork.NewChatworkClient(`your-api-key-here`)

	roomID := 1
	params := chatwork.RoomMessagePostParam{Body: "Hi there.", SelfUnread: optional.BoolTrue()}

	if r, err := client.Rooms().Messages().Post(roomID, params); err == nil {
		fmt.Println(r.MessageID)
	} else {
		panic(err)
	}

}
