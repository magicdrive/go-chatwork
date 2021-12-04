package example

import (
	"fmt"

	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

func PostMessage() {

	client := chatwork.NewChatworkClient(`your-api-key-here`)

	room_id := 1
	params := chatwork.RoomMessagePostParam{Body: "Hi there.", SelfUnread: optional.BoolTrue()}

	if r, err := client.Rooms().Message().Post(room_id, params); err == nil {
		fmt.Println(r.MessageId)
	} else {
		panic(err)
	}

}
